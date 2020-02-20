package main

import (
	"log"
	"thingz-server/lib"
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

func main() {
	pb := lib.NewMqttProtocolBus("tcp://localhost:1883", "server.Mqtt.", func(topic string, payload []byte) {
		log.Println(topic, payload)
	})

	err := pb.Connect()
	if err != nil {
		panic(err)
	}

	defer pb.Close()
	err = pb.RegisterListener("test", func(msg *lib.Msg) {
		log.Println(string(msg.Msg))
	})

	select {}
}
