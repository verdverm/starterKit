package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/zaha-io/zaha/apps"
)

func create_app_noninteractive(appname string) (cfg apps.AppConfig, err error) {
	appcfg, err := readAppConfig("default_app")
	if err != nil {
		panic(err)
		return cfg, err
	}
	printAppConfig(appcfg)

	appcfg.Name = appname

	appcfg.SourcePath, err = os.Getwd()
	if err != nil {
		panic(err)
		return cfg, err
	}
	appcfg.SourcePath += "/" + appname
	appcfg.Description = appname + "... a Zaha app"

	printAppConfig(appcfg)
	// writeAppConfig(APPCFG)

	err = initAppDirectory(appcfg)
	panicOnError(err, "error initializing app directory")

	return appcfg, nil
}

func initAppDirectory(appcfg apps.AppConfig) error {
	// create app base directory in current directory
	tmplpath := ZAHACONFIG.ConfigAbs + "template"
	dockpath := ZAHACONFIG.ConfigAbs + "dockers/"
	wapppath := ZAHACONFIG.ConfigAbs + "webapp/"

	apppath := appcfg.SourcePath
	err := os.Mkdir(apppath, 0755)
	if err != nil {
		return err
	}

	// copy temlate
	err = mustCopyDir(apppath, tmplpath, nil)
	panicOnError(err, "copying template directory")

	// copy dockers
	docks := []string{}
	docks = append(docks, appcfg.Frontends...)
	docks = append(docks, appcfg.Services...)
	docks = append(docks, appcfg.Databases...)
	docks = append(docks, appcfg.Backends...)
	for _, d := range docks {
		// copy yaml file
		mustCopyFile(apppath+"/dockers/"+d+".yaml", dockpath+d+".yaml")

		// create storage directories (this should check whether the docker needs storage or not)
		err = os.MkdirAll(apppath+"/storage/"+d, 0755)

		// add code to somewhere to be injected into the template
	}

	// copy webapp template & fill in this map that is used for rendering
	err = mustCopyDir(apppath, wapppath, nil)
	panicOnError(err, "copying template directory")

	return nil
}

func create_app_interactive(appname string) (err error) {
	appcfg := apps.AppConfig{}
	appcfg.Name = appname

	appcfg.SourcePath, err = os.Getwd()
	if err != nil {
		return err
	}
	appcfg.SourcePath += "/" + appname

	// fmt.Println("  dir: ", appcfg.SourcePath)

	// fmt.Println("\nPlease enter a short destription of the app:")
	// reader := bufio.NewReader(os.Stdin)
	// resp, err := reader.ReadString('\n')
	// if err != nil {
	// 	return err
	// }
	appcfg.Description = appname + "... a Zaha app"

	// fmt.Println("\n\n")
	fmt.Println("\nFor the questions, please enter numbers separated by spaces.\n")

	fes, err := question(fe_qna)
	if err != nil {
		return err
	}
	appcfg.Frontends = fes

	dbs, err := question(db_qna)
	if err != nil {
		return err
	}
	appcfg.Databases = dbs

	bes, err := question(be_qna)
	if err != nil {
		return err
	}
	appcfg.Backends = bes

	fmt.Println(appname + " configuration complete:\n")
	err = printAppConfig(appcfg)
	if err != nil {
		return err
	}

	return nil
}

// Yaml this

func question(qna qna_struct) (responses []string, err error) {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println(qna.text)
	for i, d := range qna.descrips {
		fmt.Printf("   %d: %s\n", i+1, d)
	}
	fmt.Print("> ")
	resp, _ := reader.ReadString('\n')

	flds := strings.Fields(resp)
	for _, f := range flds {
		n, err := strconv.Atoi(f)
		if err != nil {
			fmt.Println("Bad input: ", f)
			return nil, err
		}
		n--
		if n < 0 || n >= len(qna.choices) {
			fmt.Println("Bad index; ", n+1)
			return nil, errors.New("invalid choice index")
		}
		responses = append(responses, qna.choices[n])
	}

	return responses, nil
}

type qna_struct struct {
	text     string
	choices  []string
	descrips []string
}

var (
	fe_qna = qna_struct{
		text: "Which frontends?  ",
		choices: []string{
			"nginx",
			"flask",
			"redis",
		},
		descrips: []string{
			"nginx     (proxy)",
			"flask     (framework)",
			"redis     (cache)",
		},
	}

	db_qna = qna_struct{
		text: "Which databases?  ",
		choices: []string{
			"psql",
			"neo4j",
			"couchdb",
		},
		descrips: []string{
			"psql      (sql)",
			"neo4j     (graph)",
			"couchdb   (doc)",
		},
	}

	be_qna = qna_struct{
		text: "Which databases?  ",
		choices: []string{
			"rabbitmq",
			"ipynb",
		},
		descrips: []string{
			"rabbitmq  (messaging)",
			"ipython   (the notebook)",
		},
	}
)
