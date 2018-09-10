package pathmonitor

import (
	"errors"
	"github.com/fsnotify/fsnotify"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

const (
	UPDIR = ".."
)

type Monitor struct {
	watcher *fsnotify.Watcher
	config  Config
}

func NewMonitor(c Config) *Monitor {
	m := new(Monitor)
	m.config = c
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		panic(err)
	}
	m.watcher = watcher

	for _, path := range m.config.Paths {
		m.addRecursive(path.Path)
	}

	return m
}

func (m *Monitor) addRecursive(p string) {
	filepath.Walk(p, func(path string, info os.FileInfo, err error) error {
		Info.Println("Adding path:", path)
		err = m.watcher.Add(path)
		if err != nil {
			Warning.Println("Skipping path", p, ":", err)
		}
		return nil
	})
}

func isDirectory(path string) bool {
	fi, err := os.Stat(path)
	if err != nil {
		Info.Println("Testing", path, ":", err)
		return false
	}
	return fi.IsDir()
}

func (m *Monitor) scanDirectory(dir string) {
	Info.Println("Scanning new directory:", dir)
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if !isDirectory(path) {
			m.executeIfFileMatches(path)
		}
		return nil
	})
}

func (m *Monitor) Run() {
	for {
		select {
		case event := <-m.watcher.Events:
			if event.Op&(fsnotify.Create|fsnotify.Rename) > 0 {
				if isDirectory(event.Name) {
					m.scanDirectory(event.Name)
					m.addRecursive(event.Name)
				}
				Info.Println("File added:", event.Name)
				m.executeIfFileMatches(event.Name)
			}
		case err := <-m.watcher.Errors:
			log.Println("error:", err)
		}
	}
}

func (m *Monitor) executeIfFileMatches(filename string) bool {
	pattern, command, err := m.findPatternCommandForPath(filename)
	Trace.Printf("Checking pattern '%s' for file '%s'", pattern, filename)

	if err == nil {
		if match, err := regexp.MatchString(pattern, strings.ToLower(filename)); err != nil {
			Error.Printf("Can't match pattern '%s' to file '%s': %s",
				pattern, filename, err)
		} else if match {
			Trace.Println("Executing command:", command)
			res, err := ExecuteCommand(ReplacePlaceholders(command, filename))
			if err != nil {
				Error.Printf("Error executing '%s': %s", command, err)
				return false
			}
			Trace.Println("Command executed:", res)
			return res
		}
	}

	return false
}

func (m *Monitor) findPatternCommandForPath(filename string) (string, string, error) {
	Trace.Println("Looking for dir for file:", filename)
	for _, path := range m.config.Paths {
		Trace.Println("Checking", path.Path)
		rel, err := filepath.Rel(path.Path, filename)
		if err == nil && !strings.HasPrefix(rel, UPDIR) {
			return path.Filter, path.Command, nil
		}
	}
	return "", "", errors.New("Path not found")
}
