package main

import (
	"errors"
	"log"
	"net/http"
	"thingz-server/lib"
	thingP "thingz-server/thing/proto"
	thingT "thingz-server/thing/topics"

	"github.com/labstack/echo"
)

func (a *app) createThing(c echo.Context) error {
	thing := &thingP.Thing{}

	err := c.Bind(thing)
	if err != nil {
		log.Println("bind error")
		return a.makeError(c, http.StatusBadRequest, err)
	}

	/* for _, ch := range thing.Channels {
		ch.Id = lib.NewObjectID()
	} */

	req := &thingP.CreateThingRequest{
		Thing: thing,
	}

	res := &thingP.CreateThingResponse{}

	err = a.eb.RequestMessage(thingT.CreateThing, req, res, lib.DefaultTimeout)
	if err != nil {
		log.Printf("error: %+v", err)
		return a.makeError(c, http.StatusInternalServerError, err)
	}

	if !res.GetSuccess() {
		return a.makeError(c, http.StatusUnauthorized, errors.New(res.GetError()))
	}
	return a.sendSucess(c, echo.Map{
		"msg": "ok",
		"id":  res.GetId(),
	})
}

func (a *app) projectThings(c echo.Context) error {
	project := c.Param("id")
	req := &thingP.ProjectThingsRequest{
		Project: project,
	}
	res := &thingP.ProjectThingsResponse{}

	err := a.eb.RequestMessage(thingT.ProjectThings, req, res, lib.DefaultTimeout)
	if err != nil {
		log.Printf("error: %+v", err)
		return a.makeError(c, http.StatusInternalServerError, err)
	}

	if !res.GetSuccess() {
		return a.makeError(c, http.StatusUnauthorized, errors.New(res.GetError()))
	}
	return a.sendSucess(c, echo.Map{
		"msg":    "ok",
		"things": res.GetThings(),
	})
}
