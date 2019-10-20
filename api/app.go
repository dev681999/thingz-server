package main

import (
	"context"
	"log"
	"thingz-server/api/topics"
	"thingz-server/lib"
	"time"

	nats "github.com/nats-io/nats.go"

	"github.com/alexandrevicenzi/go-sse"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type appConfig struct {
	Addr      string `json:"addr"`
	DBURL     string `json:"dbUrl"`
	DBUser    string `json:"dbUser"`
	DBPass    string `json:"dbPass"`
	DBName    string `json:"dbName"`
	JwtSecret string `json:"jwt"`
	NATSUrl   string `json:"natsUrl"`
}

type app struct {
	eb          *lib.EventBus
	c           *appConfig
	e           *echo.Echo
	eventServer *sse.Server
}

func newApp(config *appConfig) *app {
	if config == nil {
		return nil
	}

	if config.NATSUrl == "" {
		config.NATSUrl = nats.DefaultURL
	}

	eventServer := sse.NewServer(nil)

	return &app{
		c:           config,
		eb:          lib.NewEventBusUnsecure(config.NATSUrl),
		e:           echo.New(),
		eventServer: eventServer,
	}
}

func (a *app) Init() error {
	log.Println("init start")
	log.Println("connecting nats")

	err := a.eb.Connect(lib.ProtobufEnc)
	if err != nil {
		a.Close()
		return err
	}

	log.Println("connecting nats sucess")
	log.Println("registering to event-bus")

	// log.Println("rule topic", topics.CheckThingRule)

	listeners := []lib.Listener{
		lib.Listener{
			Topic: topics.SendThingUpdate,
			Func:  a.sendUpdateThing,
		},
	}

	err = a.eb.RegisterListeners(listeners)
	if err != nil {
		a.Close()
		return err
	}

	log.Println("registering to event-bus complete")
	log.Println("starting HTTP server")

	a.e.HideBanner = true
	a.e.Use(middleware.Logger())
	a.e.Use(middleware.Recover())
	a.e.Use(middleware.CORS())

	a.e.POST("/login", a.login)
	a.e.POST("/register", a.register)

	a.e.POST("/assign-thing", a.assignThing)
	a.e.POST("/assign-thing/", a.assignThing)

	api := a.e.Group("/api")

	api.Use(middleware.JWT([]byte(a.c.JwtSecret)))

	project := api.Group("/project")
	thing := api.Group("/thing")
	rule := api.Group("/rule")

	project.POST("", a.createProject)
	project.POST("/", a.createProject)

	project.GET("", a.userProjects)
	project.GET("/", a.userProjects)

	project.DELETE("/:id", a.deleteProject)
	project.DELETE("/:id/", a.deleteProject)

	project.GET("/:id/things", a.projectThings)
	project.GET("/:id/things/", a.projectThings)

	project.GET("/:id/rules", a.projectRules)
	project.GET("/:id/rules/", a.projectRules)

	thing.POST("", a.createThing)
	thing.POST("/", a.createThing)

	thing.GET("/:id", a.getThing)
	thing.GET("/:id/", a.getThing)

	thing.GET("/:id/series", a.getThingSeries)
	thing.GET("/:id/series/", a.getThingSeries)

	thing.POST("/:id", a.generateAssignThing)
	thing.POST("/:id/", a.generateAssignThing)

	// thing.DELETE("/:id", a.deleteThing)
	// thing.DELETE("/:id/", a.deleteThing)

	thing.DELETE("/:id", a.deassignThing)
	thing.DELETE("/:id/", a.deassignThing)

	thing.PATCH("/:id/channel", a.updateChannel)
	thing.PATCH("/:id/channel/", a.updateChannel)

	thing.Any("/events", a.handleUpdateThingEvent)

	rule.POST("", a.createRule)
	rule.POST("/", a.createRule)

	rule.DELETE("/:id", a.deleteRule)
	rule.DELETE("/:id/", a.deleteRule)

	go func() {
		if err := a.e.Start(a.c.Addr); err != nil {
			log.Println("HTTP server shutdown")
		}
	}()

	log.Println("init complete")
	return nil
}

func (a *app) Close() {
	log.Println("closing start")
	log.Println("closing HTTP server")

	if a.e != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		if err := a.e.Shutdown(ctx); err != nil {
			log.Println(err)
		}
		cancel()
	}

	log.Println("closed echo server")
	log.Println("closing nats connection")

	if a.eb != nil {
		a.eb.Close()
	}

	log.Println("closed nats connection")
	log.Println("close complete")
}
