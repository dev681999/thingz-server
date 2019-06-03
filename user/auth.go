package main

import (
	"log"
	"thingz-server/lib"
	proto "thingz-server/user/proto"

	"github.com/globalsign/mgo/bson"
)

func (a *app) createUser(_, reply string, req *proto.CreateUserRequest) {
	resp := &proto.CreateUserResponse{}
	db, err := a.db.GetMongoSession()
	if err == nil {
		defer db.Close()
		hash, err := a.h.Generate(req.User.Password)
		if err == nil {
			req.User.Password = hash
			req.User.Id = lib.NewObjectID()
			err = db.DB("").C(collectionName).Insert(req.User)
		}
	}
	if err != nil {
		resp.Success = false
		resp.Error = err.Error()
	}

	if reply != "" {
		a.eb.SendMessage(reply, resp)
	} else {
		log.Printf("Request: %+v, Resposne: %+v", req, resp)
	}
}

func (a *app) verifyUser(_, reply string, req *proto.VerifyUserRequest) {
	resp := &proto.VerifyUserResponse{}
	user := &proto.User{}
	db, err := a.db.GetMongoSession()
	if err == nil {
		defer db.Close()
		err = db.DB("").C(collectionName).Find(bson.M{"email": req.Email}).One(user)
		if err == nil {
			err = a.h.Compare(user.Password, req.Password)
		}
	}
	if err != nil {
		resp.Success = false
		resp.Error = err.Error()
	} else {
		resp.User = user
	}

	if reply != "" {
		a.eb.SendMessage(reply, resp)
	} else {
		log.Printf("Request: %+v, Resposne: %+v", req, resp)
	}
}
