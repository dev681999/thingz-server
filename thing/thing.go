package main

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"thingz-server/lib"
	mqttProto "thingz-server/mqtt/proto"
	mqttTopics "thingz-server/mqtt/topics"
	ruleProto "thingz-server/rule/proto"
	ruleTopics "thingz-server/rule/topics"
	proto "thingz-server/thing/proto"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/google/uuid"
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

func (a *app) getThing(_, reply string, req *proto.GetThingRequest) {
	log.Printf("create req: %+v", req)
	resp := &proto.GetThingResponse{}
	db, err := a.db.GetMongoSession()
	t := &proto.Thing{}
	if err == nil {
		defer db.Close()
		err = db.DB("").C(collectionName).Find(lib.M{"_id": req.Thing}).Select(bson.M{
			"secret": 0,
		}).One(&t)
	}
	if err != nil {
		resp.Success = false
		resp.Error = err.Error()
	} else {
		resp.Success = true
		resp.Thing = t
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
		err = db.DB("").C(collectionName).Find(lib.M{"project": req.Project}).Select(bson.M{
			"secret": 0,
		}).All(&things)
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

func (a *app) projectDelete(_, reply string, req *proto.ProjectDeleteRequest) {
	log.Printf("projectDelete req: %+v", req)
	resp := &proto.ProjectDeleteResponse{}
	db, err := a.db.GetMongoSession()
	if err == nil {
		defer db.Close()
		_, err = db.DB("").C(collectionName).UpdateAll(lib.M{
			"project": req.Project,
		}, lib.M{
			"$set": lib.M{
				"project": "",
			},
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

func (a *app) deleteThing(_, reply string, req *proto.DeleteThingRequest) {
	log.Printf("projectDelete req: %+v", req)
	resp := &proto.DeleteThingResponse{}
	db, err := a.db.GetMongoSession()
	if err == nil {
		defer db.Close()
		_, err = db.DB("").C(collectionName).UpdateAll(lib.M{
			"_id": req.Thing,
		}, lib.M{
			"$set": lib.M{
				"project": "",
			},
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

/* func (a *app) updateChannel(_, reply string, req *proto.UpdateChannelRequest) {
	log.Printf("create req: %+v", req)
	resp := &proto.UpdateChannelResponse{}
	db, err := a.db.GetMongoSession()
	if err == nil {
		defer db.Close()
		update := lib.M{}
		switch req.Channel.Unit {
		case proto.Unit_BOOL:
			update["channels.$.boolValue"] = req.Channel.BoolValue
		case proto.Unit_NUMBER:
			update["channels.$.floatValue"] = req.Channel.FloatValue
		case proto.Unit_STRING:
			update["channels.$.stringValue"] = req.Channel.StringValue
		case proto.Unit_DATA:
			update["channels.$.dataValue"] = req.Channel.DataValue
		}
		err = db.DB("").C(collectionName).Update(lib.M{
			"_id":          req.Channel.Thing,
			"channels._id": req.Channel.Id,
		}, bson.M{
			"$set": update,
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
} */

func (a *app) updateChannels(_, reply string, req *proto.UpdateChannelsRequest) {
	log.Printf("create req: %+v", req)
	resp := &proto.UpdateChannelsResponse{}
	db, err := a.db.GetMongoSession()
	if err == nil {
		defer db.Close()
		for _, channel := range req.Channels {
			update := lib.M{}
			switch channel.Unit {
			case proto.Unit_BOOL:
				update["channels.$.boolValue"] = channel.BoolValue
			case proto.Unit_NUMBER:
				update["channels.$.floatValue"] = channel.FloatValue
			case proto.Unit_STRING:
				update["channels.$.stringValue"] = channel.StringValue
			case proto.Unit_DATA:
				update["channels.$.dataValue"] = channel.DataValue
			}
			err = db.DB("").C(collectionName).Update(lib.M{
				"_id":          req.Thing,
				"channels._id": channel.Id,
			}, bson.M{
				"$set": update,
			})

			if err != nil {
				break
			} else {
				a.insertSeries(req.Thing, req.Channels, db)
				rule := &ruleProto.Rule{
					Thing:       req.Thing,
					FloatValue:  channel.FloatValue,
					BoolValue:   channel.BoolValue,
					StringValue: channel.StringValue,
					DataValue:   channel.DataValue,
					Unit:        int32(channel.Unit),
					Channel:     channel.Id,
				}

				ruleReq := &ruleProto.CheckThingRuleRequest{
					Update: rule,
				}

				// log.Println("rule topic", ruleTopics.CheckThingRule)

				err = a.eb.SendMessage(ruleTopics.CheckThingRule, ruleReq)
				if err != nil {
					log.Println("rule err", err)
					err = nil
				}
			}
		}

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

func (a *app) updateThingsChannels(_, reply string, req *proto.UpdateThingsChannelsRequest) {
	log.Printf("create req: %+v", req)
	resp := &proto.UpdateThingsChannelsResponse{}
	db, err := a.db.GetMongoSession()
	if err == nil {
		defer db.Close()
		for _, thing := range req.Things {
			currRes := &proto.Result{
				Success: true,
				Thing:   thing.Id,
			}
			mqttResp := &mqttProto.UpdateThingResponse{}
			updates := map[string]lib.M{}

			req := &mqttProto.UpdateThingRequest{
				Thing: &mqttProto.Thing{
					Id:       thing.Id,
					Channels: []*mqttProto.Channel{},
				},
			}

			for _, channel := range thing.Channels {
				if !channel.IsSensor {
					update := lib.M{}
					thingChannel := &mqttProto.Channel{
						Id:   channel.Id,
						Unit: int32(channel.Unit),
					}
					switch channel.Unit {
					case proto.Unit_BOOL:
						update["channels.$.boolValue"] = channel.BoolValue
						thingChannel.BoolValue = channel.BoolValue
					case proto.Unit_NUMBER:
						update["channels.$.floatValue"] = channel.FloatValue
						thingChannel.FloatValue = channel.FloatValue
					case proto.Unit_STRING:
						update["channels.$.stringValue"] = channel.StringValue
						thingChannel.StringValue = channel.StringValue
					case proto.Unit_DATA:
						update["channels.$.dataValue"] = channel.DataValue
						thingChannel.DataValue = channel.DataValue
					}

					updates[channel.Id] = update

					req.Thing.Channels = append(req.Thing.Channels, thingChannel)
				}
			}

			err = a.eb.RequestMessage(mqttTopics.UpdateThing, req, mqttResp, lib.DefaultTimeout)
			if err == nil {
				if mqttResp.Success {
					for id, update := range updates {
						err = db.DB("").C(collectionName).Update(lib.M{
							"_id":          thing.Id,
							"channels._id": id,
						}, bson.M{
							"$set": update,
						})

						if err != nil {
							break
						} else {
							a.insertSeries(thing.Id, thing.Channels, db)
						}
					}
				} else {
					err = errors.New(mqttResp.Error)
				}
			}

			if err != nil {
				currRes.Success = false
				currRes.Error = err.Error()
				err = nil
			} else {
				go func(a *app, thing *proto.ThingChannels) {
					for _, channel := range thing.Channels {
						rule := &ruleProto.Rule{
							Thing:       thing.Id,
							FloatValue:  channel.FloatValue,
							BoolValue:   channel.BoolValue,
							StringValue: channel.StringValue,
							DataValue:   channel.DataValue,
							Unit:        int32(channel.Unit),
							Channel:     channel.Name,
						}

						ruleReq := ruleProto.CheckThingRuleRequest{
							Update: rule,
						}

						a.eb.SendMessage(ruleTopics.CheckThingRule, ruleReq)
					}
				}(a, thing)
			}

			resp.Results = append(resp.Results, currRes)
		}
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

func (a *app) generateAssignThing(_, reply string, req *proto.GenerateAssignThingRequest) {
	log.Printf("create req: %+v", req)
	resp := &proto.GenerateAssignThingResponse{}
	db, err := a.db.GetMongoSession()
	if err == nil {
		defer db.Close()
		var n int
		n, err = db.DB("").C(collectionName).Find(bson.M{
			"_id":     req.Id,
			"project": "",
		}).Count()

		if n < 1 {
			err = errors.New("not found")
		}

		if err == nil {
			t, _ := uuid.NewRandom()
			token := t.String()

			resp.Token = token

			err = a.cacheClient.Set(token, fmt.Sprintf("%s-%s", req.Id, req.Project), lib.TokenTimeout).Err()
		}
	}

	if err != nil {
		resp.Success = false
		resp.Error = err.Error()
	} else {
		resp.Success = true
	}

	log.Println("new token", fmt.Sprintf("%s-%s", req.Id, req.Project))

	log.Printf("Request: %+v, Resposne: %+v", req, resp)
	if reply != "" {
		a.eb.SendMessage(reply, resp)
	}
}

func (a *app) assignThing(_, reply string, req *proto.AssignThingRequest) {
	log.Printf("create req: %+v", req)
	resp := &proto.AssignThingResponse{}
	data, err := a.cacheClient.Get(req.Token).Result()
	if err == nil {
		id := strings.Split(data, "-")[0]
		project := strings.Split(data, "-")[1]

		log.Println(project, id, "PIPI")

		var db *mgo.Session

		db, err = a.db.GetMongoSession()
		if err == nil {
			defer db.Close()
			t := &proto.Thing{}
			var change *mgo.ChangeInfo
			change, err = db.DB("").C(collectionName).Find(bson.M{
				"_id":     id,
				"project": "",
				"key":     req.Key,
			}).Apply(mgo.Change{
				Update: bson.M{
					"$set": bson.M{
						"project": project,
					},
				},
			}, t)
			log.Printf("Change %+v, err %v", change, err)
			if err == nil {
				if change.Matched == 0 || change.Updated == 0 {
					err = errors.New("wrong key id pair")
				}
			} else {
				err = errors.New("wrong key id pair")
			}

		}
	}

	log.Println(err)

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

func (a *app) deassignThing(_, reply string, req *proto.DeassignThingRequest) {
	log.Printf("create req: %+v", req)
	resp := &proto.DeassignThingResponse{}
	db, err := a.db.GetMongoSession()
	if err == nil {
		defer db.Close()
		err = db.DB("").C(collectionName).Update(bson.M{
			"_id":     req.Id,
			"project": req.Project,
		}, bson.M{
			"$set": bson.M{
				"project": "",
			},
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

func (a *app) thingSeries(_, reply string, req *proto.ThingSeriesRequest) {
	log.Printf("series req: %+v", req)
	resp := &proto.ThingSeriesResponse{}
	values := []*proto.Series{}
	db, err := a.db.GetMongoSession()
	if err == nil {
		defer db.Close()
		err = db.DB("").C(timeSeriesCollectionName).Find(bson.M{
			"thing": req.Id,
		}).All(&values)
		if err == mgo.ErrNotFound {
			err = nil
		}
	}
	if err != nil {
		resp.Success = false
		resp.Error = err.Error()
	} else {
		resp.Success = true
		resp.Values = values
	}

	log.Printf("Request: %+v, Resposne: %+v", req, resp)
	if reply != "" {
		a.eb.SendMessage(reply, resp)
	}
}

/* {
	"_id": "5cffbd85e64d1e22a6114e6a"
},
{
	$set: {
		"channels.$[aa].boolValue": false,
		"channels.$[bb].floatValue": 58.0,
	}
},
{
	upsert: false,
	arrayFilters: [
		{ "aa._id": "light1" },
		{ "bb._id": "fan1" },
	]
} */
