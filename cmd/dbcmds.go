package main

import (
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/fsouza/go-dockerclient"

	"github.com/zaha-io/zaha/apps"
)

func db_initialize(c *cli.Context) {
	if len(c.Args()) != 0 {
		fmt.Println("Error: bad args to 'db initialize' !!!")
		cli.ShowSubcommandHelp(c)
		return
	}

	db_cmds := [][]string{
		[]string{"psql", "init"},
		[]string{"psql", "migrate"},
		[]string{"psql", "upgrade"},
		[]string{"dbfill"},
	}

	opts, err := apps.ReadDockerOpts("flask", APPCFG)
	if err != nil {
		fmt.Println("while getting docker options for: flask\n" + err.Error())
		return
	}

	apps.PrepDocker(&opts, APPCFG)

	opts.Name = "dbinit." + opts.CreateOpts.Config.Domainname
	opts.CreateOpts.Name = opts.Name
	opts.CreateOpts.Config.Hostname = "dbinit"
	opts.CreateOpts.Config.ExposedPorts = nil
	opts.StartOpts.PortBindings = nil

	// printDockerOpts(opts)

	client, err := docker.NewClient(ZAHACONFIG.DockerHost)
	if err != nil {
		fmt.Println("while getting docker client: \n" + err.Error())
		return
	}

	for _, cmd := range db_cmds {
		fmt.Println("dbinit: ", cmd)

		// setup the command
		cmd_n_args := []string{
			"/sbin/my_init",
			"python",
			"/src/run.py",
		}
		cmd_n_args = append(cmd_n_args, cmd...)
		opts.CreateOpts.Config.Cmd = cmd_n_args

		// run the docker
		err = apps.StartDocker(&opts)
		if err != nil {
			fmt.Println("while init'n DB\n" + err.Error())
			return
		}

		// wait for the docker to exit
		done := false
		exit := 0
		for !done {
			dock, err := client.InspectContainer(opts.Name)
			if err != nil {
				fmt.Println("while inspecting container: ", opts.Name, "\n"+err.Error())
				return
			} else if !dock.State.Running {
				exit = dock.State.ExitCode
				if exit != 0 {
					fmt.Printf("%v %v", dock.State.ExitCode, dock.State.Error)
					return
				}
				break
			}
		}

		// remove docker
		err = apps.RemoveDocker(opts.Name)
		if err != nil {
			fmt.Println("while removing container: ", opts.Name, "\n"+err.Error())
			return
		}
	}

}

func db_migrate(c *cli.Context) {
	if len(c.Args()) != 0 {
		fmt.Println("Error: bad args to 'db migrate' !!!")
		cli.ShowSubcommandHelp(c)
		return
	}
	// migrate_cmd := "db migrate"
	// upgrade_cmd := "db upgrade"
	fmt.Println("Not implemented yet...")
}

func db_reset(c *cli.Context) {
	if len(c.Args()) != 0 {
		fmt.Println("Error: bad args to 'db reset' !!!")
		cli.ShowSubcommandHelp(c)
		return
	}

	fmt.Println("Not implemented yet...")
}

func db_export(c *cli.Context) {
	if len(c.Args()) != 0 {
		fmt.Println("Error: bad args to 'db export' !!!")
		cli.ShowSubcommandHelp(c)
		return
	}
	fmt.Println("Not implemented yet...")
}

func db_import(c *cli.Context) {
	if len(c.Args()) != 0 {
		fmt.Println("Error: bad args to 'db import' !!!")
		cli.ShowSubcommandHelp(c)
		return
	}
	fmt.Println("Not implemented yet...")
}

func db_fix(c *cli.Context) {
	if len(c.Args()) != 0 {
		fmt.Println("Error: bad args to 'db fix' !!!")
		cli.ShowSubcommandHelp(c)
		return
	}

	opts, err := apps.ReadDockerOpts("flask", APPCFG)
	if err != nil {
		fmt.Println("while getting docker options for: flask\n" + err.Error())
		return
	}

	apps.PrepDocker(&opts, APPCFG)

	opts.Name = "dbfix." + opts.CreateOpts.Config.Domainname
	opts.CreateOpts.Name = opts.Name
	opts.CreateOpts.Config.Hostname = "dbfix"
	opts.CreateOpts.Config.ExposedPorts = nil
	opts.StartOpts.PortBindings = nil

	// printDockerOpts(opts)

	client, err := docker.NewClient(ZAHACONFIG.DockerHost)
	if err != nil {
		fmt.Println("while getting docker client: \n" + err.Error())
		return
	}

	// setup the command
	cmd_n_args := []string{
		"/sbin/my_init",
		"python",
		"/src/run.py",
		"dbfix",
	}
	fmt.Println("db fix: ", cmd_n_args)

	opts.CreateOpts.Config.Cmd = cmd_n_args

	// run the docker
	err = apps.StartDocker(&opts)
	if err != nil {
		fmt.Println("while init'n DB\n" + err.Error())
		return
	}

	// wait for the docker to exit
	done := false
	exit := 0
	for !done {
		dock, err := client.InspectContainer(opts.Name)
		if err != nil {
			fmt.Println("while inspecting container: ", opts.Name, "\n"+err.Error())
			return
		} else if !dock.State.Running {
			exit = dock.State.ExitCode
			if exit != 0 {
				fmt.Printf("%v %v", dock.State.ExitCode, dock.State.Error)
				return
			}
			break
		}
	}

	// remove docker
	err = apps.RemoveDocker(opts.Name)
	if err != nil {
		fmt.Println("while removing container: ", opts.Name, "\n"+err.Error())
		return
	}

}
