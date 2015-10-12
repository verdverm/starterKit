package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/codegangsta/cli"

	"github.com/zaha-io/zaha/apps"
	"github.com/zaha-io/zaha/mesos"
	"github.com/zaha-io/zaha/providers"
)

func set_app(c *cli.Context) {
	if len(c.Args()) != 1 {
		fmt.Println("Error: bad args to 'app set'")
		cli.ShowSubcommandHelp(c)
		return
	}
	appname := c.Args()[0]

	cfg, err := readAppConfig(appname)
	if err != nil {
		fmt.Println("no app with that name...  use 'zaha create " + appname + "'")
		return
	}

	APPCFG = cfg
	ZAHACONFIG.CurrentApp = appname
	writeZahaConfig()
}

func create_app(c *cli.Context) {
	if len(c.Args()) != 1 {
		fmt.Println("Error: bad args to 'app create'")
		cli.ShowSubcommandHelp(c)
		return
	}

	appname := c.Args()[0]

	// How to handle multiple apps with the same name?
	// i.e. we have two copies of the app in different directories
	_, err := readAppConfig(appname)
	if err == nil {
		fmt.Println("App with that name already exists, please choose another")
		return
	}

	fmt.Println("\n\nZaha!   creating: ", appname)
	APPCFG, err = create_app_noninteractive(appname)
	if err != nil {
		fmt.Println("Error creating app: ", err)
		return
	}
}

func start_app(c *cli.Context) {
	args := c.Args()
	L := len(args)

	// check we have a current app and some action to perform
	checkCurrApp(c)
	if L > 2 {
		fmt.Println("bad args to 'start <component> [cloudname]")
		cli.ShowSubcommandHelp(c)
		return
	}
	action := "start"
	subject := "app"
	if L >= 1 {
		subject = args[0]
	}

	location := ""
	if L == 2 {
		location = args[1]
	}
	err := do_action_on_app(action, subject, location)
	if err != nil {
		msg := fmt.Sprint("Error during "+action+" "+subject+": ", err.Error())
		panicOnError(err, msg)
	}
}

func status_app(c *cli.Context) {
	args := c.Args()
	L := len(args)

	// check we have a current app and some action to perform
	checkCurrApp(c)
	if L != 2 {
		fmt.Println("bad args to 'status <component> <cloudname>")
		cli.ShowSubcommandHelp(c)
		return
	}
	action := "status"
	subject := args[0]

	location := args[1]

	err := do_action_on_app(action, subject, location)
	if err != nil {
		msg := fmt.Sprint("Error during "+action+" "+subject+": ", err.Error())
		panicOnError(err, msg)
	}
}

func stop_app(c *cli.Context) {
	args := c.Args()
	L := len(args)

	// check we have a current app and some action to perform
	checkCurrApp(c)
	if L > 2 {
		fmt.Println("bad args to 'stop <component> [cloudname]")
		cli.ShowSubcommandHelp(c)
		return
	}
	action := "stop"
	subject := "app"
	if L >= 1 {
		subject = args[0]
	}

	location := ""
	if L == 2 {
		location = args[1]
	}
	err := do_action_on_app(action, subject, location)
	if err != nil {
		msg := fmt.Sprint("Error during "+action+" "+subject+": ", err.Error())
		panicOnError(err, msg)
	}
}

func restart_app(c *cli.Context) {
	args := c.Args()
	L := len(args)

	// check we have a current app and some action to perform
	checkCurrApp(c)
	if L > 2 {
		fmt.Println("bad args to 'status <component> [cloudname]")
		cli.ShowSubcommandHelp(c)
		return
	}
	action := "restart"
	subject := "app"
	if L >= 1 {
		subject = args[0]
	}

	location := ""
	if L == 2 {
		location = args[1]
	}
	err := do_action_on_app(action, subject, location)
	if err != nil {
		msg := fmt.Sprint("Error during "+action+" "+subject+": ", err.Error())
		panicOnError(err, msg)
	}
}

func do_action_on_app(action, subject, location string) error {
	fmt.Printf("Zaha!  %s: %s\n", action, subject)
	// turn the subject into a list of dockers to act upon
	dock_list, err := apps.GetAppDockerList(subject, APPCFG)
	if err != nil {
		return errors.New("while getting docker list: " + subject + "\n" + err.Error())
	}

	// read and collect the docker options
	var dock_opts []*apps.DockerOptions
	for _, dock := range dock_list {
		// fn := APPCFG.SourcePath + fmt.Sprintf("/dockers/%s.yaml", dock)
		opts, err := apps.ReadDockerOpts(dock, APPCFG)
		if err != nil {
			return errors.New("while getting docker list: " + subject + "\n" + err.Error())
		}

		dock_opts = append(dock_opts, &opts)
	}

	// perform the requested action upon the dockers
	switch action {
	case "start":
		err = start_app_dockers(location, dock_opts)
	case "stop":
		err = stop_app_dockers(location, dock_opts)
	case "restart":
		err = restart_app_dockers(location, dock_opts)
	case "status":
		err = status_app_dockers(location, dock_opts)
	}
	if err != nil {
		return err
	}
	return nil
}

func start_app_dockers(location string, dock_list []*apps.DockerOptions) error {
	apps.ReverseDockerList(dock_list)
	for _, opts := range dock_list {
		if location == "" {
			err := apps.PrepDocker(opts, APPCFG)
			if err != nil {
				return errors.New("error preparing docker: " + opts.Name + "\n" + err.Error())
			}
			err = apps.StartDocker(opts)
			if err != nil {
				return errors.New("error starting docker: " + opts.Name + "\n" + err.Error())
			}
		} else {
			ccfg, err := readCloudConfig(APPCFG.Name, location)
			if err != nil {
				return err
			}

			err = mesos.StartMesosphereDocker(opts, APPCFG, ccfg.Name, ccfg.URI)
			if err != nil {
				return err
			}

			// some sort of sleep or wait check on dependent dockers
		}
	}
	return nil
}

func stop_app_dockers(location string, dock_list []*apps.DockerOptions) error {
	for _, opts := range dock_list {
		if location == "" {
			err := apps.StopDocker(opts)
			if err != nil && !strings.Contains(err.Error(), "non-existant") {
				return errors.New("error stopping docker: " + opts.Name + "\n" + err.Error())
			}
		} else {
			ccfg, err := readCloudConfig(APPCFG.Name, location)
			if err != nil {
				return err
			}

			err = mesos.StopMesosphereDocker(ccfg.URI, opts.Name)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func restart_app_dockers(location string, dock_list []*apps.DockerOptions) error {
	err := stop_app_dockers(location, dock_list)
	if err != nil {
		return err
	}

	err = start_app_dockers(location, dock_list)
	if err != nil {
		return err
	}

	return nil
}

func status_app_dockers(location string, dock_list []*apps.DockerOptions) error {
	if location == "" {
		for _, opts := range dock_list {
			err := apps.StatusDocker(opts)
			if err != nil {
				return errors.New("error statusing docker: " + opts.Name + "\n" + err.Error())
			}
		}
	} else {
		ccfg, err := readCloudConfig(APPCFG.Name, location)
		if err != nil {
			return err
		}

		appTasks, err := mesos.StatusMesosphereTasks(APPCFG, ccfg.Name, ccfg.URI)
		if err != nil {
			return err
		}

		fmt.Printf("  %-20s  %s\n", "COMPONENT", "STATUS")

		for _, opts := range dock_list {
			found := false
			for _, task := range appTasks.Tasks {
				if opts.Name == task.AppId[1:] {
					fmt.Printf("  %-20s   %-16s %v\n", opts.Name, task.Host, task.Ports)
					found = true
					break
				}
			}
			if !found {
				fmt.Printf("  %-20s   non-existant\n", opts.Name)
			}
		}

	}
	return nil
}

func checkCurrApp(c *cli.Context) {
	if ZAHACONFIG.CurrentApp == "none" || ZAHACONFIG.CurrentApp == "" {
		fmt.Println("no app set, please set one first")
		cli.ShowSubcommandHelp(c)
		os.Exit(1)
	}
	var err error
	APPCFG, err = readAppConfig(ZAHACONFIG.CurrentApp)
	panicOnError(err, "reading app config")
}

func push_app(c *cli.Context) {

	args := c.Args()
	L := len(args)

	// check we have a current app and some action to perform
	checkCurrApp(c)
	if L != 2 {
		fmt.Println("bad args to 'push <appname> <cloudname>")
		cli.ShowSubcommandHelp(c)
		return
	}
	appname := args[0]
	cloudname := args[1]

	acfg, err := readAppConfig(appname)
	if err != nil {
		fmt.Println("error reading app config for: ", appname)
		return
	}

	ccfg, err := readCloudConfig(acfg.Name, cloudname)
	if err != nil {
		fmt.Println("error reading cloud config for: ", cloudname)
		return
	}

	fmt.Printf("Pushing %s to %s (%s)\n", acfg.Name, ccfg.Name, ccfg.URI)

	err = providers.CloudPushApp(ccfg, acfg)
	if err != nil {
		fmt.Println("error pushing", acfg.Name, "to", ccfg.Name)
		return
	}
	fmt.Println("Zaha! done pushing app, you can now (re)start it.")
}
