package main

import (
	"encoding/json"
	"errors"
	"net/http"
	proto "thingz-server/api/proto"
	"thingz-server/lib"
	thingP "thingz-server/thing/proto"
	thingT "thingz-server/thing/topics"

	log "github.com/sirupsen/logrus"

	"github.com/alexandrevicenzi/go-sse"
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

func (a *app) getThing(c echo.Context) error {
	thing := c.Param("id")
	req := &thingP.GetThingRequest{
		Thing: thing,
	}
	res := &thingP.GetThingResponse{}

	err := a.eb.RequestMessage(thingT.GetThing, req, res, lib.DefaultTimeout)
	if err != nil {
		log.Printf("error: %+v", err)
		return a.makeError(c, http.StatusInternalServerError, err)
	}

	if !res.GetSuccess() {
		return a.makeError(c, http.StatusUnauthorized, errors.New(res.GetError()))
	}

	return a.sendSucess(c, echo.Map{
		"msg":   "ok",
		"thing": res.Thing,
	})
}

func (a *app) updateThing(c echo.Context) error {
	thing := c.Param("id")
	req := &thingP.UpdateThingRequest{
		Thing: thing,
		Name:  c.QueryParam("name"),
		Group: c.QueryParam("group"),
	}
	res := &thingP.UpdateThingResponse{}

	err := a.eb.RequestMessage(thingT.UpdateThing, req, res, lib.DefaultTimeout)
	if err != nil {
		log.Printf("error: %+v", err)
		return a.makeError(c, http.StatusInternalServerError, err)
	}

	if !res.GetSuccess() {
		return a.makeError(c, http.StatusUnauthorized, errors.New(res.GetError()))
	}

	return a.sendSucess(c, echo.Map{
		"msg": "ok",
	})
}

func (a *app) updateThingChannelName(c echo.Context) error {
	thing := c.Param("id")
	req := &thingP.UpdateChannelNameRequest{
		Thing:   thing,
		Channel: c.QueryParam("channel"),
		Name:    c.QueryParam("name"),
	}
	res := &thingP.UpdateChannelNameResponse{}

	err := a.eb.RequestMessage(thingT.UpdateChannelName, req, res, lib.DefaultTimeout)
	if err != nil {
		log.Printf("error: %+v", err)
		return a.makeError(c, http.StatusInternalServerError, err)
	}

	if !res.GetSuccess() {
		return a.makeError(c, http.StatusUnauthorized, errors.New(res.GetError()))
	}

	return a.sendSucess(c, echo.Map{
		"msg": "ok",
	})
}

func (a *app) getThingSeries(c echo.Context) error {
	thing := c.Param("id")
	req := &thingP.ThingSeriesRequest{
		Id:      thing,
		Channel: c.QueryParam("channel"),
	}
	res := &thingP.ThingSeriesResponse{}

	err := a.eb.RequestMessage(thingT.ThingSeries, req, res, lib.DefaultTimeout)
	if err != nil {
		log.Printf("error: %+v", err)
		return a.makeError(c, http.StatusInternalServerError, err)
	}

	if !res.GetSuccess() {
		return a.makeError(c, http.StatusUnauthorized, errors.New(res.GetError()))
	}

	return a.sendSucess(c, echo.Map{
		"msg":    "ok",
		"values": res.Values,
	})
}

func (a *app) projectThings(c echo.Context) error {
	project := c.Param("id")
	group := c.QueryParam("group")
	things := []*thingP.Thing{}

	if group == "" {
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

		if res.Things == nil {
			res.Things = []*thingP.Thing{}
		}

		things = res.Things
	} else {
		req := &thingP.ProjectGroupThingsRequest{
			Project: project,
		}
		res := &thingP.ProjectGroupThingsResponse{}

		err := a.eb.RequestMessage(thingT.ProjectGroupThings, req, res, lib.DefaultTimeout)
		if err != nil {
			log.Printf("error: %+v", err)
			return a.makeError(c, http.StatusInternalServerError, err)
		}

		if !res.GetSuccess() {
			return a.makeError(c, http.StatusUnauthorized, errors.New(res.GetError()))
		}

		if res.Things == nil {
			res.Things = []*thingP.Thing{}
		}

		things = res.Things
	}

	log.Println("got things", things)

	return a.sendSucess(c, echo.Map{
		"msg":    "ok",
		"things": things,
	})
}

func (a *app) deleteProjectThings(c echo.Context) error {
	project := c.Param("id")
	reqT := &thingP.ProjectDeleteRequest{
		Project: project,
	}
	resT := &thingP.ProjectDeleteResponse{}

	err := a.eb.RequestMessage(thingT.ProjectDelete, reqT, resT, lib.DefaultTimeout)
	if err != nil {
		log.Printf("error: %+v", err)
		return a.makeError(c, http.StatusInternalServerError, err)
	}

	if !resT.GetSuccess() {
		return a.makeError(c, http.StatusUnauthorized, errors.New(resT.GetError()))
	}

	return a.sendSucess(c, echo.Map{
		"msg": "ok",
	})
}

func (a *app) deleteThing(c echo.Context) error {
	thing := c.Param("id")
	req := &thingP.DeleteThingRequest{
		Thing: thing,
	}
	res := &thingP.DeleteThingResponse{}

	err := a.eb.RequestMessage(thingT.DeleteThing, req, res, lib.DefaultTimeout)
	if err != nil {
		log.Printf("error: %+v", err)
		return a.makeError(c, http.StatusInternalServerError, err)
	}

	if !res.GetSuccess() {
		return a.makeError(c, http.StatusUnauthorized, errors.New(res.GetError()))
	}

	return a.sendSucess(c, echo.Map{
		"msg": "ok",
	})
}

func (a *app) generateAssignThing(c echo.Context) error {
	thing := c.Param("id")
	project := c.QueryParam("project")
	req := &thingP.GenerateAssignThingRequest{
		Id:      thing,
		Project: project,
	}
	res := &thingP.GenerateAssignThingResponse{}

	err := a.eb.RequestMessage(thingT.GenerateAssignThing, req, res, lib.DefaultTimeout)
	if err != nil {
		log.Printf("error: %+v", err)
		return a.makeError(c, http.StatusInternalServerError, err)
	}

	if !res.GetSuccess() {
		return a.makeError(c, http.StatusUnauthorized, errors.New(res.GetError()))
	}

	return a.sendSucess(c, res.Token)
}

func (a *app) assignThing(c echo.Context) error {
	token := c.QueryParam("token")
	key := c.QueryParam("key")
	req := &thingP.AssignThingRequest{
		Token: token,
		Key:   key,
	}
	res := &thingP.AssignThingResponse{}

	err := a.eb.RequestMessage(thingT.AssignThing, req, res, lib.DefaultTimeout)
	if err != nil {
		log.Printf("error: %+v", err)
		return a.makeError(c, http.StatusInternalServerError, err)
	}

	if !res.GetSuccess() {
		return a.makeError(c, http.StatusUnauthorized, errors.New(res.GetError()))
	}

	return a.sendSucess(c, echo.Map{
		"msg": "ok",
	})
}

func (a *app) deassignThing(c echo.Context) error {
	thing := c.Param("id")
	project := c.QueryParam("project")
	req := &thingP.DeassignThingRequest{
		Id:      thing,
		Project: project,
	}
	res := &thingP.DeassignThingResponse{}

	err := a.eb.RequestMessage(thingT.DeassignThing, req, res, lib.DefaultTimeout)
	if err != nil {
		log.Printf("error: %+v", err)
		return a.makeError(c, http.StatusInternalServerError, err)
	}

	if !res.GetSuccess() {
		return a.makeError(c, http.StatusUnauthorized, errors.New(res.GetError()))
	}

	return a.sendSucess(c, echo.Map{
		"msg": "ok",
	})
}

func (a *app) updateChannel(c echo.Context) error {
	thing := c.Param("id")
	physicalID := c.QueryParam("physicalId")
	updateReq := &thingP.UpdateThingsChannelsRequest{
		Things: []*thingP.ThingChannels{},
	}

	thingChannel := &thingP.ThingChannels{
		Id:         thing,
		PhysicalId: physicalID,
		Channels:   []*thingP.Channel{},
	}

	channel := &thingP.Channel{}
	err := c.Bind(channel)
	if err != nil {
		log.Println(err)
		return a.makeError(c, http.StatusBadRequest, err)
	}

	thingChannel.Channels = append(thingChannel.Channels, channel)

	updateReq.Things = append(updateReq.Things, thingChannel)

	updateRes := &thingP.UpdateThingsChannelsResponse{}

	log.Printf("updateCHanel phy: %v || %+v", physicalID, updateReq)

	err = a.eb.RequestMessage(thingT.UpdateThingsChannels, updateReq, updateRes, lib.DefaultTimeout)
	if err != nil {
		log.Println(err)
		return a.makeError(c, http.StatusBadRequest, err)
	}

	return a.sendSucess(c, echo.Map{
		"msg": "ok",
	})
}

func (a *app) handleUpdateThingEvent(c echo.Context) error {
	a.eventServer.ServeHTTP(c.Response(), c.Request())
	return nil
}

func (a *app) sendUpdateThing(_, reply string, req *proto.SendThingUpdateRequest) {
	for _, ch := range req.Update.Channels {
		ch.Thing = req.Update.Thing
	}

	b, err := json.Marshal(req.Update.Channels)
	if err != nil {
		log.Println(err)
		return
	}

	a.eventServer.SendMessage("/api/thing/events", sse.SimpleMessage(string(b)))

	log.Printf("event sent to %v, %+v", req.Update.Thing, req.Update.Channels)
}

func (a *app) updateThingConfig(c echo.Context) error {
	req := &thingP.UpdateThingConfigRequest{
		Thing:  c.Param("id"),
		Config: c.QueryParam("config"),
	}
	resp := &thingP.UpdateThingConfigResponse{}

	err := a.eb.RequestMessage(thingT.UpdateThingConfig, req, resp, lib.DefaultTimeout)
	if err != nil {
		log.Println(err)
		return a.makeError(c, http.StatusBadRequest, err)
	}
	return a.sendSucess(c, echo.Map{
		"msg": "ok",
	})
}
