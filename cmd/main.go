package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/codegangsta/cli"

	"github.com/zaha-io/zaha/apps"
)

var (
	do_pat = os.Getenv("DIGITAL_OCEAN_TOKEN")

	app = cli.NewApp()

	// YAML This
	commands = []string{"start", "status", "stop", "restart"}
	subcmds  = []string{"app", "db", "fe", "be", "services"}
	db_extra = []string{"load", "export"}
	// End YAML

	// ZAHAMETA ZahaMeta
	ZAHACONFIG ZahaConfig
	APPCFG     apps.AppConfig

	// Put these somewhere else
	config_dir    string
	config_fn     string
	config_fn_abs string
)

func init() {
	app.Name = "zaha"
	app.Version = "0.2.3"
	app.Usage = "named after a great architect"
	app.Flags = []cli.Flag{}
	app.Commands = menu
	// app.Action = zaha_app

	usr, err := user.Current()
	if err != nil {
		panicOnError(err, "unable to get current user")
	}
	config_dir = usr.HomeDir + "/.zaha/"
	config_fn = "zaha.yaml"
	config_fn_abs = config_dir + config_fn

	ZAHACONFIG, err = readZahaConfig()
	if err != nil {
		fmt.Errorf("couldn't read zaha config: %v\n", err)
		initZahaConfig()
		if err != nil {
			fmt.Errorf("couldn't init zaha config: %v\n", err)
		}
	}
	ZAHACONFIG.ConfigAbs = config_dir
	APPCFG, err = readAppConfig(ZAHACONFIG.CurrentApp)
	if err != nil {
		fmt.Println("error reading current app '" + ZAHACONFIG.CurrentApp + "' config")
	}
}

func main() {
	app.Run(os.Args)
}

func setupZaha(c *cli.Context) {
	if len(c.Args()) > 0 {
		fmt.Println("Error: bad args to 'zaha config setup'")
		cli.ShowSubcommandHelp(c)
		return
	}

	initZahaConfig()

	// pull dockers

}

func dev(c *cli.Context) {
	if len(c.Args()) != 1 {
		fmt.Println("Error: bad args to 'dev' !!!")
		cli.ShowSubcommandHelp(c)
		return
	}
	fmt.Println("Not Implemented")

}
