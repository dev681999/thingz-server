package main

import (
	"errors"
	"net/http"
	"thingz-server/lib"
	projectP "thingz-server/project/proto"
	projectT "thingz-server/project/topics"

	log "github.com/sirupsen/logrus"

	"github.com/labstack/echo"
)

func (a *app) createProject(c echo.Context) error {
	project := &projectP.Project{}

	err := c.Bind(project)
	if err != nil {
		log.Println("bind error")
		return a.makeError(c, http.StatusBadRequest, err)
	}

	owner := getUserFromContext(c)["id"].(string)
	project.Owner = owner

	req := &projectP.CreateProjectRequest{
		Project: project,
	}

	res := &projectP.CreateProjectResponse{}

	err = a.eb.RequestMessage(projectT.CreateProject, req, res, lib.DefaultTimeout)
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

func (a *app) userProjects(c echo.Context) error {
	owner := getUserFromContext(c)["id"].(string)
	req := &projectP.UserProjectsRequest{
		Owner: owner,
	}
	res := &projectP.UserProjectsResponse{}

	err := a.eb.RequestMessage(projectT.UserProjects, req, res, lib.DefaultTimeout)
	if err != nil {
		log.Printf("error: %+v", err)
		return a.makeError(c, http.StatusInternalServerError, err)
	}

	if !res.GetSuccess() {
		return a.makeError(c, http.StatusUnauthorized, errors.New(res.GetError()))
	}

	if res.Projects == nil {
		res.Projects = []*projectP.Project{}
	}

	return a.sendSucess(c, echo.Map{
		"msg":      "ok",
		"projects": res.GetProjects(),
	})
}

func (a *app) deleteProject(c echo.Context) error {
	project := c.Param("id")
	req := &projectP.DeleteProjectRequest{
		Id:    project,
		Owner: getUserFromContext(c)["id"].(string),
	}
	res := &projectP.DeleteProjectResponse{}

	err := a.eb.RequestMessage(projectT.DeleteProject, req, res, lib.DefaultTimeout)
	if err != nil {
		log.Printf("error: %+v", err)
		return a.makeError(c, http.StatusInternalServerError, err)
	}

	if !res.GetSuccess() {
		return a.makeError(c, http.StatusUnauthorized, errors.New(res.GetError()))
	}

	return a.deleteProjectThings(c)
}

func (a *app) addProjectGroup(c echo.Context) error {
	project := c.Param("id")
	req := &projectP.AddProjectGroupRequest{
		Id:    project,
		Owner: getUserFromContext(c)["id"].(string),
		Name:  c.QueryParam("name"),
	}
	res := &projectP.AddProjectGroupResponse{}

	log.Printf("sending request %+v", *req)

	err := a.eb.RequestMessage(projectT.AddProjectGroup, req, res, lib.DefaultTimeout)
	if err != nil {
		log.Printf("error: %+v", err)
		return a.makeError(c, http.StatusInternalServerError, err)
	}

	if !res.GetSuccess() {
		return a.makeError(c, http.StatusUnauthorized, errors.New(res.GetError()))
	}

	return a.deleteProjectThings(c)
}

func (a *app) deleteProjectGroup(c echo.Context) error {
	project := c.Param("id")
	req := &projectP.DeleteProjectGroupRequest{
		Id:    project,
		Owner: getUserFromContext(c)["id"].(string),
		Name:  c.QueryParam("name"),
	}
	res := &projectP.DeleteProjectGroupResponse{}

	err := a.eb.RequestMessage(projectT.DeleteProjectGroup, req, res, lib.DefaultTimeout)
	if err != nil {
		log.Printf("error: %+v", err)
		return a.makeError(c, http.StatusInternalServerError, err)
	}

	if !res.GetSuccess() {
		return a.makeError(c, http.StatusUnauthorized, errors.New(res.GetError()))
	}

	return a.deleteProjectThings(c)
}
