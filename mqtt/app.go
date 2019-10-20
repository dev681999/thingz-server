package main

import (
	"log"
	"thingz-server/lib"
	"thingz-server/mqtt/topics"

	nats "github.com/nats-io/nats.go"
)

var collectionName = "mqtts"

type appConfig struct {
	DBURL   string `json:"dbUrl"`
	DBUser  string `json:"dbUser"`
	DBPass  string `json:"dbPass"`
	DBName  string `json:"dbName"`
	NATSUrl string `json:"natsUrl"`
	MQTTUrl string `json:"mqttUrl"`
}

type app struct {
	eb *lib.EventBus
	db *lib.Store
	c  *appConfig
	h  *lib.Hash
	pb *lib.MqttProtocolBus
}

func newApp(config *appConfig) *app {
	if config == nil {
		return nil
	}

	s := &lib.Store{
		Address:  config.DBURL,
		Database: config.DBName,
		Password: config.DBPass,
		Username: config.DBUser,
	}

	if config.NATSUrl == "" {
		config.NATSUrl = nats.DefaultURL
	}

	return &app{
		c:  config,
		db: s,
		eb: lib.NewEventBusUnsecure(config.NATSUrl),
		h:  &lib.Hash{},
		pb: lib.NewMqttProtocolBus(config.MQTTUrl, topics.Base, func(topic string, payload []byte) {
			log.Println(topic, payload)
		}),
	}
}

func (a *app) Init() error {
	log.Println("init start")
	log.Println("connecting db")

	err := a.db.Connect()
	if err != nil {
		a.Close()
		return err
	}

	log.Println("connecting db success")
	log.Println("connecting mqtt")

	err = a.pb.Connect()
	if err != nil {
		a.Close()
		return err
	}

	log.Println("connecting mqtt sucess")
	log.Println("connecting nats")

	err = a.eb.Connect(lib.ProtobufEnc)
	if err != nil {
		a.Close()
		return err
	}

	log.Println("connecting nats sucess")
	log.Println("registering to mqtt-bus", topics.UpdateThing)

	err = a.pb.RegisterListener(topics.UpdateThing, a.updateChannel)
	if err != nil {
		a.Close()
		return err
	}

	log.Println("registering to mqtt-bus complete")
	log.Println("registering to event-bus")

	listeners := []lib.Listener{
		lib.Listener{
			Topic: topics.CreateMqtt,
			Func:  a.createMqtt,
		},
		lib.Listener{
			Topic: topics.UserMqtts,
			Func:  a.userMqtts,
		},
		lib.Listener{
			Topic: topics.UpdateThing,
			Func:  a.updateThing,
		},
	}

	err = a.eb.RegisterListeners(listeners)
	if err != nil {
		a.Close()
		return err
	}

	log.Println("registering to event-bus complete")
	log.Println("init complete")

	return nil
}

func (a *app) Close() {
	log.Println("closing start")
	log.Println("closing db connection")

	if a.db != nil {
		a.db.Close()
	}

	log.Println("closed db connection")
	log.Println("closing nats connection")

	if a.eb != nil {
		a.eb.Close()
	}

	log.Println("closed nats connection")
	log.Println("closing nats connection")

	if a.eb != nil {
		a.pb.Close()
	}

	log.Println("closed nats connection")
	log.Println("close complete")
}

func (a *app) Test() string {
	return "user-srv"
}
