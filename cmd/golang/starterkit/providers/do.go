package providers

import (
	// "bytes"
	"fmt"
	// "io/ioutil"
	"os"
	// "os/user"
	"strings"
	"time"

	"code.google.com/p/goauth2/oauth"
	"github.com/digitalocean/godo"

	"github.com/zaha-io/zaha/apps"
	"github.com/zaha-io/zaha/mesos"
)

var (
	do_pat = os.Getenv("DIGITAL_OCEAN_TOKEN")
)

func newOceanClient() *godo.Client {
	t := &oauth.Transport{
		Token: &oauth.Token{AccessToken: do_pat},
	}
	return godo.NewClient(t.Client())
}

func oceanCloudCreate(ccfg CloudConfig) error {
	return nil
}

func oceanCloudStatus(ccfg CloudConfig) error {
	fmt.Println("  Mesos:")

	url := ccfg.URI + ":5050"

	reg, err := mesos.GetMesosRegistrar(url)
	if err != nil {
		fmt.Println("while getting mesos registrar", err)
		return err
	}
	fmt.Printf("\n%+v\n\n", reg)

	stats, err := mesos.GetMesosStats(url)
	if err != nil {
		fmt.Println("while getting mesos registrar", err)
		return err
	}
	fmt.Printf("\n%+v\n\n", stats)

	metrics, err := mesos.GetMesosMetrics(url)
	if err != nil {
		fmt.Println("while getting mesos registrar", err)
		return err
	}
	fmt.Printf("\n%+v\n\n", metrics)

	// client, err := newMarathonClient(":8080")
	// if err != nil {
	// 	return err
	// }

	// // List all apps
	// r, err := client.ListApps()
	// if err != nil {
	// 	return err
	// }
	// v, _ := json.Marshal(r)
	// fmt.Printf("%s", v)

	return nil

	return nil
}
func oceanCloudResize(ccfg CloudConfig) error {
	return nil
}
func oceanCloudDestroy(ccfg CloudConfig) error {

	client := newOceanClient()

	droplets, err := getDropletList(client)
	if err != nil {
		return err
	}

	for i, d := range droplets {
		if strings.Contains(d.Name, ccfg.Name) {
			fmt.Printf("D%02d:   %v   %v   %v\n", i, "deleting", d.Name, d.ID)
			deleteCloudNode(d.ID, client)
		}
	}
	return nil
}
func oceanCloudPushApp(ccfg CloudConfig, acfg apps.AppConfig) error {
	return nil
}

func createOceanNode(owner_name, drop_name, cloudProvider string) (*godo.Droplet, error) {
	client := newOceanClient()

	drop, err := createDroplet(drop_name, client)
	if err != nil {
		return nil, err
	}

	// need to for ssh to come up
	time.Sleep(time.Second * 6)
	sleepiness := 1
	cnt := 0
	for cnt < 30 {
		drop_chk, _, err := client.Droplets.Get(drop.ID)
		if err != nil {
			return nil, err
		}

		if drop_chk.Droplet.Status == "active" {
			if len(drop_chk.Droplet.Networks.V4) > 0 {
				drop = drop_chk.Droplet
				break
			} else {

				fmt.Printf("Droplet ID(%02d): XXX  %d   %s\n", cnt, drop_chk.Droplet.ID, drop_chk.Droplet.Status)
			}
		} else {
			fmt.Printf("Droplet ID(%02d):      %d   %s\n", cnt, drop_chk.Droplet.ID, drop_chk.Droplet.Status)
		}

		time.Sleep(time.Second * 3 * time.Duration(sleepiness))
		cnt++
		if cnt == 30 && sleepiness == 1 {
			cnt = 0
			sleepiness = 2
		}
	}

	// do the actual setup on a new node
	ip := drop.Networks.V4[0].IPAddress // still an issue
	execCommand("/bin/bash", "/home/tony/.zaha/cloud/scripts/start-mesos-node.sh", ip, owner_name)

	if err != nil {
		return nil, err
	}

	return drop, nil
}

func createDroplet(name string, client *godo.Client) (*godo.Droplet, error) {
	ssh_key_id := 326841
	ints := make([]interface{}, 1)
	ints[0] = interface{}(ssh_key_id)

	createRequest := &godo.DropletCreateRequest{
		Name:    name,
		Region:  "nyc3",
		Size:    "2gb",
		Image:   "mesos-boot-image",
		SSHKeys: []interface{}(ints),
	}

	newDroplet, _, err := client.Droplets.Create(createRequest)
	if err != nil {
		fmt.Printf("Something bad happened: %s\n\n", err)
		return nil, err
	}

	fmt.Printf("New Droplet: %s\n%+v\n\n", name, newDroplet)

	return newDroplet.Droplet, nil
}

func deleteCloudNode(dId int, client *godo.Client) error {
	fmt.Println("deleteing: ", dId)
	_, err := client.Droplets.Delete(dId)

	if err != nil {
		fmt.Printf("Something bad happened: %s\n\n", err)
		return err
	}
	return nil
}

func getDropletList(client *godo.Client) ([]godo.Droplet, error) {
	// create a list to hold our droplets
	list := []godo.Droplet{}

	// create options. initially, these will be blank
	opt := &godo.ListOptions{}
	for {
		droplets, resp, err := client.Droplets.List(opt)
		if err != nil {
			return nil, err
		}

		// append the current page's droplets to our list
		for _, d := range droplets {
			list = append(list, d)
		}

		// if we are at the last page, break out the for loop
		if resp.Links.IsLastPage() {
			break
		}

		page, err := resp.Links.CurrentPage()
		if err != nil {
			return nil, err
		}

		// set the page we want for the next request
		opt.Page = page + 1
	}

	return list, nil
}

func startMesosStack(droplet *godo.Droplet, id int, role, name, master_ips string) error {

	fmt.Println("\n\n\n\nstarting mesos on ", name, "\n\n\n")

	// // rename droplet through DO
	// t := &oauth.Transport{
	// 	Token: &oauth.Token{AccessToken: do_pat},
	// }

	// client := godo.NewClient(t.Client())
	// _, _, err := client.DropletActions.Rename(id, name)
	// if err != nil {
	// 	return err
	// }

	ip := droplet.Networks.V4[0].IPAddress
	id_str := fmt.Sprint(id)
	execCommand("/bin/bash", "/home/tony/.zaha/cloud/scripts/start-mesos-stack.sh", id_str, ip, role, name, master_ips)

	return nil
}
