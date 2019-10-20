package main

import (
	"log"
	"thingz-server/lib"
	proto "thingz-server/rule/proto"
	thingProto "thingz-server/thing/proto"
	thingTopics "thingz-server/thing/topics"
)

func (a *app) createRule(_, reply string, req *proto.CreateRuleRequest) {
	log.Printf("create req: %+v", req)
	resp := &proto.CreateRuleResponse{}
	db, err := a.db.GetMongoSession()
	if err == nil {
		defer db.Close()
		req.Rule.Id = lib.NewObjectID()
		err = db.DB("").C(collectionName).Insert(req.Rule)
	}

	if err != nil {
		resp.Success = false
		resp.Error = err.Error()
	} else {
		resp.Success = true
		resp.Id = req.Rule.Id
	}

	log.Printf("Request: %+v, Resposne: %+v", req, resp)
	if reply != "" {
		a.eb.SendMessage(reply, resp)
	}
}

func (a *app) deleteRule(_, reply string, req *proto.DeleteRuleRequest) {
	log.Printf("deleteRule req: %+v", req)
	resp := &proto.DeleteRuleResponse{}
	db, err := a.db.GetMongoSession()
	if err == nil {
		defer db.Close()
		err = db.DB("").C(collectionName).Remove(lib.M{
			"_id": req.Rule,
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

func (a *app) projectRules(_, reply string, req *proto.ProjectRulesRequest) {
	log.Printf("create req: %+v", req)
	resp := &proto.ProjectRulesResponse{}
	rules := []*proto.Rule{}
	db, err := a.db.GetMongoSession()
	if err == nil {
		defer db.Close()
		db.DB("").C(collectionName).Find(lib.M{"project": req.Project}).All(&rules)
	}
	if err != nil {
		resp.Success = false
		resp.Error = err.Error()
	} else {
		resp.Success = true
		resp.Rules = rules
	}

	log.Printf("Request: %+v, Resposne: %+v", req, resp)
	if reply != "" {
		a.eb.SendMessage(reply, resp)
	}
}

func (a *app) checkThingRule(_, reply string, req *proto.CheckThingRuleRequest) {
	log.Printf("create req: %+v", req)
	resp := &proto.CheckThingRuleResponse{}
	rules := []*proto.Rule{}
	db, err := a.db.GetMongoSession()
	if err == nil {
		defer db.Close()
		err = db.DB("").C(collectionName).Find(lib.M{
			"thing":   req.Update.Thing,
			"channel": req.Update.Channel,
		}).All(&rules)
	}
	if err != nil {
		log.Println(err)
		resp.Success = false
		resp.Error = err.Error()
	} else {
		log.Println("checking rules", rules)
		matchedRules := []*proto.Rule{}

		for _, rule := range rules {
			log.Printf("checking %v with %v", rule.Unit, req.Update.Unit)
			switch thingProto.Unit(rule.Unit) {
			case thingProto.Unit_BOOL:
				switch rule.Operation {
				case proto.Operation_EQUAL:
					if rule.BoolValue == req.Update.BoolValue {
						matchedRules = append(matchedRules, rule)
					}
				case proto.Operation_NOTEQUAL:
					if rule.BoolValue != req.Update.BoolValue {
						matchedRules = append(matchedRules, rule)
					}
				}
			case thingProto.Unit_NUMBER:
				switch rule.Operation {
				case proto.Operation_EQUAL:
					if rule.FloatValue == req.Update.FloatValue {
						matchedRules = append(matchedRules, rule)
					}
				case proto.Operation_NOTEQUAL:
					if rule.FloatValue != req.Update.FloatValue {
						matchedRules = append(matchedRules, rule)
					}
				case proto.Operation_LESSTHAN:
					log.Println(rule.FloatValue, ">", req.Update.FloatValue, ": ", rule.FloatValue > req.Update.FloatValue)
					if rule.FloatValue > req.Update.FloatValue {
						matchedRules = append(matchedRules, rule)
					}
				case proto.Operation_LESSTHANEQUAL:
					if rule.FloatValue >= req.Update.FloatValue {
						matchedRules = append(matchedRules, rule)
					}
				case proto.Operation_GREATERTHAN:
					if rule.FloatValue < req.Update.FloatValue {
						matchedRules = append(matchedRules, rule)
					}
				case proto.Operation_GREATERTHANEQUAL:
					if rule.FloatValue <= req.Update.FloatValue {
						matchedRules = append(matchedRules, rule)
					}
				}
			case thingProto.Unit_DATA:
				switch rule.Operation {
				case proto.Operation_EQUAL:
					if rule.DataValue == req.Update.DataValue {
						matchedRules = append(matchedRules, rule)
					}
				case proto.Operation_NOTEQUAL:
					if rule.DataValue != req.Update.DataValue {
						matchedRules = append(matchedRules, rule)
					}
				case proto.Operation_LESSTHAN:
					if rule.DataValue > req.Update.DataValue {
						matchedRules = append(matchedRules, rule)
					}
				case proto.Operation_LESSTHANEQUAL:
					if rule.DataValue >= req.Update.DataValue {
						matchedRules = append(matchedRules, rule)
					}
				case proto.Operation_GREATERTHAN:
					if rule.DataValue < req.Update.DataValue {
						matchedRules = append(matchedRules, rule)
					}
				case proto.Operation_GREATERTHANEQUAL:
					if rule.DataValue <= req.Update.DataValue {
						matchedRules = append(matchedRules, rule)
					}
				}
			case thingProto.Unit_STRING:
				switch rule.Operation {
				case proto.Operation_EQUAL:
					if rule.StringValue == req.Update.StringValue {
						matchedRules = append(matchedRules, rule)
					}
				case proto.Operation_NOTEQUAL:
					if rule.StringValue != req.Update.StringValue {
						matchedRules = append(matchedRules, rule)
					}
				case proto.Operation_LESSTHAN:
					if rule.StringValue > req.Update.StringValue {
						matchedRules = append(matchedRules, rule)
					}
				case proto.Operation_LESSTHANEQUAL:
					if rule.StringValue >= req.Update.StringValue {
						matchedRules = append(matchedRules, rule)
					}
				case proto.Operation_GREATERTHAN:
					if rule.StringValue < req.Update.StringValue {
						matchedRules = append(matchedRules, rule)
					}
				case proto.Operation_GREATERTHANEQUAL:
					if rule.StringValue <= req.Update.StringValue {
						matchedRules = append(matchedRules, rule)
					}
				}
			}
		}

		updateReq := &thingProto.UpdateThingsChannelsRequest{
			Things: []*thingProto.ThingChannels{},
		}

		for _, rule := range matchedRules {
			log.Println("matched rule", rule)
			switch rule.TriggerType {
			case proto.TriggerType_COMMAND:
				cmds := rule.TriggerCommands

				for _, cmd := range cmds {
					thingChannel := &thingProto.ThingChannels{
						Id:       cmd.Thing,
						Channels: []*thingProto.Channel{},
					}

					for _, channel := range cmd.Channels {
						thingChannel.Channels = append(thingChannel.Channels, &thingProto.Channel{
							Id:          channel.Id,
							StringValue: channel.StringValue,
							DataValue:   channel.DataValue,
							FloatValue:  channel.FloatValue,
							BoolValue:   channel.BoolValue,
							Unit:        thingProto.Unit(channel.Unit),
						})
					}

					updateReq.Things = append(updateReq.Things, thingChannel)
				}

				updateRes := &thingProto.UpdateThingsChannelsResponse{}

				err = a.eb.RequestMessage(thingTopics.UpdateThingsChannels, updateReq, updateRes, lib.DefaultTimeout)
				log.Println("exec rule", updateReq, thingTopics.UpdateChannels, err)
			}
		}
		resp.Success = true
	}

	log.Printf("Request: %+v, Resposne: %+v", req, resp)
	if reply != "" {
		a.eb.SendMessage(reply, resp)
	}
}
