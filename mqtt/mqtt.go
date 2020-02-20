package main

import (
	"encoding/json"
	"reflect"
	"strings"
	"thingz-server/lib"
	proto "thingz-server/mqtt/proto"
	topics "thingz-server/mqtt/topics"
	thingProto "thingz-server/thing/proto"
	thingTopics "thingz-server/thing/topics"

	log "github.com/sirupsen/logrus"
)

type channel struct {
	ID    string      `json:"id,omitempty"`
	Type  int         `json:"type,omitempty"`
	Value interface{} `json:"value,omitempty"`
}

type mqttPacket struct {
	Thing    string    `json:"thing"`
	Channels []channel `json:"channels"`
}

func (a *app) updateChannel(msg *lib.Msg) {
	log.Println("updateChanel")
	c := &mqttPacket{}
	/* h := new(codec.MsgpackHandle)
	dec := codec.NewDecoderBytes(msg.Msg, h)
	err := dec.Decode(c) */
	// err := msgpack.NewDecoder(bytes.NewReader(msg.Msg)).UseJSONTag(true).Decode(c)
	err := json.Unmarshal(msg.Msg, c)
	if err != nil {
		log.Printf("mqtt update channel error: %+v", err)
	}

	log.Println("c", c)

	req := &thingProto.UpdateChannelsRequest{
		Thing: c.Thing,
		// PhysicalId: c.PhysicalID,
		Channels: []*thingProto.Channel{},
	}

	for _, channel := range c.Channels {
		pc := &thingProto.Channel{
			Id:   channel.ID,
			Unit: thingProto.Unit(channel.Type),
		}

		switch pc.Unit {
		case thingProto.Unit_BOOL:
			pc.BoolValue = channel.Value.(bool)
		case thingProto.Unit_NUMBER:
			if v, ok := channel.Value.(float32); ok {
				pc.FloatValue = float64(v)
			} else if v, ok := channel.Value.(float64); ok {
				pc.FloatValue = v
			} else if v, ok := channel.Value.(int); ok {
				pc.FloatValue = float64(v)
			} else if v, ok := channel.Value.(int8); ok {
				pc.FloatValue = float64(v)
			} else {
				log.Println(reflect.TypeOf(channel.Value))
				return
			}
		case thingProto.Unit_DATA:
			pc.DataValue = channel.Value.(string)
		case thingProto.Unit_STRING:
			pc.StringValue = channel.Value.(string)
		}

		req.Channels = append(req.Channels, pc)
	}

	log.Println(req)

	a.eb.SendMessage(thingTopics.UpdateChannels, req)
}

func (a *app) updateThing(_, reply string, req *proto.UpdateThingRequest) {
	log.Printf("updateThing req: %+v", req)
	resp := &proto.UpdateThingResponse{}

	msg := &mqttPacket{
		Thing:    req.Thing.Id,
		Channels: []channel{},
	}

	for _, c := range req.Thing.Channels {
		var val interface{}

		switch thingProto.Unit(c.Unit) {
		case thingProto.Unit_BOOL:
			val = c.BoolValue
		case thingProto.Unit_NUMBER:
			val = c.FloatValue
		case thingProto.Unit_DATA:
			val = c.DataValue
		case thingProto.Unit_STRING:
			val = c.StringValue
		}

		msg.Channels = append(msg.Channels, channel{
			ID:    c.Id,
			Type:  int(c.Unit),
			Value: val,
		})
	}

	// var b bytes.Buffer
	// buffer := bufio.NewWriter(&b)

	// err := msgpack.NewEncoder(buffer).UseJSONTag(true).Encode(msg)

	/* var b = make([]byte, 0, 1024)
	var h codec.Handle = new(codec.MsgpackHandle)
	var enc *codec.Encoder = codec.NewEncoderBytes(&b, h)
	var err = enc.Encode(msg) //any of v1 ... v8 */

	b, err := json.Marshal(msg)
	if err == nil {
		err = a.pb.SendMessage(topics.UpdateDevice+"/"+strings.Trim(req.Thing.PhysicalId, " "), b)
	}

	if err != nil {
		resp.Success = false
		resp.Error = err.Error()
	} else {
		resp.Success = true
		log.Println(b, msg, topics.UpdateDevice+"/"+req.Thing.PhysicalId)
	}

	log.Printf("Request: %+v, Resposne: %+v", req, resp)
	if reply != "" {
		a.eb.SendMessage(reply, resp)
	}
}

func (a *app) createMqtt(_, reply string, req *proto.CreateMqttRequest) {
	log.Printf("create req: %+v", req)
	resp := &proto.CreateMqttResponse{}
	db, err := a.db.GetMongoSession()
	if err == nil {
		defer db.Close()
		req.Mqtt.Id = lib.NewObjectID()
		err = db.DB("").C(collectionName).Insert(req.Mqtt)
	}
	if err != nil {
		resp.Success = false
		resp.Error = err.Error()
	} else {
		resp.Success = true
		resp.Id = req.Mqtt.Id
	}

	log.Printf("Request: %+v, Resposne: %+v", req, resp)
	if reply != "" {
		a.eb.SendMessage(reply, resp)
	}
}

func (a *app) userMqtts(_, reply string, req *proto.UserMqttsRequest) {
	log.Printf("create req: %+v", req)
	resp := &proto.UserMqttsResponse{}
	mqtts := []*proto.Mqtt{}
	db, err := a.db.GetMongoSession()
	if err == nil {
		defer db.Close()
		db.DB("").C(collectionName).Find(lib.M{"owner": req.GetOwner()}).All(&mqtts)
	}
	if err != nil {
		resp.Success = false
		resp.Error = err.Error()
	} else {
		resp.Success = true
		resp.Mqtts = mqtts
	}

	log.Printf("Request: %+v, Resposne: %+v", req, resp)
	if reply != "" {
		a.eb.SendMessage(reply, resp)
	}
}
