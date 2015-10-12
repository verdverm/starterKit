package mesos

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/zaha-io/gomarathon"

	"github.com/zaha-io/zaha/apps"
)

func newMarathonClient(host string) (*gomarathon.Client, error) {
	return gomarathon.NewClient(host, nil)
}

func GetMesosRegistrar(host string) (mesos MesosRegistrar, err error) {
	url := "http://" + host + "/registrar(1)/registry"
	resp, err := http.Get(url)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	r, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	// fmt.Println(string(r))

	err = json.Unmarshal(r, &mesos)
	if err != nil {
		return
	}

	return
}

func GetMesosStats(host string) (mesos MesosStats, err error) {
	url := "http://" + host + "/system/stats.json"
	resp, err := http.Get(url)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	r, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	// fmt.Println(string(r))

	err = json.Unmarshal(r, &mesos)
	if err != nil {
		return
	}

	return
}

func GetMesosMetrics(host string) (mesos MesosMetrics, err error) {
	url := "http://" + host + "/metrics/snapshot"
	resp, err := http.Get(url)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	r, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	r2 := bytes.Replace(r, []byte("\\/"), []byte("__"), -1)

	// fmt.Println(string(r2))

	err = json.Unmarshal(r2, &mesos)
	if err != nil {
		return
	}

	return
}

func GetMesosphereApps(host string) (app gomarathon.Application, err error) {
	url := "http://" + host + "/v2/tasks"
	resp, err := http.Get(url)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	r, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	r2 := bytes.Replace(r, []byte("\\/"), []byte("__"), -1)

	fmt.Println(string(r2))

	err = json.Unmarshal(r2, &app)
	if err != nil {
		return
	}

	return
}

func StatusMesosphereApp(acfg apps.AppConfig, cloudName, cloudURI string) error {

	return nil
}

func StatusMesosphereDockers(docks_opts []*apps.DockerOptions, acfg apps.AppConfig, cloudName, cloudURI string) error {

	return nil
}

func StatusMesosphereTasks(acfg apps.AppConfig, cloudName, cloudURI string) (appTasks MesosAppTaskStruct, err error) {

	host := cloudURI
	url := "http://" + host + ":8080/v2/tasks"
	fmt.Printf("    %-16s (%s)\n", cloudName, url)

	resp, err := http.Get(url)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	r, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	r2 := bytes.Replace(r, []byte("\\/"), []byte("__"), -1)

	// fmt.Println(string(r2))

	err = json.Unmarshal(r2, &appTasks)
	if err != nil {
		return
	}

	return

}

func StartMesosphereApp(acfg apps.AppConfig, cloudName, cloudURI string) error {
	// func StartMesosphereApp(app gomarathon.Application, host string) error {

	dock_list, err := apps.GetAppDockerList("app", acfg)
	if err != nil {
		return err
	}

	var dock_opts []*apps.DockerOptions
	for _, dock_name := range dock_list {

		dopts, err := apps.ReadDockerOpts(dock_name, acfg)
		if err != nil {
			return err
		}

		fmt.Printf("%+v\n\n", dopts)
		dock_opts = append(dock_opts, &dopts)
	}

	apps.ReverseDockerList(dock_opts)

	for _, dopts := range dock_opts {
		err = StartMesosphereDocker(dopts, acfg, cloudName, cloudURI)
		if err != nil {
			return err
		}
	}

	return nil
}

func StartMesosphereDockers(docks_opts []*apps.DockerOptions, acfg apps.AppConfig, cloudName, cloudURI string) error {
	for _, d := range docks_opts {
		err := StartMesosphereDocker(d, acfg, cloudName, cloudURI)
		if err != nil {
			return err
		}
	}
	return nil
}

func StartMesosphereDocker(dopts *apps.DockerOptions, acfg apps.AppConfig, cloudName, cloudURI string) error {
	// need to prep dockers here
	err := PrepMesosphereDocker(dopts, acfg, cloudName)
	if err != nil {
		return err
	}

	host := cloudURI
	url := "http://" + host + ":8080/v2/apps"
	fmt.Printf("  %16s   ->  %s\n", dopts.Name, url)

	var app MesosphereContainerStruct

	app.Id = dopts.Name
	app.Container.Type = "DOCKER"
	app.Container.Docker.Network = "BRIDGE"
	app.Container.Docker.Image = dopts.Image

	app.Cpus = dopts.Cpus
	app.Mem = dopts.Mem
	app.Instances = dopts.Instances
	app.UpgradeStrategy = UpgradeStrategyStruct{1.0}

	// deal with ports
	for _, port := range dopts.Ports {
		cp, hp := port, "0"
		if strings.Contains(port, ":") {
			flds := strings.Split(port, ":")
			if len(flds) != 2 {
				return errors.New("bad port spec for: " + dopts.Name + " ... " + port)
			}
			cp, hp = flds[0], flds[1]
		}
		fmt.Println("ports:", port, cp, hp)
		cpi, err := strconv.ParseInt(cp, 10, 64)
		if err != nil {
			return err
		}
		hpi, err := strconv.ParseInt(hp, 10, 64)
		if err != nil {
			return err
		}
		app.Ports = append(app.Ports, int(cpi))
		pm := PortMappingsStruct{int(cpi), int(hpi), "tcp", 0}
		app.Container.Docker.PortMappings = append(app.Container.Docker.PortMappings, pm)
	}

	// deal with mounted volumes
	for _, bind := range dopts.Binds {
		flds := strings.Split(bind, ":")
		if len(flds) != 2 && flds[0] != "" && flds[1] != "" {
			return errors.New("bad bind spec for: " + dopts.Name + " ... " + bind)
		}
		hpath, cpath := flds[0], flds[1]
		v := VolumesStruct{cpath, hpath, "RW"}
		app.Container.Volumes = append(app.Container.Volumes, v)
	}

	app.Env = make(map[string]string)
	for _, env := range dopts.Env {
		flds := strings.Split(env, "=")
		if len(flds) != 2 && flds[0] != "" && flds[1] != "" {
			return errors.New("bad env var spec for: " + dopts.Name + " ... " + env)
		}

		key := flds[0]
		value := flds[1]
		if strings.Contains(value, "$") {
			value = os.Getenv(value[1:])
		}
		app.Env[key] = value
	}

	jsonBytes, err := json.MarshalIndent(app, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(jsonBytes))

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonBytes))

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error posting json:", err)
	}
	fmt.Println(string(body))

	return nil
}

func StopMesosphereDocker(host, dock_id string) error {
	url := "http://" + host + ":8080/v2/apps/" + dock_id

	fmt.Println("DEL >", url)

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		// fmt.Println("response Status:", resp.Status)
		// fmt.Println("response Headers:", resp.Header)
		// fmt.Println("response Body:", string(body))
		return errors.New("error: " + resp.Status)
	}
	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	return err
	// }
	return nil
}

func PrepMesosphereDocker(opts *apps.DockerOptions, acfg apps.AppConfig, cloudName string) (err error) {

	err = environPrepper(opts, acfg, cloudName)
	if err != nil {
		return err
	}

	if opts.UseStorage {
		err = storagePrepper(opts, acfg, cloudName)
		if err != nil {
			return err
		}
	}

	prepper, ok := local_preppers[opts.Component]
	if ok {
		err = prepper(opts, acfg, cloudName)
		if err != nil {
			return err
		}
	}

	return nil
}

type dockPrepper func(opts *apps.DockerOptions, acfg apps.AppConfig, cloudName string) error

var (
	local_preppers = map[string]dockPrepper{
		"storage": storagePrepper,
		"flask":   flaskPrepper,
		"ipynb":   ipynbPrepper,
		"nginx":   nginxPrepper,
	}
)

func storagePrepper(opts *apps.DockerOptions, acfg apps.AppConfig, cloudName string) error {
	storage_dir := fmt.Sprintf("/storage/%s/%s/%s/", cloudName, acfg.Name, opts.Component)
	binds := opts.StartOpts.Binds
	for i, b := range binds {
		if b[0] != '/' {
			binds[i] = storage_dir + b
		}
	}
	opts.Binds = binds[:]
	return nil
}

func environPrepper(opts *apps.DockerOptions, acfg apps.AppConfig, cloudName string) error {
	envs := opts.CreateOpts.Config.Env
	for i, e := range envs {
		flds := strings.Split(e, "$")
		if len(flds) == 2 {
			val := os.Getenv(flds[1])
			if val == "" {
				return errors.New("ENV var: " + flds[1] + " not set for use in " + opts.Name)
			}
			envs[i] = flds[0] + val
		}
	}
	return nil
}

func flaskPrepper(opts *apps.DockerOptions, acfg apps.AppConfig, cloudName string) error {
	opts.StartOpts.Links = []string{
		"psql." + acfg.Name + ":db_psql",
		"couchdb." + acfg.Name + ":db_couchdb",
		"redis." + acfg.Name + ":redis",
	}
	appsrc_bind := fmt.Sprintf("/storage/%s/%s/webapp:/src", cloudName, acfg.Name, opts.Component)
	opts.StartOpts.Binds = append(opts.StartOpts.Binds, appsrc_bind)
	return nil
}

func ipynbPrepper(opts *apps.DockerOptions, acfg apps.AppConfig, cloudName string) error {
	notebooks_bind := fmt.Sprintf("/storage/%s/%s/notebooks:/ipynb/notebooks", cloudName, acfg.Name, opts.Component)
	opts.StartOpts.Binds = append(opts.StartOpts.Binds, notebooks_bind)
	// printDockerOpts(opts)
	return nil
}

func nginxPrepper(opts *apps.DockerOptions, acfg apps.AppConfig, cloudName string) error {
	static_bind := fmt.Sprintf("/storage/%s/%s/static:/static", cloudName, acfg.Name, opts.Component)
	certs_bind := "/etc/ssl:/etc/ssl"
	opts.StartOpts.Binds = append(opts.StartOpts.Binds, static_bind)
	opts.StartOpts.Binds = append(opts.StartOpts.Binds, certs_bind)
	return nil
}
