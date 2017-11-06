package pathmonitor_test

import (
	"fmt"
	"github.com/dargad/pathmonitor/internal/app/pathmonitor"
	"io/ioutil"
	"os"
	"testing"
	"time"
)

var (
	monitor       *pathmonitor.Monitor
	testDir       string
	config        pathmonitor.Config
	configContent string = `
log:
    level: debug
    output: file:///tmp/pathmonitor.log
paths:
    - path: %s
      filter: .*
      command: rm {}
`
)

func AssertTrue(t *testing.T, cond bool, failMsg string) {
	if !cond {
		t.Error("assertion failed", failMsg)
	}
}

func Init() {
	if monitor == nil {
		var err error
		pathmonitor.LogInit()
		testDir, err = ioutil.TempDir("", "")
		if err != nil {
			pathmonitor.Error.Println("Can't create test dir: ", err)
		}
		pathmonitor.Trace.Println("Created temp dir:", testDir)
		configContent = fmt.Sprintf(configContent, testDir)

		configFile, err := ioutil.TempFile(testDir, "config.yaml")
		configFile.Write([]byte(configContent))
		defer func() { configFile.Close(); os.Remove(configFile.Name()) }()

		config, err = pathmonitor.ReadConfig(configFile.Name())
	}
}

func TestFileAdded(t *testing.T) {
	Init()
	monitor = pathmonitor.NewMonitor(config)
	go monitor.Run()
	f, err := ioutil.TempFile(testDir, "")
	if err != nil {
		pathmonitor.Error.Println("Failed creating test file.")
	}
	f.Close()

	time.Sleep(5 * time.Millisecond)

	_, err = os.Stat(f.Name())
	AssertTrue(t, os.IsNotExist(err), "failed to remove file")
}

func TestFileDetection(t *testing.T) {
	//t.Error("not implemented")
}
