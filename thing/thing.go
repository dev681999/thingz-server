package main

import (
	"log"
	"thingz-server/lib"
	proto "thingz-server/thing/proto"
)

func (a *app) createThing(_, reply string, req *proto.CreateThingRequest) {
	log.Printf("create req: %+v", req)
	resp := &proto.CreateThingResponse{}
	db, err := a.db.GetMongoSession()
	if err == nil {
		defer db.Close()
		req.Thing.Id = lib.NewObjectID()
		err = db.DB("").C(collectionName).Insert(req.Thing)
	}
	if err != nil {
		resp.Success = false
		resp.Error = err.Error()
	} else {
		resp.Success = true
		resp.Id = req.Thing.Id
	}

	log.Printf("Request: %+v, Resposne: %+v", req, resp)
	if reply != "" {
		a.eb.SendMessage(reply, resp)
	}
}

func (a *app) projectThings(_, reply string, req *proto.ProjectThingsRequest) {
	log.Printf("create req: %+v", req)
	resp := &proto.ProjectThingsResponse{}
	things := []*proto.Thing{}
	db, err := a.db.GetMongoSession()
	if err == nil {
		defer db.Close()
		db.DB("").C(collectionName).Find(lib.M{"project": req.Project}).All(&things)
	}
	if err != nil {
		resp.Success = false
		resp.Error = err.Error()
	} else {
		resp.Success = true
		resp.Things = things
	}

	log.Printf("Request: %+v, Resposne: %+v", req, resp)
	if reply != "" {
		a.eb.SendMessage(reply, resp)
	}
}
