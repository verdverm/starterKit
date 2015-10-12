package main

import (
	"archive/zip"
	"fmt"
	// "io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/codegangsta/cli"
	"gopkg.in/yaml.v2"

	"github.com/zaha-io/zaha/apps"
	"github.com/zaha-io/zaha/providers"
)

type ZahaConfig struct {
	Description string "description"
	DockerHost  string "dockerhost"
	CurrentUser string "user"
	CurrentApp  string "app"
	ConfigAbs   string "-"
}

var defaultZahaConfig = `
description: "Zaha Configuration"
dockerhost: "unix:///var/run/docker.sock"
user: none
app:  none
`

func config_app(c *cli.Context) {
	if len(c.Args()) != 0 {
		fmt.Println("Error: bad args to 'config app'")
		cli.ShowSubcommandHelp(c)
		return
	}
	fmt.Println("Zaha! ", APPCFG.Name, "config.")
}

func config_show(c *cli.Context) {
	if len(c.Args()) != 0 {
		fmt.Println("Error: bad args to 'config show'")
		cli.ShowSubcommandHelp(c)
		return
	}
	printZahaConfig()
	printAppConfig(APPCFG)
}

func readZahaConfig() (ZahaConfig, error) {
	var zahacfg ZahaConfig
	data, err := ioutil.ReadFile(config_fn_abs)
	if err != nil {
		return zahacfg, err
	}
	err = yaml.Unmarshal([]byte(data), &zahacfg)
	if err != nil {
		return zahacfg, err
	}
	return zahacfg, nil
}

func printZahaConfig() error {
	fmt.Println(ZAHACONFIG)
	d, err := yaml.Marshal(&ZAHACONFIG)
	if err != nil {
		fmt.Println("Couldn't marshal Zaha config")
		return err
	}
	fmt.Println(string(d))
	return nil
}

func writeZahaConfig() error {
	d, err := yaml.Marshal(&ZAHACONFIG)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(config_fn_abs, d, 0644)
	if err != nil {
		return err
	}
	fmt.Println("Zaha! saved zaha config file")
	return nil
}

func initZahaConfig() error {
	dotzaha_zip_url := "https://github.com/iassic/zaha-dotzaha/archive/master.zip"

	fmt.Println("creating default Zaha config file in: ", config_dir)
	err := os.MkdirAll(config_dir, 0755)
	panicOnError(err, "couldn't mkdir "+config_dir)

	resp, err := http.Get(dotzaha_zip_url)
	panicOnError(err, "couldn't get zip file from github")

	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	panicOnError(err, "couldn't read response from github")

	zip_fn := config_dir + "dotzaha.zip"
	err = ioutil.WriteFile(zip_fn, body, 0755)
	panicOnError(err, "couldn't write temporary zip file")

	reader, err := zip.OpenReader(zip_fn)
	panicOnError(err, "couldn't open temporary zip file")
	defer reader.Close()

	err = extractZipfile(config_dir, &reader.Reader)
	panicOnError(err, "couldn't extract temporary zip file")

	// remove
	return nil
}

func readAppConfig(appname string) (apps.AppConfig, error) {
	var acfg, zcfg apps.AppConfig

	zahaCopyOfConfigLocation := ZAHACONFIG.ConfigAbs + fmt.Sprintf("apps/%s/%s.yaml", appname, appname)
	// fmt.Println(zahaCopyOfConfigLocation)
	zcfg_info, err := os.Lstat(zahaCopyOfConfigLocation)
	if err != nil {
		fmt.Println("zcfg info error: " + zahaCopyOfConfigLocation)
		return zcfg, err
	}
	// fmt.Printf("%+v\n", zcfg_info)
	zdata, err := ioutil.ReadFile(zahaCopyOfConfigLocation)
	if err != nil {
		fmt.Println("zcfg read error: " + zahaCopyOfConfigLocation)
		return zcfg, err
	}
	err = yaml.Unmarshal([]byte(zdata), &zcfg)
	if err != nil {
		fmt.Println("zcfg unmarshall error: " + zahaCopyOfConfigLocation)
		return zcfg, err
	}

	appConfigLocation := zcfg.SourcePath + fmt.Sprintf("/config/%s.yaml", zcfg.Name)
	// fmt.Println(appConfigLocation)
	acfg_info, err := os.Lstat(appConfigLocation)
	if err != nil {
		fmt.Println("acfg info error: " + appConfigLocation)
		return zcfg, err
	}
	// fmt.Printf("%+v\n", acfg_info)

	// Check file times
	if zcfg_info.ModTime().Before(acfg_info.ModTime()) {
		// if the app has a newer mod time, write the file
		fmt.Println("Zaha copy of app config is stale, saving new copy of your app config to ~/.zaha/apps")
		copyFileContents(appConfigLocation, zahaCopyOfConfigLocation)
	} else if zcfg_info.ModTime().After(acfg_info.ModTime()) {
		// if we have a newer time, then something weird happened
		// this seems to always happen actually...
		// fmt.Println("Zaha copy of app config newer than your app directory copy. Please manually fix this by copying your app config to ~/.zaha/apps")
	}

	adata, err := ioutil.ReadFile(appConfigLocation)
	if err != nil {
		fmt.Println("acfg read error: " + appConfigLocation)
		return zcfg, err
	}
	err = yaml.Unmarshal([]byte(adata), &acfg)
	if err != nil {
		fmt.Println("acfg unmarshall error: ", appConfigLocation)
		return zcfg, err
	}

	return acfg, nil
}

func printAppConfig(appcfg apps.AppConfig) error {
	data, err := yaml.Marshal(&appcfg)
	if err != nil {
		return err
	}
	fmt.Println(string(data))
	return nil
}

func writeAppConfig(appcfg apps.AppConfig) error {
	d, err := yaml.Marshal(&appcfg)
	if err != nil {
		return err
	}
	fn := ZAHACONFIG.ConfigAbs + fmt.Sprintf("apps/%s/%s.yaml", appcfg.Name, appcfg.Name)
	err = ioutil.WriteFile(fn, d, 0644)
	if err != nil {
		return err
	}
	fn = appcfg.SourcePath + fmt.Sprintf("/config/%s.yaml", appcfg.Name)
	err = ioutil.WriteFile(fn, d, 0644)
	if err != nil {
		return err
	}
	fmt.Println("Zaha!   saved " + appcfg.Name + "'s config")
	return nil
}

func readCloudConfig(appname, cloudname string) (providers.CloudConfig, error) {
	var (
		acfg       apps.AppConfig
		ccfg, zcfg providers.CloudConfig
	)

	acfg, err := readAppConfig(appname)

	zahaCopyOfConfigLocation := ZAHACONFIG.ConfigAbs + fmt.Sprintf("apps/%s/clouds/%s.yaml", appname, cloudname)
	// fmt.Println(zahaCopyOfConfigLocation)
	zcfg_info, err := os.Lstat(zahaCopyOfConfigLocation)
	if err != nil {
		fmt.Println("zcfg info error: " + zahaCopyOfConfigLocation)
		return zcfg, err
	}
	// fmt.Printf("%+v\n", zcfg_info)
	zdata, err := ioutil.ReadFile(zahaCopyOfConfigLocation)
	if err != nil {
		fmt.Println("zcfg read error: " + zahaCopyOfConfigLocation)
		return zcfg, err
	}
	err = yaml.Unmarshal([]byte(zdata), &zcfg)
	if err != nil {
		fmt.Println("zcfg unmarshall error: " + zahaCopyOfConfigLocation)
		return zcfg, err
	}

	cloudConfigLocation := acfg.SourcePath + fmt.Sprintf("/clouds/%s.yaml", zcfg.Name)
	// fmt.Println(cloudConfigLocation)
	ccfg_info, err := os.Lstat(cloudConfigLocation)
	if err != nil {
		fmt.Println("ccfg info error: " + cloudConfigLocation)
		return zcfg, err
	}
	// fmt.Printf("%+v\n", ccfg_info)

	// Check file times
	if zcfg_info.ModTime().Before(ccfg_info.ModTime()) {
		// if the app has a newer mod time, write the file
		fmt.Println("Zaha copy of app cloud is stale, saving new copy of your app cloud to ~/.zaha/apps")
		copyFileContents(cloudConfigLocation, zahaCopyOfConfigLocation)
	} else if zcfg_info.ModTime().After(ccfg_info.ModTime()) {
		// if we have a newer time, then something weird happened
		// this seems to always happen actually...
		// fmt.Println("Zaha copy of app cloud newer than your app directory copy. Please manually fix this by copying your app config to ~/.zaha/apps")
	}

	adata, err := ioutil.ReadFile(cloudConfigLocation)
	if err != nil {
		fmt.Println("ccfg read error: " + cloudConfigLocation)
		return zcfg, err
	}
	err = yaml.Unmarshal([]byte(adata), &ccfg)
	if err != nil {
		fmt.Println("ccfg unmarshall error: ", cloudConfigLocation)
		return zcfg, err
	}

	return ccfg, nil
}

func printCloudConfig(cloudcfg providers.CloudConfig) error {
	data, err := yaml.Marshal(&cloudcfg)
	if err != nil {
		return err
	}
	fmt.Println(string(data))
	return nil
}

func writeCloudConfig(appname string, cloudcfg providers.CloudConfig) error {
	acfg, err := readAppConfig(appname)
	if err != nil {
		return err
	}
	d, err := yaml.Marshal(&cloudcfg)
	if err != nil {
		return err
	}
	fn := ZAHACONFIG.ConfigAbs + fmt.Sprintf("apps/%s/clouds/%s.yaml", acfg.Name, cloudcfg.Name)
	err = ioutil.WriteFile(fn, d, 0644)
	if err != nil {
		return err
	}
	fn = acfg.SourcePath + fmt.Sprintf("/clouds/%s.yaml", cloudcfg.Name)
	err = ioutil.WriteFile(fn, d, 0644)
	if err != nil {
		return err
	}
	fmt.Println("Zaha!   saved " + cloudcfg.Name + "'s config")
	return nil
}
