package main

import (
	"thingz-server/thing/topics"

	log "github.com/sirupsen/logrus"

	lib "github.com/dev681999/helperlibs"

	"github.com/go-redis/redis"
	nats "github.com/nats-io/nats.go"
)

const collectionName = "things"

type appConfig struct {
	DBURL    string `json:"dbUrl"`
	DBUser   string `json:"dbUser"`
	DBPass   string `json:"dbPass"`
	DBName   string `json:"dbName"`
	NATSUrl  string `json:"natsUrl"`
	CacheURL string `json:"cacheUrl"`
}

type app struct {
	eb          *lib.EventBus
	db          *lib.Store
	c           *appConfig
	h           *lib.Hash
	cacheClient *redis.Client
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
	log.Println("connecting cache")

	cacheClient := redis.NewClient(&redis.Options{
		Addr:     a.c.CacheURL,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err = cacheClient.Ping().Result()
	if err != nil {
		a.Close()
		return err
	}

	a.cacheClient = cacheClient

	log.Println("connecting db success")
	log.Println("connecting nats")

	err = a.eb.Connect(lib.ProtobufEnc)
	if err != nil {
		a.Close()
		return err
	}

	log.Println("connecting nats sucess")
	log.Println("registering to event-bus")

	listeners := []lib.Listener{
		lib.Listener{
			Topic: topics.CreateThing,
			Func:  a.createThing,
		}, {
			Topic: topics.CreateThings,
			Func:  a.createThings,
		}, lib.Listener{
			Topic: topics.GetThing,
			Func:  a.getThing,
		},
		lib.Listener{
			Topic: topics.ProjectThings,
			Func:  a.projectThings,
		},
		lib.Listener{
			Topic: topics.ProjectDelete,
			Func:  a.projectDelete,
		},
		lib.Listener{
			Topic: topics.DeleteThing,
			Func:  a.deleteThing,
		},
		/* lib.Listener{
			Topic: topics.UpdateChannel,
			Func:  a.updateChannel,
		}, */
		lib.Listener{
			Topic: topics.UpdateChannels,
			Func:  a.updateChannels,
		},
		lib.Listener{
			Topic: topics.UpdateThingsChannels,
			Func:  a.updateThingsChannels,
		},
		lib.Listener{
			Topic: topics.GenerateAssignThing,
			Func:  a.generateAssignThing,
		},
		lib.Listener{
			Topic: topics.AssignThing,
			Func:  a.assignThing,
		},
		lib.Listener{
			Topic: topics.DeassignThing,
			Func:  a.deassignThing,
		},
		lib.Listener{
			Topic: topics.ThingSeries,
			Func:  a.thingSeries,
		},
		lib.Listener{
			Topic: topics.GetThingsByIDs,
			Func:  a.getThingsByIDs,
		},
		lib.Listener{
			Topic: topics.UpdateThingConfig,
			Func:  a.updateThingConfig,
		},
		lib.Listener{
			Topic: topics.GetThingTypes,
			Func:  a.getThingTypes,
		}, {
			Topic: topics.UpdateThing,
			Func:  a.updateThing,
		}, {
			Topic: topics.UpdateChannelName,
			Func:  a.updateChannelName,
		},
	}

	err = a.eb.RegisterListeners(listeners)
	if err != nil {
		a.Close()
		return err
	}

	log.Println("registering to event-bus complete")
	log.Println("init complete")

	/* go func() {
		a.updateChannels("", "", &proto.UpdateChannelsRequest{
			Thing: "2",
			Channels: []*proto.Channel{{
				Id:         "sensor",
				FloatValue: 10.0,
				Unit:       1,
			}},
		})
	}() */

	/* go func(a *app) {
		resp := &mqttProto.UpdateThingResponse{}
		err := a.eb.RequestMessage(mqttTopics.UpdateThing, &mqttProto.UpdateThingRequest{
			Thing: &mqttProto.Thing{
				Id: "test",
				Channels: []*mqttProto.Channel{{
					Id:          "1",
					Unit:        int32(proto.Unit_STRING),
					StringValue: "TEST",
				}},
			},
		}, resp, lib.DefaultTimeout)
		if err != nil {
			log.Println(err)
			return
		}

		log.Println(resp)
	}(a) */

	return nil
}

func (a *app) Close() {
	log.Println("closing start")
	log.Println("closing db connection")

	if a.db != nil {
		a.db.Close()
	}

	log.Println("closed db connection")
	log.Println("closing db connection")

	if a.cacheClient != nil {
		a.cacheClient.Close()
	}

	log.Println("closed db connection")
	log.Println("closing nats connection")

	if a.eb != nil {
		a.eb.Close()
	}

	log.Println("closed nats connection")
	log.Println("close complete")
}
