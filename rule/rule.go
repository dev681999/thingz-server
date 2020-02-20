package main

import (
	"context"
	"encoding/json"
	"fmt"
	"thingz-server/lib"
	proto "thingz-server/rule/proto"
	thingProto "thingz-server/thing/proto"
	thingTopics "thingz-server/thing/topics"

	log "github.com/sirupsen/logrus"

	"github.com/dgraph-io/dgo/v2/protos/api"
)

/* func (a *app) createRule(_, reply string, req *proto.CreateRuleRequest) {
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
} */

func (a *app) findThingChUID(id string) (string, error) {
	ctx := context.Background()
	q := `{
		devices(func: eq(thingId,` + id + `)) {
			uid
		}
	}`
	/* vars := map[string]string{
		"$id": id,
	} */
	log.Println("firing for", q)
	response, err := a.dg.NewReadOnlyTxn().Query(ctx, q)
	if err != nil {
		return "", err
	}

	type device struct {
		UID string `json:"uid,omitempty"`
		// ID  string `json:"id,omitempty"`
	}

	resData := &struct {
		Devices []device `json:"devices,omitempty"`
	}{}

	log.Println("response", response.Json)

	err = json.Unmarshal(response.Json, resData)
	if err != nil {
		return "", err
	}

	return resData.Devices[0].UID, nil
}

func resolvePhysicalID(r *proto.Rule, m map[string]thingProto.Thing) {
	r.PhysicalId = m[r.Thing].PhysicalId
	for _, subRule := range r.Rules {
		resolvePhysicalID(subRule, m)
	}
}

func (a *app) createRule(_, reply string, req *proto.CreateRuleRequest) {
	log.Printf("create req: %+v", req)

	resp := &proto.CreateRuleResponse{}
	err := req.Rule.ResolveThingChLinks(a.findThingChUID)
	if err == nil {
		ctx := context.Background()
		mu := &api.Mutation{
			CommitNow: true,
		}
		thingsMap := map[string]thingProto.Thing{}
		err = json.Unmarshal([]byte(req.ThingsMap), &thingsMap)
		if err == nil {
			resolvePhysicalID(req.Rule, thingsMap)

			for _, cmd := range req.Rule.TriggerCommands {
				cmd.PhysicalId = thingsMap[cmd.Thing].PhysicalId
				log.Println("trig thing", cmd.Thing, "phycsial id", cmd.PhysicalId)
			}

			req.Rule.Evaluate(thingsMap)
			req.Rule.AssignRootID(req.Rule.GetId())

			req.Rule.IsRoot = true

			dgraphRule := req.Rule.GetWithDType()
			pb, err := json.Marshal(dgraphRule)
			if err == nil {
				log.Println("sendig dgraph", string(pb))
				mu.SetJson = pb
				_, err = a.dg.NewTxn().Mutate(ctx, mu)
				/* if err == nil {

				} */
			}
		}
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

/* func (a *app) deleteRule(_, reply string, req *proto.DeleteRuleRequest) {
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
*/

func (a *app) deleteRule(_, reply string, req *proto.DeleteRuleRequest) {
	log.Printf("deleteRule req: %+v", req)
	resp := &proto.DeleteRuleResponse{}
	ids := &struct {
		Rules []map[string]string `json:"rules"`
	}{}

	ctx := context.Background()
	dRes, err := a.dg.NewReadOnlyTxn().Query(ctx, `{
		rules	(func: eq(root, `+req.Rule+`)) {
			uid
		}
	}`)

	if err == nil {
		err = json.Unmarshal(dRes.Json, ids)
		if err == nil {
			log.Printf("got ids: %+v", ids)
			pb, err := json.Marshal(ids.Rules)
			if err == nil {
				mu := &api.Mutation{
					CommitNow:  true,
					DeleteJson: pb,
				}

				_, err = a.dg.NewTxn().Mutate(ctx, mu)
				if err == nil {
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

func (a *app) projectRules(_, reply string, req *proto.ProjectRulesRequest) {
	log.Printf("projectRules req: %+v", req)
	resp := &proto.ProjectRulesResponse{}
	rules := []*proto.Rule{}
	triggerCommands := []*proto.Rule{}
	ctx := context.Background()
	dRes, err := a.dg.NewReadOnlyTxn().Query(ctx, `{
		v as var(func: eq(project, `+req.Project+`)) @filter(eq(isRoot, true))
		rules(func: uid(v)) @recurse {
			id
			physicalId
			operation
			val
			vals
			unit
			boolVal
			floatValue
			thing
			channel
			project
			rules: ~rules
			uid
		}

		triggerCommands(func: uid(v)) {
			id
			triggerCommands {
				thing
				physicalId
				channels {
					id
					boolValue
				}
			}
		}
	}`)

	if err == nil {
		projectRules := &struct {
			Rules           []*proto.Rule `json:"rules"`
			TriggerCommands []*proto.Rule `json:"triggerCommands"`
		}{}
		err = json.Unmarshal(dRes.Json, projectRules)
		if err == nil {
			rules = projectRules.Rules
			triggerCommands = projectRules.TriggerCommands
			for idx, r := range rules {
				r.TriggerCommands = triggerCommands[idx].TriggerCommands
			}
		}
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

/* func (a *app) checkThingRule(_, reply string, req *proto.CheckThingRuleRequest) {
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
} */

func (a *app) checkThingRule(_, reply string, req *proto.CheckThingRuleRequest) {
	log.Printf("check thing rule req: %+v", req)
	resp := &proto.CheckThingRuleResponse{}
	// db, err := a.db.GetMongoSession()
	ctx := context.Background()
	q := `{
		var(func: eq(thingId,` + req.Update.Thing + "-" + req.Update.Channel + `)) {
			~thingChLink {
				ID AS uid
			}
		}
		
		rules(func: uid(ID)) @recurse {
			id: v as id
			physicalId
			operation
			val
			vals
			unit
			boolValue
			floatValue
			thing
			channel
			project
			rules: ~rules
			uid
		}

		triggerRule(func: uid(v)) {
			id
			triggerCommands {
				thing
				physicalId
				channels {
					id
					boolValue
					floatValue
					unit
				}
			}
		}
	}`
	result := &struct {
		// Rules []struct {
		Rules       []*proto.Rule `json:"rules,omitempty"`
		TriggerRule []*proto.Rule `json:"triggerRule,omitempty"`
		// } `json:"rules,omitempty"`
	}{}
	rules := map[string]*proto.Rule{}
	dResp, err := a.dg.NewTxn().Query(ctx, q)
	if err == nil {
		log.Println("res", string(dResp.GetJson()))
		err = json.Unmarshal(dResp.GetJson(), result)
		if err == nil {
			updatedRuleNodes := map[string]*proto.Rule{}
			for idx, rule := range result.Rules {
				status := rule.Check(&thingProto.Channel{
					BoolValue:  req.Update.BoolValue,
					FloatValue: req.Update.FloatValue,
					Unit:       thingProto.Unit(req.Update.Unit),
				})
				rule.Val = status
				// update := true
				// var subRule *proto.Rule

				updatedRuleNodes[rule.Id] = rule

				for len(rule.Rules) > 0 /* && len(subRule.Rules) > 0 */ {
					vals := map[string]bool{}
					subRule := rule.Rules[0]

					if r, ok := updatedRuleNodes[subRule.Id]; ok {
						// log.Println("match", rule.Id, r.Id, "-", subRule.Vals, ":::", r.Vals)
						subRule.Vals = r.Vals
					}

					json.Unmarshal([]byte(subRule.Vals), &vals)

					vals[rule.Id] = rule.Val
					b, _ := json.Marshal(vals)
					subRule.Vals = string(b)
					status := true

					for _, v := range vals {
						if (subRule.Operation == proto.Operation_AND && status) || (subRule.Operation == proto.Operation_OR && !status) {
							status = v
						}
					}

					subRule.Val = status
					updatedRuleNodes[subRule.Id] = subRule
					rule = subRule
				}

				if _, ok := rules[rule.Id]; ok {
					rule.TriggerCommands = result.TriggerRule[idx].TriggerCommands
					rules[rule.Id] = rule
				} else if rule.Val {
					rule.TriggerCommands = result.TriggerRule[idx].TriggerCommands
					rules[rule.Id] = rule
				}
			}

			updateReq := []map[string]interface{}{}

			for _, update := range updatedRuleNodes {
				updateReq = append(updateReq, map[string]interface{}{
					"uid":  update.Uid,
					"val":  update.Val,
					"vals": update.Vals,
				})

				// log.Printf("updates: id: %+v & update: %+v", id, update.Val)
			}

			log.Println("updating", updateReq)

			pb, err := json.Marshal(updateReq)
			if err == nil {
				ctx = context.Background()
				dResp, err := a.dg.NewTxn().Mutate(ctx, &api.Mutation{
					SetJson:   pb,
					CommitNow: true,
				})
				if err == nil {
					log.Println("update response", dResp.GetUids())
				}
			}

		}
	}

	if err != nil {
		log.Println(err)
		resp.Success = false
		resp.Error = err.Error()
	} else {
		updateReq := &thingProto.UpdateThingsChannelsRequest{
			Things: []*thingProto.ThingChannels{},
		}

		log.Println("cheching rules", len(rules))

		for _, rule := range rules {
			log.Println("matched rule", rule.Val)
			for _, cmd := range rule.TriggerCommands {
				log.Println("cmd", cmd.Thing)
				for _, ch := range cmd.Channels {
					log.Println("ch", *ch)
				}
			}
			switch rule.TriggerType {
			case proto.TriggerType_COMMAND:
				cmds := rule.TriggerCommands

				for _, cmd := range cmds {
					thingChannel := &thingProto.ThingChannels{
						Id:         cmd.Thing,
						Channels:   []*thingProto.Channel{},
						PhysicalId: cmd.PhysicalId,
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
			}
		}

		if len(updateReq.Things) > 0 {
			updateRes := &thingProto.UpdateThingsChannelsResponse{}

			err = a.eb.RequestMessage(thingTopics.UpdateThingsChannels, updateReq, updateRes, lib.DefaultTimeout)
			log.Println("exec rule", updateReq, thingTopics.UpdateChannels, err)
		}

		resp.Success = true
	}

	log.Printf("Request: %+v, Resposne: %+v", req, resp)
	if reply != "" {
		a.eb.SendMessage(reply, resp)
	}
}

func (a *app) createThingLink(_, reply string, req *proto.CreateThingLinkRequest) {
	log.Printf("createThingLink %+v", req.Thing)
	resp := &proto.CreateThingLinkResponse{}
	ctx := context.Background()
	mu := &api.Mutation{
		CommitNow: true,
	}
	data := []map[string]string{}

	for _, ch := range req.Channels {
		data = append(data, map[string]string{
			"thingId": fmt.Sprintf("%v-%v", req.Thing, ch),
		})
	}

	pb, err := json.Marshal(data)
	if err == nil {
		mu.SetJson = pb
		_, err = a.dg.NewTxn().Mutate(ctx, mu)
		/* if err == nil {

		} */
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
