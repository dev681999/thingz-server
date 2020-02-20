package main

import (
	"errors"
	"net/http"
	"thingz-server/lib"
	sceneP "thingz-server/scene/proto"
	sceneT "thingz-server/scene/topics"

	log "github.com/sirupsen/logrus"

	"github.com/labstack/echo"
)

func (a *app) createScene(c echo.Context) error {
	scene := &sceneP.Scene{}

	err := c.Bind(scene)
	if err != nil {
		log.Println("bind error")
		return a.makeError(c, http.StatusBadRequest, err)
	}

	owner := getUserFromContext(c)["id"].(string)
	scene.Owner = owner

	req := &sceneP.CreateSceneRequest{
		Scene: scene,
	}

	res := &sceneP.CreateSceneResponse{}

	err = a.eb.RequestMessage(sceneT.CreateScene, req, res, lib.DefaultTimeout)
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

func (a *app) userScenes(c echo.Context) error {
	project := c.Param("id")
	owner := getUserFromContext(c)["id"].(string)
	req := &sceneP.UserScenesRequest{
		Owner:   owner,
		Project: project,
	}
	res := &sceneP.UserScenesResponse{}

	err := a.eb.RequestMessage(sceneT.UserScenes, req, res, lib.DefaultTimeout)
	if err != nil {
		log.Printf("error: %+v", err)
		return a.makeError(c, http.StatusInternalServerError, err)
	}

	if !res.GetSuccess() {
		return a.makeError(c, http.StatusUnauthorized, errors.New(res.GetError()))
	}

	if res.Scenes == nil {
		res.Scenes = []*sceneP.Scene{}
	}

	log.Printf("scenes: %+v", res.Scenes)

	return a.sendSucess(c, echo.Map{
		"msg":    "ok",
		"scenes": res.GetScenes(),
	})
}

func (a *app) deleteScene(c echo.Context) error {
	scene := c.Param("id")
	req := &sceneP.DeleteSceneRequest{
		Id:    scene,
		Owner: getUserFromContext(c)["id"].(string),
	}
	res := &sceneP.DeleteSceneResponse{}

	err := a.eb.RequestMessage(sceneT.DeleteScene, req, res, lib.DefaultTimeout)
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

func (a *app) updateScene(c echo.Context) error {
	scene := &sceneP.Scene{}

	err := c.Bind(scene)
	if err != nil {
		log.Println("bind error")
		return a.makeError(c, http.StatusBadRequest, err)
	}

	owner := getUserFromContext(c)["id"].(string)
	scene.Owner = owner
	scene.Id = c.Param("id")

	req := &sceneP.UpdateSceneRequest{
		Scene: scene,
	}

	res := &sceneP.UpdateSceneResponse{}

	err = a.eb.RequestMessage(sceneT.UpdateScene, req, res, lib.DefaultTimeout)
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

func (a *app) activateScene(c echo.Context) error {
	scene := c.Param("id")
	req := &sceneP.ActivateSceneRequest{
		Id:    scene,
		Owner: getUserFromContext(c)["id"].(string),
	}
	res := &sceneP.ActivateSceneResponse{}

	log.Printf("sending request %+v", *req)

	err := a.eb.RequestMessage(sceneT.ActivateScene, req, res, lib.DefaultTimeout)
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
