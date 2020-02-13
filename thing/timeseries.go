package main

import (
	apiP "thingz-server/api/proto"
	apiT "thingz-server/api/topics"
	proto "thingz-server/thing/proto"
	"time"

	"github.com/globalsign/mgo"
)

var timeSeriesCollectionName = "timeseries"

func (a *app) insertSeries(thing string, channels []*proto.Channel, db *mgo.Session) error {
	updateChans := []*apiP.Channel{}

	for _, ch := range channels {
		updateChans = append(updateChans, &apiP.Channel{
			Id:          ch.Id,
			Thing:       ch.Thing,
			Name:        ch.Name,
			FloatValue:  ch.FloatValue,
			StringValue: ch.StringValue,
			BoolValue:   ch.BoolValue,
			DataValue:   ch.DataValue,
			Unit:        apiP.Unit(ch.Unit),
			Type:        ch.Type,
			IsSensor:    ch.IsSensor,
		})
	}

	a.eb.SendMessage(apiT.SendThingUpdate, &apiP.SendThingUpdateRequest{
		Update: &apiP.ThingUpdate{
			Thing:    thing,
			Channels: updateChans,
		},
	})
	var serieses []interface{}
	for _, c := range channels {
		s := proto.Series{
			Thing:       thing,
			Channel:     c.Id,
			FloatValue:  c.FloatValue,
			StringValue: c.StringValue,
			BoolValue:   c.BoolValue,
			DataValue:   c.DataValue,
			Unit:        c.Unit,
			TimeStamp:   time.Now().Format(time.RFC3339),
		}

		serieses = append(serieses, s)
	}
	return db.DB("").C(timeSeriesCollectionName).Insert(serieses...)
}
