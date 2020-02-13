package main

import (
	"thingz-server/lib"
	proto "thingz-server/project/proto"

	log "github.com/sirupsen/logrus"

	"github.com/globalsign/mgo/bson"
)

func (a *app) createProject(_, reply string, req *proto.CreateProjectRequest) {
	log.Printf("create req: %+v", req)
	resp := &proto.CreateProjectResponse{}
	db, err := a.db.GetMongoSession()
	if err == nil {
		defer db.Close()
		req.Project.Id = lib.NewObjectID()
		err = db.DB("").C(collectionName).Insert(req.Project)
	}
	if err != nil {
		resp.Success = false
		resp.Error = err.Error()
	} else {
		resp.Success = true
		resp.Id = req.Project.Id
	}

	log.Printf("Request: %+v, Resposne: %+v", req, resp)
	if reply != "" {
		a.eb.SendMessage(reply, resp)
	}
}

func (a *app) userProjects(_, reply string, req *proto.UserProjectsRequest) {
	log.Printf("create req: %+v", req)
	resp := &proto.UserProjectsResponse{}
	projects := []*proto.Project{}
	db, err := a.db.GetMongoSession()
	if err == nil {
		defer db.Close()
		db.DB("").C(collectionName).Find(lib.M{"owner": req.GetOwner()}).All(&projects)
	}
	if err != nil {
		resp.Success = false
		resp.Error = err.Error()
	} else {
		resp.Success = true
		resp.Projects = projects
	}

	log.Printf("Request: %+v, Resposne: %+v", req, resp)
	if reply != "" {
		a.eb.SendMessage(reply, resp)
	}
}

func (a *app) deleteProject(_, reply string, req *proto.DeleteProjectRequest) {
	log.Printf("create req: %+v", req)
	resp := &proto.DeleteProjectResponse{}
	db, err := a.db.GetMongoSession()
	if err == nil {
		defer db.Close()
		_, err = db.DB("").C(collectionName).RemoveAll(bson.M{
			"_id": req.Id,
		})
	}
	if err != nil {
		resp.Success = false
		resp.Error = err.Error()
	} else {
		resp.Success = true
	}

	log.Printf("Request: %+v, Resposne: %+v", req, resp)
	if reply != "" {
		a.eb.SendMessage(reply, resp)
	}
}
