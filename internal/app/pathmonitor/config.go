package pathmonitor

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Log struct {
	Level  string
	Output string
}

type Config struct {
	Log
	Paths []struct {
		Path    string
		Filter  string
		Command string
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadConfig(filename string) (Config, error) {
	c := Config{}
	data, err := ioutil.ReadFile(filename)
	check(err)
	err = yaml.Unmarshal([]byte(data), &c)
	if err != nil {
		Error.Println("Error while reading config:", err)
	}
	return c, err
}
