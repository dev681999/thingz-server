package ruleproto

import (
	"encoding/json"
	thingProto "thingz-server/thing/proto"

	"github.com/globalsign/mgo/bson"
)

// RuleWithDType add Dgraph Type to Rule
type RuleWithDType struct {
	*Rule
	DType []string         `json:"dgraph.type,omitempty"`
	Rules []*RuleWithDType `json:"rules,omitempty"`
}

// ThingChLinkWithDType add Dgraph Type to ThingChLink
type ThingChLinkWithDType struct {
	*ThingChLink
	DType []string `json:"dgraph.type,omitempty"`
}

// TriggerCommandWithDType add Dgraph Type to TriggerCommand
type TriggerCommandWithDType struct {
	*TriggerCommand
	Channels []*ChannelWithDType `json:"channels"`
	DType    []string            `json:"dgraph.type,omitempty"`
}

// ChannelWithDType add Dgraph Type to Channel
type ChannelWithDType struct {
	*Channel
	DType []string `json:"dgraph.type,omitempty"`
}

// Evaluate the rule
func (rule *Rule) Evaluate(things map[string]thingProto.Thing) (string, bool) {
	if rule.Id == "" {
		rule.Id = bson.NewObjectId().Hex()
	}

	res := true
	ch := &thingProto.Channel{}

	if rule.Operation != Operation_AND && rule.Operation != Operation_OR {
		for _, c := range things[rule.Thing].Channels {
			if c.Id == rule.Channel {
				ch = c
				break
			}
		}
		res = rule.Check(ch)
	} else {
		switch rule.Operation {
		case Operation_AND:
			_, res = rule.Rules[0].Evaluate(things)
			vals := map[string]bool{}
			for _, subRule := range rule.Rules {
				rID, rRes := subRule.Evaluate(things)
				vals[rID] = rRes
				if res {
					res = rRes
				}
			}
			b, _ := json.Marshal(vals)
			rule.Vals = string(b)
		case Operation_OR:
			_, res = rule.Rules[0].Evaluate(things)
			vals := map[string]bool{}
			for _, subRule := range rule.Rules {
				rID, rRes := subRule.Evaluate(things)
				vals[rID] = rRes
				if !res {
					res = rRes
				}
			}
			b, _ := json.Marshal(vals)
			rule.Vals = string(b)
		}
	}

	rule.Val = res

	return rule.Id, res
}

// Check current rule
func (rule *Rule) Check(ch *thingProto.Channel) bool {
	var res bool

	switch thingProto.Unit(rule.Unit) {
	case thingProto.Unit_BOOL:
		switch rule.Operation {
		case Operation_EQUAL:
			res = rule.BoolValue == ch.BoolValue
		case Operation_NOTEQUAL:
			res = rule.BoolValue != ch.BoolValue
		}
	case thingProto.Unit_NUMBER:
		switch rule.Operation {
		case Operation_EQUAL:
			res = rule.FloatValue == ch.FloatValue
		case Operation_NOTEQUAL:
			res = rule.FloatValue != ch.FloatValue
		case Operation_LESSTHAN:
			// log.Println(rule.FloatValue, ">", ch.FloatValue, ": ", rule.FloatValue > ch.FloatValue)
			res = rule.FloatValue > ch.FloatValue
		case Operation_LESSTHANEQUAL:
			res = rule.FloatValue >= ch.FloatValue
		case Operation_GREATERTHAN:
			res = rule.FloatValue < ch.FloatValue
		case Operation_GREATERTHANEQUAL:
			res = rule.FloatValue <= ch.FloatValue
		}
	case thingProto.Unit_DATA:
		switch rule.Operation {
		case Operation_EQUAL:
			res = rule.DataValue == ch.DataValue
		case Operation_NOTEQUAL:
			res = rule.DataValue != ch.DataValue
		case Operation_LESSTHAN:
			res = rule.DataValue > ch.DataValue
		case Operation_LESSTHANEQUAL:
			res = rule.DataValue >= ch.DataValue
		case Operation_GREATERTHAN:
			res = rule.DataValue < ch.DataValue
		case Operation_GREATERTHANEQUAL:
			res = rule.DataValue <= ch.DataValue
		}
	case thingProto.Unit_STRING:
		switch rule.Operation {
		case Operation_EQUAL:
			res = rule.StringValue == ch.StringValue
		case Operation_NOTEQUAL:
			res = rule.StringValue != ch.StringValue
		case Operation_LESSTHAN:
			res = rule.StringValue > ch.StringValue
		case Operation_LESSTHANEQUAL:
			res = rule.StringValue >= ch.StringValue
		case Operation_GREATERTHAN:
			res = rule.StringValue < ch.StringValue
		case Operation_GREATERTHANEQUAL:
			res = rule.StringValue <= ch.StringValue
		}
	}

	return res
}

// GetAllThings return all things id in rule tree
func (rule *Rule) GetAllThings() []string {
	things := []string{}
	thingsMap := map[string]bool{}

	if rule.Thing != "" {
		things = append(things, rule.Thing)
	}

	for _, subRule := range rule.Rules {
		things = append(things, subRule.GetAllThings()...)
	}

	for _, t := range things {
		thingsMap[t] = true
	}

	things = []string{}

	for t := range thingsMap {
		things = append(things, t)
	}

	return things
}

// ResolveThingChLinks resolve the thing-channel edge in Dgraph
func (rule *Rule) ResolveThingChLinks(fn func(id string) (string, error)) error {
	if rule.Thing != "" && rule.Channel != "" {
		uid, err := fn(rule.Thing + "-" + rule.Channel)
		if err != nil {
			return err
		}
		rule.ThingChLink = &ThingChLink{
			Uid: uid,
		}
	}

	for _, subRule := range rule.Rules {
		err := subRule.ResolveThingChLinks(fn)
		if err != nil {
			return err
		}
	}

	return nil
}

// AssignRootID assigns root id to entire rule tree
func (rule *Rule) AssignRootID(id string) {
	rule.Root = id
	if rule.Rules != nil {
		for _, r := range rule.Rules {
			r.AssignRootID(id)
		}
	}
}

// GetWithDType adds DType
func (rule *Rule) GetWithDType() *RuleWithDType {
	newRule := &RuleWithDType{
		Rule:  rule,
		DType: []string{"Rule"},
		Rules: []*RuleWithDType{},
	}

	for _, r := range rule.Rules {
		newRule.Rules = append(newRule.Rules, r.GetWithDType())
	}

	return newRule
}

// GetWithDType adds DType
func (t *ThingChLink) GetWithDType() *ThingChLinkWithDType {
	newT := &ThingChLinkWithDType{
		ThingChLink: t,
		DType:       []string{"ThingChLink"},
	}

	return newT
}

// GetWithDType gets DType
func (t *TriggerCommand) GetWithDType() *TriggerCommandWithDType {
	newT := &TriggerCommandWithDType{
		TriggerCommand: t,
		DType:          []string{"TriggerCommand"},
		Channels:       []*ChannelWithDType{},
	}

	if t.Channels != nil {
		for _, ch := range t.Channels {
			newT.Channels = append(newT.Channels, ch.GetWithDType())
		}
	}

	return newT
}

// GetWithDType gets DType
func (c *Channel) GetWithDType() *ChannelWithDType {
	return &ChannelWithDType{
		Channel: c,
		DType:   []string{"Channel"},
	}
}
