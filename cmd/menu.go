package main

import "github.com/codegangsta/cli"

var (
	menu = []cli.Command{
		{
			Name:   "setup",
			Usage:  "run this first!",
			Action: setupZaha,
		},
		{
			Name:   "set",
			Usage:  "set the current app",
			Action: set_app,
		},
		{
			Name:  "config",
			Usage: "manage app configuration",
			Subcommands: []cli.Command{
				{
					Name:   "app",
					Usage:  "reconfigure the current app",
					Action: config_app,
				},
				{
					Name:   "list",
					Usage:  "list the current app configuration",
					Action: config_show,
				},
			},
		},
		{
			Name:  "cloud",
			Usage: "manage app clouds",
			Subcommands: []cli.Command{
				{
					Name:   "create",
					Usage:  "create cloud for the current app | 'cloud create <name> <provider> <num_workers> [blueprint]",
					Action: cloud_create,
				},
				{
					Name:   "status",
					Usage:  "status cloud for the current app | 'cloud status <name>",
					Action: cloud_status,
				},
				{
					Name:   "resize",
					Usage:  "resize cloud for the current app | 'cloud resize <name> <size>",
					Action: cloud_resize,
				},
				{
					Name:   "destroy",
					Usage:  "destroy cloud for the current app | 'cloud destroy <name>",
					Action: cloud_destroy,
				},
				{
					Name:   "list",
					Usage:  "list the current cloud configurations",
					Action: cloud_list,
				},
			},
		},
		{
			Name:  "db",
			Usage: "manage app clouds",
			Subcommands: []cli.Command{
				{
					Name:   "init",
					Usage:  "initialize the Zaha! app databases",
					Action: db_initialize,
				},
				{
					Name:   "migrate",
					Usage:  "migrate and upgrade Zaha! app databases",
					Action: db_migrate,
				},
				{
					Name:   "reset",
					Usage:  "reset the Zaha! app databases",
					Action: db_reset,
				},
				{
					Name:   "export",
					Usage:  "export the Zaha! app databases",
					Action: db_export,
				},
				{
					Name:   "import",
					Usage:  "import the Zaha! app databases",
					Action: db_import,
				},
			},
		},
		{
			Name:   "create",
			Usage:  "create a Zaha! app",
			Action: create_app,
		},
		{
			Name:   "start",
			Usage:  "start a Zaha! app or component; locally or in the cloud",
			Action: start_app,
		},
		{
			Name:   "stop",
			Usage:  "stop a Zaha! app or component; locally or in the cloud",
			Action: stop_app,
		},
		{
			Name:   "restart",
			Usage:  "restart a Zaha! app or component; locally or in the cloud",
			Action: restart_app,
		},
		{
			Name:   "status",
			Usage:  "status of a Zaha! app or component;, locally or in the cloud",
			Action: status_app,
		},
		{
			Name:   "push",
			Usage:  "push a Zaha! app or component;, locally or into the cloud",
			Action: push_app,
		},
		{
			Name:   "dev",
			Usage:  "dev...?",
			Action: dev,
		},
	}
)
