package lib

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
)

const (
	defaultConfigFile = "./config.json"
)

type app interface {
	Init() error
	Close()
	Test() string
}

// ConfigFromFile return a config from a JSON file
func ConfigFromFile(config interface{}) error {
	file := defaultConfigFile
	if len(os.Args) > 1 && os.Args[1] != "" {
		file = os.Args[1]
	}

	data, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, config)
}

// RunApp return a config from a JSON file
func RunApp(a app) error {
	err := a.Init()
	if err != nil {
		return err
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch

	log.Println("interrupt occured")

	a.Close()

	return nil
}
