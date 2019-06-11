package main

import (
	"log"
	"thingz-server/lib"
)

func main() {
	c := &appConfig{}
	err := lib.ConfigFromFile(c)
	if err != nil {
		panic(err)
	}

	log.Printf("config: %+v", c)

	a := newApp(c)

	err = lib.RunApp(a)
	if err != nil {
		panic(err)
	}
}
