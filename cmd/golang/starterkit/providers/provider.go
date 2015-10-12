package providers

import (
	"errors"
	"fmt"
	"strings"

	"github.com/zaha-io/zaha/apps"
)

type CloudConfig struct {
	Name     string
	Provider string
	URI      string

	NumMaster int
	NumWorker int

	Status string
}

type ServiceConfig struct {
	Name     string
	Replicas int

	Status string
}

func CloudCreate(ccfg CloudConfig) error {
	fmt.Println("Cloud Create:")
	cn := strings.ToLower(ccfg.Provider)
	switch cn {

	case "local", "localhost":
		return localCloudCreate(ccfg)

	case "do", "ocean", "digitalocean":
		return oceanCloudCreate(ccfg)

	case "aws", "amazon":
		return amazonCloudCreate(ccfg)

	case "gce", "google":
		return googleCloudCreate(ccfg)

	default:
		return errors.New("Unknown cloud provider: " + cn)
	}

	return nil
}
func CloudStatus(ccfg CloudConfig) error {
	fmt.Println("Cloud Status:")
	cn := strings.ToLower(ccfg.Provider)
	switch cn {

	case "local", "localhost":
		return localCloudStatus(ccfg)

	case "do", "ocean", "digitalocean":
		return oceanCloudStatus(ccfg)

	case "aws", "amazon":
		return amazonCloudStatus(ccfg)

	case "gce", "google":
		return googleCloudStatus(ccfg)

	default:
		return errors.New("Unknown cloud provider: " + cn)
	}

	return nil
}
func CloudResize(ccfg CloudConfig) error {
	fmt.Println("Cloud Resize:")
	cn := strings.ToLower(ccfg.Provider)
	switch cn {

	case "local", "localhost":
		return localCloudResize(ccfg)

	case "do", "ocean", "digitalocean":
		return oceanCloudResize(ccfg)

	case "aws", "amazon":
		return amazonCloudResize(ccfg)

	case "gce", "google":
		return googleCloudResize(ccfg)

	default:
		return errors.New("Unknown cloud provider: " + cn)
	}

	return nil
}
func CloudDestroy(ccfg CloudConfig) error {
	fmt.Println("Cloud Destroy:")
	cn := strings.ToLower(ccfg.Provider)
	switch cn {

	case "local", "localhost":
		return localCloudDestroy(ccfg)

	case "do", "ocean", "digitalocean":
		return oceanCloudDestroy(ccfg)

	case "aws", "amazon":
		return amazonCloudDestroy(ccfg)

	case "gce", "google":
		return googleCloudDestroy(ccfg)

	default:
		return errors.New("Unknown cloud provider: " + cn)
	}

	return nil
}

func CloudPushApp(ccfg CloudConfig, acfg apps.AppConfig) error {
	// fmt.Println("Cloud PushApp: ", ccfg.Name, acfg.Name)
	cn := strings.ToLower(ccfg.Provider)
	switch cn {

	case "local", "localhost":
		return localCloudPushApp(ccfg, acfg)

	case "do", "ocean", "digitalocean":
		return oceanCloudPushApp(ccfg, acfg)

	case "aws", "amazon":
		return amazonCloudPushApp(ccfg, acfg)

	case "gce", "google":
		return googleCloudPushApp(ccfg, acfg)

	default:
		return errors.New("Unknown cloud provider: " + cn)
	}

	return nil
}
