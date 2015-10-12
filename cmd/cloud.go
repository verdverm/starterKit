package main

import (
	"fmt"
	"io/ioutil"
	// "strconv"
	"strings"
	// "sync"

	"github.com/codegangsta/cli"

	"github.com/zaha-io/zaha/providers"
)

func cloud_create(c *cli.Context) {
	args := c.Args()
	L := len(args)
	if L < 2 || L > 3 {
		fmt.Println("Error: bad args to 'cloud create <provider> <cloudname>' !!!")
		cli.ShowSubcommandHelp(c)
		return
	}

	// appName := ZAHACONFIG.CurrentApp
	// ownerName := ZAHACONFIG.CurrentUser
	// cloudName := args[0]
	// dropPrefix := appName + "-" + cloudName

	// provider := args[1] // cloud provider
	// numW, err := strconv.ParseInt(args[2], 10, 64)
	// if err != nil {
	// 	fmt.Println("while parsing numWorkers int: ", err)
	// }

	// numMaster := 1
	// numWorker := int(numW)
	// if numWorker > 5 {
	// 	numMaster = 3
	// } // another tier?
	// num_nodes := numMaster + numWorker

	// var node_names []string
	// for i := 0; i < num_nodes; i++ {
	// 	name := fmt.Sprintf(dropPrefix+"-node-%06d", i+1)
	// 	node_names = append(node_names, name)
	// }

	// drop_chan := make(chan *godo.Droplet, 16)
	// // start up nodes for all roles
	// var wg1 sync.WaitGroup
	// for _, name := range node_names {
	// 	wg1.Add(1)

	// 	go func(nn string) {
	// 		defer wg1.Done()
	// 		drop, err := CreateCloudNode(nn, provider)
	// 		if err != nil {
	// 			fmt.Println("Error setting up", nn, ":  ", err)
	// 			return
	// 		} else {
	// 			drop_chan <- drop
	// 		}

	// 	}(name)
	// }

	// // wait for first M nodes completed and set them up as masters
	// master_drops := make([]*godo.Droplet, numMaster)
	// master_ips := ""
	// for i := 0; i < numMaster; i++ {
	// 	drop := <-drop_chan
	// 	master_drops[i] = drop
	// 	master_ips += " " + drop.Networks.V4[0].IPAddress
	// }
	// // remove first comma
	// master_ips = master_ips[1:]

	// var wg2 sync.WaitGroup
	// for i := 0; i < numMaster; i++ {
	// 	wg2.Add(1)
	// 	drop := master_drops[i]

	// 	name := dropPrefix + fmt.Sprintf("-master-%d", i+1)

	// 	go func(D *godo.Droplet, ID int, R, N, MIP string) {
	// 		defer wg2.Done()

	// 		err := StartMesosStack(D, ID, R, N, MIP)
	// 		if err != nil {
	// 			fmt.Println("Error setting up", N, ":  ", err)
	// 		}
	// 	}(drop, i+1, "master", name, master_ips)
	// }

	// // setup the remaining nodes
	// for i := 0; i < numWorker; i++ {
	// 	drop := <-drop_chan
	// 	wg2.Add(1)

	// 	name := dropPrefix + fmt.Sprintf("-worker-%06d", i+1)

	// 	go func(D *godo.Droplet, ID int, R, N, MIP string) {
	// 		defer wg2.Done()

	// 		err := StartMesosStack(D, ID, R, N, MIP)
	// 		if err != nil {
	// 			fmt.Println("Error setting up", N, ":  ", err)
	// 		}
	// 	}(drop, i+1, "worker", name, master_ips)
	// }

	// wg1.Wait()
	// wg2.Wait()

	// fmt.Println("\n\n\n\nZaha!   your cluster is ready\n\n\n")

}

func cloud_destroy(c *cli.Context) {
	if len(c.Args()) != 1 {
		fmt.Println("Error: bad args to 'cloud destroy'")
		cli.ShowSubcommandHelp(c)
		return
	}

	fmt.Println("GOT HERE")

	clustername := c.Args()[0]

	ccfg, err := readCloudConfig(APPCFG.Name, clustername)
	if err != nil {
		fmt.Println("error reading cloud config for: ", clustername)
		return
	}

	fmt.Println("Not quite destroying:", ccfg.Name)

}

func cloud_list(c *cli.Context) {
	args := c.Args()
	L := len(args)
	if L != 0 {
		fmt.Println("Error: bad args to 'cloud list'")
		cli.ShowSubcommandHelp(c)
		return
	}
	appname := APPCFG.Name

	// get a directory listing
	finfos, err := ioutil.ReadDir(APPCFG.SourcePath + "/clouds")
	if err != nil {
		fmt.Println("while reading clouds directory: ", err)
	}

	fmt.Printf("\n%-16s %-15s %-12s   %s\n", "Name", "Host", "Status", "Master:Worker counts")
	for _, fi := range finfos {
		if fi.IsDir() {
			continue
		}
		cloudname := strings.Replace(fi.Name(), ".yaml", "", -1)
		cloudcfg, err := readCloudConfig(appname, cloudname)
		if err != nil {
			fmt.Println("while reading cloud config: ", err)
		}

		fmt.Printf("%-16s %-15s %-12s   %d:%d\n", cloudcfg.Name, cloudcfg.URI, cloudcfg.Status, cloudcfg.NumMaster, cloudcfg.NumWorker)
	}
}

func cloud_status(c *cli.Context) {
	args := c.Args()
	L := len(args)
	if L != 1 {
		fmt.Println("Error: bad args to 'cloud status <cloudname>'")
		cli.ShowSubcommandHelp(c)
		return
	}

	appname := APPCFG.Name
	cloudname := args[0]

	cloudcfg, err := readCloudConfig(appname, cloudname)
	if err != nil {
		fmt.Println("while reading cloud config: ", err)
		return
	}

	// print status known in the config file
	fmt.Printf("\n%-16s %-15s %-12s   %s\n", "Name", "Host", "Status", "Master:Worker counts")
	fmt.Printf("%-16s %-15s %-12s   %d:%d\n\n", cloudcfg.Name, cloudcfg.URI, cloudcfg.Status, cloudcfg.NumMaster, cloudcfg.NumWorker)

	// print status from mesos

	providers.CloudStatus(cloudcfg)

	// print status for the app

	// print status from mesosphere ? or is the service status

}

func cloud_resize(c *cli.Context) {
	args := c.Args()
	L := len(args)
	if L != 1 {
		fmt.Println("Error: bad args to 'cloud status <cloudname>'")
		cli.ShowSubcommandHelp(c)
		return
	}

	appname := APPCFG.Name
	cloudname := args[0]

	cloudcfg, err := readCloudConfig(appname, cloudname)
	if err != nil {
		fmt.Println("while reading cloud config: ", err)
	}

	fmt.Printf("\n%-16s %-15s %-12s   %s\n", "Name", "Host", "Status", "Master:Worker counts")
	fmt.Printf("%-16s %-15s %-12s   %d:%d\n", cloudcfg.Name, cloudcfg.URI, cloudcfg.Status, cloudcfg.NumMaster, cloudcfg.NumWorker)

}
