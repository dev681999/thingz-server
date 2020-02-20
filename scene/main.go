package main

import (
	"thingz-server/lib"

	log "github.com/sirupsen/logrus"
)

func main() {
	c := &appConfig{}
	log.SetReportCaller(true)
	err := lib.ConfigFromFile(c)
	if err != nil {
		panic(err)
	}

	log.Printf("config: %+v", c)

	lib.WaitFor([]string{c.NATSUrl + ":4222"})

	a := newApp(c)

	err = lib.RunApp(a)
	if err != nil {
		panic(err)
	}
}
