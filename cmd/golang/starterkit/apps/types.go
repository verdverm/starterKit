package apps

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/fsouza/go-dockerclient"
	"gopkg.in/yaml.v2"
)

type AppConfig struct {
	Name         string
	Description  string
	SourcePath   string
	StaticPath   string
	NotebookPath string

	Frontends []string
	Services  []string
	Databases []string
	Backends  []string
}

type DockerOptions struct {
	Name       string
	Component  string
	UseStorage bool

	Cpus      float64
	Mem       float64
	Instances int

	Image string
	Ports []string
	Binds []string
	Env   []string

	CreateOpts docker.CreateContainerOptions
	StartOpts  *docker.HostConfig

	// Other dockers this one depends on and links to
	Deps []string
}

func ReadDockerOpts(component_name string, acfg AppConfig) (DockerOptions, error) {
	var opts DockerOptions

	filename := acfg.SourcePath + fmt.Sprintf("/dockers/%s.yaml", component_name)
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return opts, err
	}
	err = yaml.Unmarshal([]byte(data), &opts)
	if err != nil {
		return opts, err
	}

	return opts, nil
}

func PrintDockerOpts(opts *DockerOptions) error {
	data, err := yaml.Marshal(opts)
	if err != nil {
		return err
	}
	fmt.Printf(string(data))
	return nil
}

func WriteDockerOpts(opts *DockerOptions, acfg AppConfig) error {
	filename := acfg.SourcePath + fmt.Sprintf("/dockers/%s.yaml", opts.Component)

	d, err := yaml.Marshal(opts)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filename, d, 0644)
	if err != nil {
		return err
	}
	// fmt.Println("saved Zaha docker opts")
	return nil
}

func StartDocker(opts *DockerOptions) error {
	name := opts.Name
	client, err := docker.NewClient(DOCKER_HOST)
	if err != nil {
		return err
	}
	_, err = client.CreateContainer(opts.CreateOpts)
	if err != nil {
		return err
	}
	err = client.StartContainer(name, opts.StartOpts)
	if err != nil {
		return err
	}
	return nil
}

func StopDocker(opts *DockerOptions) error {
	name := opts.Name

	client, err := docker.NewClient(DOCKER_HOST)
	if err != nil {
		return err
	}

	nc_str := "No such container:"
	err = client.KillContainer(docker.KillContainerOptions{ID: name})
	if err != nil {
		if strings.Contains(err.Error(), nc_str) {
			return errors.New("non-existant")
		} else {
			return err
		}
	}

	err = RemoveDocker(name)
	if err != nil {
		return err
	}

	return nil
}

func StatusDocker(opts *DockerOptions) error {
	name := opts.Name

	client, err := docker.NewClient(DOCKER_HOST)
	if err != nil {
		return err
	}

	state := "running"
	nc_str := "No such container:"
	dock, err := client.InspectContainer(name)
	if err != nil {
		if strings.Contains(err.Error(), nc_str) {
			state = "non-existant"
		} else {
			return err
		}
	} else if !dock.State.Running {
		state = fmt.Sprintf("%v %v", dock.State.ExitCode, dock.State.Error)
	}

	status := fmt.Sprintf(container_status_short_fmt,
		name,
		state,
	)

	out := status

	fmt.Println(out)
	return nil
}

const container_status_short_fmt = `  %-20s  %s`

func RemoveDocker(name string) error {
	client, err := docker.NewClient(DOCKER_HOST)
	if err != nil {
		return err
	}

	// remove options
	ropts := docker.RemoveContainerOptions{
		ID:            name,
		RemoveVolumes: false,
		Force:         true,
	}
	err = client.RemoveContainer(ropts)
	if err != nil {
		return err
	}
	return nil
}
