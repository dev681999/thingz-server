package main

import (
	"thingz-server/lib"

	log "github.com/sirupsen/logrus"

	"github.com/google/uuid"
)

func main() {

	c := &appConfig{}
	log.SetReportCaller(true)
	err := lib.ConfigFromFile(c)
	if err != nil {
		panic(err)
	}

	if c.JwtSecret == "" {
		t, _ := uuid.NewRandom()
		jwt := t.String()

		c.JwtSecret = jwt

		err = lib.SaveConfigToFile(c)
		if err != nil {
			panic(err)
		}
	}

	log.Printf("config: %+v", c)

	lib.WaitFor([]string{c.NATSUrl + ":4222"})

	a := newApp(c)

	err = lib.RunApp(a)
	if err != nil {
		panic(err)
	}
}
