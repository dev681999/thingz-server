package main

import (
	"log"
	"thingz-server/rule/topics"

	lib "github.com/dev681999/helperlibs"

	nats "github.com/nats-io/nats.go"
)

var collectionName = "rules"

type appConfig struct {
	DBURL   string `json:"dbUrl"`
	DBUser  string `json:"dbUser"`
	DBPass  string `json:"dbPass"`
	DBName  string `json:"dbName"`
	NATSUrl string `json:"natsUrl"`
}

type app struct {
	eb *lib.EventBus
	db *lib.Store
	c  *appConfig
	h  *lib.Hash
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
	log.Println("connecting nats")

	err = a.eb.Connect(lib.ProtobufEnc)
	if err != nil {
		a.Close()
		return err
	}

	log.Println("connecting nats sucess")
	log.Println("registering to event-bus")

	// log.Println("rule topic", topics.CheckThingRule)

	listeners := []lib.Listener{
		lib.Listener{
			Topic: topics.CreateRule,
			Func:  a.createRule,
		},
		lib.Listener{
			Topic: topics.DeleteRule,
			Func:  a.deleteRule,
		},
		lib.Listener{
			Topic: topics.CheckThingRule,
			Func:  a.checkThingRule,
		},
		lib.Listener{
			Topic: topics.ProjectRules,
			Func:  a.projectRules,
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
	log.Println("close complete")
}

func (a *app) Test() string {
	return "user-srv"
}
