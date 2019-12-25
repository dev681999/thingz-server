package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
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

// go-run --thing 8TZVXBQ0 --channel value --valType 2 --stringVal 'A'

func main() {
	thingID := flag.String("thing", "", "Thing ID")
	channelID := flag.String("channel", "", "Channel ID")
	floatVal := flag.Float64("floatValue", 0, "Float Value")
	boolVal := flag.Bool("boolValue", false, "Bool Value")
	stringVal := flag.String("stringVal", "", "String Value")
	valType := flag.Int("valType", 0, "Value Type")

	flag.Parse()

	fmt.Println(*thingID, *channelID, *floatVal, *boolVal, *valType, *stringVal)

	opts := mqtt.NewClientOptions().AddBroker("tcp://ec2-54-191-142-124.us-west-2.compute.amazonaws.com:1883").SetClientID("thingz-mqtt-test")
	// opts := mqtt.NewClientOptions().AddBroker("tcp://localhost:1883").SetClientID("thingz-mqtt-test")
	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	p := mqttPacket{
		Thing: *thingID,
		Channels: []channel{{
			ID:   *channelID,
			Type: *valType,
		}},
	}
	switch *valType {
	case 0:
		p.Channels[0].Value = boolVal
	case 1:
		p.Channels[0].Value = floatVal
	case 2:
		p.Channels[0].Value = stringVal
	}
	b, _ := json.Marshal(p)

	token := c.Publish("server.Mqtt.UpdateThing", 0, false, string(b))
	token.Wait()

	c.Disconnect(250)

	time.Sleep(1 * time.Second)
}
