package main

import (
	"context"
	"log"
	"thingz-server/lib"
	"time"

	nats "github.com/nats-io/nats.go"

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
	eb *lib.EventBus
	c  *appConfig
	e  *echo.Echo
}

func newApp(config *appConfig) *app {
	if config == nil {
		return nil
	}

	if config.NATSUrl == "" {
		config.NATSUrl = nats.DefaultURL
	}

	return &app{
		c:  config,
		eb: lib.NewEventBusUnsecure(config.NATSUrl),
		e:  echo.New(),
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
	log.Println("starting HTTP server")

	a.e.HideBanner = true
	a.e.Use(middleware.Logger())
	a.e.Use(middleware.Recover())
	a.e.Use(middleware.CORS())

	a.e.POST("/login", a.login)
	a.e.POST("/register", a.register)

	api := a.e.Group("/api")

	api.Use(middleware.JWT([]byte(a.c.JwtSecret)))

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

func (a *app) Test() string {
	return "api-srv"
}
