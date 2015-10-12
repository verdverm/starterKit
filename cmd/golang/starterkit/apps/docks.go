package apps

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

var (
	available = []string{"psql", "neo4j", "couchdb", "redis", "rabbitmq", "ipynb", "flask", "nginx"}
	storage   = []string{"psql", "neo4j", "couchdb", "redis", "rabbitmq", "ipynb", "flask", "nginx"}

	DOCKER_HOST = "unix:///var/run/docker.sock"
)

func GetAppDockerList(arg string, acfg AppConfig) (dock_list []string, err error) {
	switch arg {
	case "app":
		// XXX  This case is the only place where docker init order is defined anywhere XXX
		// well the only one that matters as of this writting
		// the other places that order comes out is in the yaml files and in main.go var(...) section
		dock_list = append(dock_list, acfg.Frontends...)
		dock_list = append(dock_list, acfg.Services...)
		dock_list = append(dock_list, acfg.Backends...)
		dock_list = append(dock_list, acfg.Databases...)
	case "fe":
		dock_list = acfg.Frontends
	case "db":
		dock_list = acfg.Databases
	case "be":
		dock_list = acfg.Backends
	case "services":
		dock_list = acfg.Services
	default:
		// look for arg in available dockers
		found := false
		for _, d := range available {
			if d == arg {
				found = true
				break
			}
		}
		if !found {
			text := fmt.Sprintf("invalid arguement: '%s'\n", arg)
			err = errors.New(text)
		} else {
			dock_list = append(dock_list, arg)
		}

	}
	return dock_list, err
}

// this is used when starting right now
// should really do a topo sort, but need to inject deps in both internal code and user configs
func ReverseDockerList(list []*DockerOptions) {
	i := 0
	L := len(list) - 1
	for i < L {
		list[i], list[L] = list[L], list[i]
		i++
		L--
	}
}

func PrepDocker(opts *DockerOptions, acfg AppConfig) (err error) {

	err = environPrepper(opts, acfg)
	if err != nil {
		return err
	}

	if opts.UseStorage {
		err = storagePrepper(opts, acfg)
		if err != nil {
			return err
		}
	}

	prepper, ok := local_preppers[opts.Component]
	if ok {
		err = prepper(opts, acfg)
		if err != nil {
			return err
		}
	}

	return nil
}

type dockPrepper func(opts *DockerOptions, acfg AppConfig) error

var (
	local_preppers = map[string]dockPrepper{
		"storage": storagePrepper,
		"flask":   flaskPrepper,
		"ipynb":   ipynbPrepper,
		"nginx":   nginxPrepper,
	}
)

func storagePrepper(opts *DockerOptions, acfg AppConfig) error {
	fmt.Println("APPCFG SRCPATH: ", acfg.SourcePath)
	storage_dir := acfg.SourcePath + "/storage/" + opts.Component + "/"
	binds := opts.StartOpts.Binds
	for i, b := range binds {
		if b[0] != '/' {
			binds[i] = storage_dir + b
		}
	}
	opts.Binds = binds[:]
	return nil
}

func environPrepper(opts *DockerOptions, acfg AppConfig) error {
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

func flaskPrepper(opts *DockerOptions, acfg AppConfig) error {
	opts.StartOpts.Links = []string{
		"psql." + acfg.Name + ":db_psql",
		"couchdb." + acfg.Name + ":db_couchdb",
		"redis." + acfg.Name + ":redis",
	}
	appsrc_bind := acfg.SourcePath + "/webapp:/src"
	opts.StartOpts.Binds = append(opts.StartOpts.Binds, appsrc_bind)
	nbsrc_bind := acfg.NotebookPath + ":/notebooks"
	opts.StartOpts.Binds = append(opts.StartOpts.Binds, nbsrc_bind)
	return nil
}

func ipynbPrepper(opts *DockerOptions, acfg AppConfig) error {
	notebooks_bind := acfg.NotebookPath + ":/ipynb/notebooks"
	opts.StartOpts.Binds = append(opts.StartOpts.Binds, notebooks_bind)
	// printDockerOpts(opts)
	return nil
}

func nginxPrepper(opts *DockerOptions, acfg AppConfig) error {
	static_bind := acfg.StaticPath + ":/static"
	certs_bind := "/etc/ssl:/etc/ssl"
	nbsrc_bind := acfg.NotebookPath + ":/static/notebooks"
	opts.StartOpts.Binds = append(opts.StartOpts.Binds, static_bind)
	opts.StartOpts.Binds = append(opts.StartOpts.Binds, certs_bind)
	opts.StartOpts.Binds = append(opts.StartOpts.Binds, nbsrc_bind)
	return nil
}
