package main

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"thingz-server/lib"
	mqttProto "thingz-server/mqtt/proto"
	mqttTopics "thingz-server/mqtt/topics"
	ruleProto "thingz-server/rule/proto"
	ruleTopics "thingz-server/rule/topics"
	proto "thingz-server/thing/proto"

	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/google/uuid"
)

// go-run --thing VOEZGHDB --channel value --valType 2 --stringVal '18.5213,73.9523'

var (
	thingTypes     = []thingInfo{}
	thingTypeToInt = map[string]int32{}
)

const letterBytes = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func getNewID(c *redis.Client) string {
	for {
		id := randStringBytes(8)
		_, err := c.Get(fmt.Sprintf("phy_id_%s", id)).Result()
		if err == redis.Nil {
			c.Set(fmt.Sprintf("phy_id_%s", id), "TAKEN", 0)
			return id
		}
	}
}

func createHubChannels() []*proto.Channel {
	channels := []*proto.Channel{}

	return channels
}

func (a *app) createThing(_, reply string, req *proto.CreateThingRequest) {
	log.Printf("create req: %+v", req)
	resp := &proto.CreateThingResponse{}
	db, err := a.db.GetMongoSession()
	if err == nil {
		defer db.Close()
		if req.Thing.Name == "" {
			req.Thing.Name = thingTypes[req.Thing.Type].Name
		}

		req.Thing.Channels = []*proto.Channel{}
		for _, ch := range thingTypes[req.Thing.Type].Channels {
			req.Thing.Channels = append(req.Thing.Channels, &proto.Channel{
				Id:       ch.Id,
				Name:     ch.Name,
				Unit:     ch.Unit,
				Type:     ch.Type,
				IsSensor: ch.IsSensor,
			})
		}
		/* switch req.Thing.Type {
		case thingTypeToInt["HUB"]:
			req.Thing.Channels = createHubChannels()
		default:
			req.Thing.Channels = []*proto.Channel{}
			for _, ch := range thingTypes[req.Thing.Type].Channels {
				req.Thing.Channels = append(req.Thing.Channels, &proto.Channel{
					Id:       ch.Id,
					Name:     ch.Name,
					Unit:     ch.Unit,
					Type:     ch.Type,
					IsSensor: ch.IsSensor,
				})
			}
		} */

		// for {
		req.Thing.PhysicalId = getNewID(a.cacheClient)
		// req.Thing.Id = bson.NewObjectId().Hex()
		req.Thing.Id = req.Thing.PhysicalId
		err = db.DB("").C(collectionName).Insert(req.Thing)
		// }
	}
	if err != nil {
		resp.Success = false
		resp.Error = err.Error()
	} else {
		resp.Success = true
		resp.Id = req.Thing.Id
		channels := []string{}

		for _, ch := range req.Thing.Channels {
			channels = append(channels, ch.GetId())
		}

		resp.Channels = channels
	}

	log.Printf("Request: %+v, Resposne: %+v", req, resp)
	if reply != "" {
		a.eb.SendMessage(reply, resp)
	}
}

func (a *app) createThings(_, reply string, req *proto.CreateThingsRequest) {
	log.Printf("create req: %+v", req)
	resp := &proto.CreateThingsResponse{
		Things: []*proto.ThingStrChannels{},
	}
	db, err := a.db.GetMongoSession()
	if err == nil {
		defer db.Close()

		physicalID := getNewID(a.cacheClient)
		for _, t := range req.Things {
			t.PhysicalId = physicalID
			if t.Name == "" {
				t.Name = thingTypes[t.Type].Name
			}

			t.Channels = []*proto.Channel{}
			for _, ch := range thingTypes[t.Type].Channels {
				t.Channels = append(t.Channels, &proto.Channel{
					Id:       ch.Id,
					Name:     ch.Name,
					Unit:     ch.Unit,
					Type:     ch.Type,
					IsSensor: ch.IsSensor,
				})
			}

			t.Id = bson.NewObjectId().Hex()
			err = db.DB("").C(collectionName).Insert(t)
			if err != nil {
				break
			}
		}
		/* switch req.Thing.Type {
		case thingTypeToInt["HUB"]:
			req.Thing.Channels = createHubChannels()
		default:
			req.Thing.Channels = []*proto.Channel{}
			for _, ch := range thingTypes[req.Thing.Type].Channels {
				req.Thing.Channels = append(req.Thing.Channels, &proto.Channel{
					Id:       ch.Id,
					Name:     ch.Name,
					Unit:     ch.Unit,
					Type:     ch.Type,
					IsSensor: ch.IsSensor,
				})
			}
		} */
	}
	if err != nil {
		resp.Success = false
		resp.Error = err.Error()
	} else {
		resp.Success = true
		for _, t := range req.Things {
			channels := []string{}

			for _, ch := range t.Channels {
				channels = append(channels, ch.GetId())
			}

			resp.Things = append(resp.Things, &proto.ThingStrChannels{
				Id:       t.Id,
				Channels: channels,
			})
		}
		// resp.Id = req.Thing.Id
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

func (a *app) updateThing(_, reply string, req *proto.UpdateThingRequest) {
	log.Printf("create req: %+v", req)
	resp := &proto.UpdateThingResponse{}
	db, err := a.db.GetMongoSession()
	if err == nil {
		defer db.Close()
		err = db.DB("").C(collectionName).Update(lib.M{
			"_id": req.Thing,
		}, bson.M{
			"$set": bson.M{
				"name":  req.Name,
				"group": req.Group,
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

func (a *app) updateChannelName(_, reply string, req *proto.UpdateChannelNameRequest) {
	log.Printf("create req: %+v", req)
	resp := &proto.UpdateChannelNameResponse{}
	db, err := a.db.GetMongoSession()
	if err == nil {
		defer db.Close()
		err = db.DB("").C(collectionName).Update(lib.M{
			"_id":         req.Thing,
			"channels.id": req.Channel,
		}, bson.M{
			"$set": bson.M{
				"channels.$.name": req.Name,
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
	log.Printf("updateThingsChannels req: %+v", req)
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
					Id:         thing.Id,
					PhysicalId: thing.PhysicalId,
					Channels:   []*mqttProto.Channel{},
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
							Channel:     channel.Id,
						}

						ruleReq := ruleProto.CheckThingRuleRequest{
							Update: rule,
						}

						log.Printf("sending to rule engine %+v", ruleReq)

						err := a.eb.SendMessage(ruleTopics.CheckThingRule, &ruleReq)
						if err != nil {
							log.Println(err)
						}
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
			"physicalId": req.Id,
			"$or": []bson.M{{
				"project": "",
			}, {
				"project": bson.M{
					"$exists": false,
				}},
			},
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
			// t := &proto.Thing{}
			var change *mgo.ChangeInfo
			change, err = db.DB("").C(collectionName).UpdateAll(bson.M{
				"physicalId": id,
				"$or": []bson.M{{
					"project": "",
				}, {
					"project": bson.M{
						"$exists": false,
					}},
				},
				// "key": req.Key,
			}, bson.M{
				"$set": bson.M{
					"project": project,
				},
			})
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
			"thing":   req.Id,
			"channel": req.Channel,
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

func (a *app) getThingsByIDs(_, reply string, req *proto.GetThingsByIDsRequest) {
	log.Printf("getThingsByIDs req: %+v", req)
	resp := &proto.GetThingsByIDsResponse{}
	things := []*proto.Thing{}
	query := lib.M{}

	if len(req.Ids) > 0 {
		query = lib.M{"_id": lib.M{"$in": req.Ids}}
	}

	db, err := a.db.GetMongoSession()
	if err == nil {
		defer db.Close()
		err = db.DB("").C(collectionName).Find(query).Select(bson.M{
			"secret":  0,
			"project": 0,
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

func (a *app) updateThingConfig(_, reply string, req *proto.UpdateThingConfigRequest) {
	log.Printf("updateThingConfig req: %+v", req)
	resp := &proto.UpdateThingConfigResponse{}
	db, err := a.db.GetMongoSession()
	if err == nil {
		defer db.Close()
		err = db.DB("").C(collectionName).Update(lib.M{
			"_id": req.Thing,
		}, lib.M{
			"$set": lib.M{
				"config": req.Config,
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

func (a *app) getThingTypes(_, reply string, req *proto.GetThingTypesRequest) {
	types := []string{}
	for _, tt := range thingTypes {
		types = append(types, tt.Name)
	}
	if reply != "" {
		a.eb.SendMessage(reply, &proto.GetThingTypesResponse{
			Success: true,
			Types:   types,
		})
	}
}

func (a *app) projectGroupThings(_, reply string, req *proto.ProjectGroupThingsRequest) {
	log.Printf("create req: %+v", req)
	resp := &proto.ProjectGroupThingsResponse{}
	things := []*proto.Thing{}
	db, err := a.db.GetMongoSession()
	if err == nil {
		defer db.Close()
		err = db.DB("").C(collectionName).Find(lib.M{"project": req.Project, "group": req.Group}).Select(bson.M{
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
