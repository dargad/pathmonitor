package main

import (
	"flag"
	"github.com/dargad/pathmonitor/internal/app/pathmonitor"
)

var (
	configPath string
)

func initFlags() {
	const usageConfig = "path of the config file"
	flag.StringVar(&configPath, "c", "/etc/pathmonitor.conf", usageConfig)
}

func main() {
	initFlags()
	flag.Parse()
	pathmonitor.LogInit(pathmonitor.Log{Level: "error", Output: "stdout"})
	pathmonitor.Trace.Println("Reading config.")
	c, err := pathmonitor.ReadConfig(configPath)
	if err != nil {
		pathmonitor.Error.Println("Failed loading config (", configPath, ")",
			err)
		return
	}
	pathmonitor.LogInit(c.Log)
	pathmonitor.Trace.Println("Creating monitor.")
	m := pathmonitor.NewMonitor(c)
	pathmonitor.Trace.Println("Running monitor...")
	m.Run()
	pathmonitor.Trace.Println("Monitor exited.")
}
