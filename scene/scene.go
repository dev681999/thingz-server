package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"thingz-server/lib"
	proto "thingz-server/scene/proto"
	thingP "thingz-server/thing/proto"

	thingT "thingz-server/thing/topics"

	log "github.com/sirupsen/logrus"

	"github.com/globalsign/mgo/bson"
)

func init() {
	b, _ := json.MarshalIndent(proto.Scene{
		Id:    "fghjk",
		Name:  "fghjkl",
		Owner: "fghjkl",
		Things: []*proto.Thing{{
			Id: "fgbnm",
			Channels: []*proto.Channel{{
				Type:        1,
				FloatValue:  7,
				BoolValue:   true,
				DataValue:   "sad",
				StringValue: "fgh",
			}},
		}},
	}, "", "	")
	fmt.Println(string(b))
}

func (a *app) createScene(_, reply string, req *proto.CreateSceneRequest) {
	log.Printf("create req: %+v", req)
	resp := &proto.CreateSceneResponse{}
	db, err := a.db.GetMongoSession()
	if err == nil {
		defer db.Close()
		req.Scene.Id = lib.NewObjectID()
		err = db.DB("").C(collectionName).Insert(req.Scene)
	}
	if err != nil {
		resp.Success = false
		resp.Error = err.Error()
	} else {
		resp.Success = true
		resp.Id = req.Scene.Id
	}

	log.Printf("Request: %+v, Resposne: %+v", req, resp)
	if reply != "" {
		a.eb.SendMessage(reply, resp)
	}
}

func (a *app) userScenes(_, reply string, req *proto.UserScenesRequest) {
	log.Printf("userScenes req: %+v", req)
	resp := &proto.UserScenesResponse{}
	scenes := []*proto.Scene{}
	db, err := a.db.GetMongoSession()
	if err == nil {
		defer db.Close()
		db.DB("").C(collectionName).Find(lib.M{"owner": req.GetOwner(), "project": req.GetProject()}).All(&scenes)
	}
	if err != nil {
		resp.Success = false
		resp.Error = err.Error()
	} else {
		resp.Success = true
		resp.Scenes = scenes
	}

	log.Printf("Request: %+v, Resposne: %+v", req, resp)
	if reply != "" {
		err = a.eb.SendMessage(reply, resp)
		if err != nil {
			log.Println("reply err", err)
		}
	}
}

func (a *app) deleteScene(_, reply string, req *proto.DeleteSceneRequest) {
	log.Printf("create req: %+v", req)
	resp := &proto.DeleteSceneResponse{}
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

func (a *app) updateScene(_, reply string, req *proto.UpdateSceneRequest) {
	log.Printf("update req: %+v", req)
	resp := &proto.UpdateSceneResponse{}
	db, err := a.db.GetMongoSession()
	if err == nil {
		defer db.Close()
		err = db.DB("").C(collectionName).Update(bson.M{"_id": req.Scene.Id}, req.Scene)
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

func (a *app) activateScene(_, reply string, req *proto.ActivateSceneRequest) {
	log.Printf("userScenes req: %+v", req)
	resp := &proto.ActivateSceneResponse{}
	scene := &proto.Scene{}
	db, err := a.db.GetMongoSession()
	if err == nil {
		defer db.Close()
		db.DB("").C(collectionName).Find(lib.M{"_id": req.GetId(), "owner": req.GetOwner()}).One(scene)

		thingReq := thingP.UpdateThingsChannelsRequest{
			Things: []*thingP.ThingChannels{},
		}

		for _, t := range scene.Things {
			channels := []*thingP.Channel{}
			for _, ch := range t.Channels {
				channels = append(channels, &thingP.Channel{
					DataValue:   ch.DataValue,
					BoolValue:   ch.BoolValue,
					FloatValue:  ch.FloatValue,
					StringValue: ch.StringValue,
					Type:        ch.Type,
					Unit:        thingP.Unit(ch.Unit),
				})
			}

			thingReq.Things = append(thingReq.Things, &thingP.ThingChannels{
				Id:         t.Id,
				Channels:   channels,
				PhysicalId: t.PhysicalId,
			})
		}

		if len(thingReq.Things) > 0 {
			log.Printf("scene handle %+v", thingReq)
			updateRes := &thingP.UpdateThingsChannelsResponse{}

			err = a.eb.RequestMessage(thingT.UpdateThingsChannels, &thingReq, updateRes, lib.DefaultTimeout)
			if err == nil {
				log.Printf("exec scene %+v, %+v", scene, err)

				if updateRes.GetError() != "" {
					err = errors.New(updateRes.GetError())
				}
			}
		}

		resp.Success = true
	}
	if err != nil {
		resp.Success = false
		resp.Error = err.Error()
	} else {
		resp.Success = true
		// resp.Scenes = scene
	}

	log.Printf("Request: %+v, Resposne: %+v", req, resp)
	if reply != "" {
		err = a.eb.SendMessage(reply, resp)
		if err != nil {
			log.Println("reply err", err)
		}
	}
}

func (a *app) activateScenes(_, reply string, req *proto.ActivateScenesRequest) {
	log.Printf("actviateScenes req: %+v", req)
	resp := &proto.ActivateScenesResponse{}
	scenes := []*proto.Scene{}
	db, err := a.db.GetMongoSession()
	if err == nil {
		defer db.Close()
		db.DB("").C(collectionName).Find(lib.M{"_id": bson.M{"$in": req.Ids}, "owner": req.GetOwner()}).All(&scenes)

		thingReq := thingP.UpdateThingsChannelsRequest{
			Things: []*thingP.ThingChannels{},
		}

		for _, scene := range scenes {

			for _, t := range scene.Things {
				channels := []*thingP.Channel{}
				for _, ch := range t.Channels {
					channels = append(channels, &thingP.Channel{
						DataValue:   ch.DataValue,
						BoolValue:   ch.BoolValue,
						FloatValue:  ch.FloatValue,
						StringValue: ch.StringValue,
						Type:        ch.Type,
						Unit:        thingP.Unit(ch.Unit),
					})
				}

				thingReq.Things = append(thingReq.Things, &thingP.ThingChannels{
					Id:         t.Id,
					Channels:   channels,
					PhysicalId: t.PhysicalId,
				})
			}
		}

		if len(thingReq.Things) > 0 {
			log.Printf("scene handle %+v", thingReq)
			updateRes := &thingP.UpdateThingsChannelsResponse{}

			err = a.eb.RequestMessage(thingT.UpdateThingsChannels, &thingReq, updateRes, lib.DefaultTimeout)
			if err == nil {
				// log.Printf("exec scene %+v, %+v", scene, err)

				if updateRes.GetError() != "" {
					err = errors.New(updateRes.GetError())
				}
			}
		}

		resp.Success = true
	}
	if err != nil {
		resp.Success = false
		resp.Error = err.Error()
	} else {
		resp.Success = true
		// resp.Scenes = scene
	}

	log.Printf("Request: %+v, Resposne: %+v", req, resp)
	if reply != "" {
		err = a.eb.SendMessage(reply, resp)
		if err != nil {
			log.Println("reply err", err)
		}
	}
}
