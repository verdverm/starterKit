package providers

import (
	// "bytes"
	"fmt"
	// "io/ioutil"
	// "os"
	// "os/user"
	// "time"
	// "encoding/json"

	// "github.com/zaha-io/gomarathon"

	"github.com/zaha-io/zaha/apps"
	"github.com/zaha-io/zaha/mesos"
)

func localCloudCreate(ccfg CloudConfig) error {
	return nil
}
func localCloudStatus(ccfg CloudConfig) error {
	fmt.Println("  Mesos:")

	reg, err := mesos.GetMesosRegistrar("localhost:5050")
	if err != nil {
		fmt.Println("while getting mesos registrar", err)
		return err
	}
	fmt.Printf("\n%+v\n\n", reg)

	stats, err := mesos.GetMesosStats("localhost:5050")
	if err != nil {
		fmt.Println("while getting mesos registrar", err)
		return err
	}
	fmt.Printf("\n%+v\n\n", stats)

	// metrics, err := GetMesosMetrics("localhost:5050")
	// if err != nil {
	// 	fmt.Println("while getting mesos registrar", err)
	// 	return err
	// }
	// fmt.Printf("\n%+v\n\n", metrics)

	return nil
}
func localCloudResize(ccfg CloudConfig) error {
	return nil
}
func localCloudDestroy(ccfg CloudConfig) error {

	return nil
}
func localCloudPushApp(ccfg CloudConfig, acfg apps.AppConfig) error {
	fmt.Println()
	// fmt.Println("Packaging files")

	fmt.Println("Copying files")

	sourceDirFrom := acfg.SourcePath + "/webapp/*"
	staticDirFrom := acfg.StaticPath + "/*"
	notebookDirFrom := acfg.NotebookPath + "/*"

	rootDirTo := "/storage/" + ccfg.Name + "/" + acfg.Name

	sourceDirTo := rootDirTo + "/webapp"
	staticDirTo := rootDirTo + "/static"
	notebookDirTo := rootDirTo + "/notebooks"

	fmt.Println("  cp", sourceDirFrom, sourceDirTo)
	fmt.Println("  cp", staticDirFrom, staticDirTo)
	fmt.Println("  cp", notebookDirFrom, notebookDirTo)

	execCommand("sudo", "mkdir", "-p", sourceDirTo)
	execCommand("sudo", "mkdir", "-p", staticDirTo)
	execCommand("sudo", "mkdir", "-p", notebookDirTo)

	execCommand("bash", "-c", "sudo cp -R "+sourceDirFrom+" "+sourceDirTo)
	execCommand("bash", "-c", "sudo cp -R "+staticDirFrom+" "+staticDirTo)
	execCommand("bash", "-c", "sudo cp -R "+notebookDirFrom+" "+notebookDirTo)

	// fmt.Println("Unpackaging files")

	fmt.Println()
	return nil
}
