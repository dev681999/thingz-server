package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
	googlesmarthome "thingz-server/api/google-smart-home"
	"thingz-server/lib"
	sceneP "thingz-server/scene/proto"
	sceneT "thingz-server/scene/topics"
	thingP "thingz-server/thing/proto"
	thingT "thingz-server/thing/topics"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
)

func (a *app) googleSmartHomeHandler(c echo.Context) error {
	// log.Printf("Req, %+v", string(c.Request().Body))
	gsr := &googlesmarthome.Request{}
	// err := c.Bind(req)

	body, err := ioutil.ReadAll(c.Request().Body)
	err = json.Unmarshal(body, gsr)

	user := c.Request().Header.Get("Authorization")
	jwtToken := strings.Split(user, " ")
	log.Println("jwt", a.c.JwtSecret)

	token, err := jwt.ParseWithClaims(jwtToken[1], &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		// since we only use the one private key to sign the tokens,
		// we also only use its public counter part to verify
		return []byte(a.c.JwtSecret), nil
	})
	claims := *(token.Claims.(*jwt.MapClaims))
	// fatal(err)

	// claims := token.(jwt.MapClaims)

	if err != nil || len(gsr.Inputs) <= 0 {
		log.Printf("Bind error: %+v", err)
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "incorrect request",
		})
	}
	/* b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Printf("Body: %s", string(b))
	} */
	/* hub = claims["hId"].(string)
	log.Printf("GSH Req: %+v %+v, Hub: %s", req, claims, hub)
	*/
	project := "5e46773fe64d1e0ff31b7493"
	if gsr.Inputs[0].Intent == googlesmarthome.IntentSync {
		req := &thingP.ProjectThingsRequest{
			Project: project,
		}
		res := &thingP.ProjectThingsResponse{}

		err := a.eb.RequestMessage(thingT.ProjectThings, req, res, lib.DefaultTimeout)
		if err != nil {
			log.Printf("error: %+v", err)
			return a.makeError(c, http.StatusInternalServerError, err)
		}

		if !res.GetSuccess() {
			return a.makeError(c, http.StatusUnauthorized, errors.New(res.GetError()))
		}

		if res.Things == nil {
			res.Things = []*thingP.Thing{}
		}

		sreq := &sceneP.UserScenesRequest{
			Owner:   claims["id"].(string),
			Project: project,
		}
		sres := &sceneP.UserScenesResponse{}

		err = a.eb.RequestMessage(sceneT.UserScenes, sreq, sres, lib.DefaultTimeout)
		if err != nil {
			log.Printf("error: %+v", err)
			return a.makeError(c, http.StatusInternalServerError, err)
		}

		if !sres.GetSuccess() {
			return a.makeError(c, http.StatusUnauthorized, errors.New(res.GetError()))
		}

		if sres.Scenes == nil {
			sres.Scenes = []*sceneP.Scene{}
		}

		resDevices := []googlesmarthome.SyncDevice{}
		devices := res.Things
		// scenes := res.Home.Scenes
		for _, device := range devices {
			traits := []string{}
			for _, ch := range device.Channels {
				traits = append(traits, ch.Id)
			}
			resDevices = append(resDevices,
				googlesmarthome.SyncDevice{
					ID: device.Id,
					Name: googlesmarthome.Name{
						Name: device.Name,
					},
					RoomHint:        "",
					Traits:          googlesmarthome.ConvertCommandToTraits(traits),
					Type:            googlesmarthome.DeviceTypeMap[int(device.Type)],
					WillReportState: false,
				},
			)
		}

		for _, scene := range sres.Scenes {
			resDevices = append(resDevices,
				googlesmarthome.SyncDevice{
					ID: scene.GetId(),
					Name: googlesmarthome.Name{
						Name: scene.Name,
					},
					RoomHint:        "",
					Traits:          []string{googlesmarthome.TraitScene},
					Type:            googlesmarthome.TypeScene,
					WillReportState: false,
				},
			)
		}

		syncRes := googlesmarthome.SyncResponse{
			RequestID: gsr.RequestID,
			Payload: googlesmarthome.SyncPayload{
				AgentUserID: claims["id"].(string),
				Devices:     resDevices,
			},
		}

		b, _ := json.MarshalIndent(syncRes, "", "	")

		log.Printf("Sync Res: %+v", string(b))

		return c.JSON(http.StatusOK, syncRes)
	}

	if gsr.Inputs[0].Intent == googlesmarthome.IntentQuery {
		queryReq := &googlesmarthome.Query{}
		log.Println("body", string(body))
		err = json.Unmarshal(body, queryReq)

		if err != nil {
			log.Printf("Query error: %+v", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"message": err.Error(),
			})
		}

		payload := queryReq.Inputs[0].Payload
		queryRes := &googlesmarthome.QueryResponse{
			RequestID: queryReq.RequestID,
		}

		ids := []string{}

		for _, device := range payload.Devices {
			ids = append(ids, device.ID)
		}

		log.Println("payload devices", ids)

		req := &thingP.ProjectThingsRequest{
			Project: project,
		}
		res := &thingP.ProjectThingsResponse{}

		err := a.eb.RequestMessage(thingT.ProjectThings, req, res, lib.DefaultTimeout)
		if err != nil {
			log.Printf("error: %+v", err)
			return a.makeError(c, http.StatusInternalServerError, err)
		}

		if !res.GetSuccess() {
			return a.makeError(c, http.StatusUnauthorized, errors.New(res.GetError()))
		}

		if res.Things == nil {
			res.Things = []*thingP.Thing{}
		}

		thingsMap := map[string]*thingP.Thing{}
		for _, t := range res.Things {
			thingsMap[t.Id] = t
		}

		queryRes.Payload = googlesmarthome.QueryResponsePayload{
			Devices: make(map[string]map[string]interface{}, len(ids)),
		}

		for _, d := range ids {
			t := thingsMap[d]

			log.Println("t", *t)

			queryRes.Payload.Devices[t.Id] = make(map[string]interface{})
			for _, ch := range t.Channels {
				st := ch.Id
				if st == "onOff" {
					st = "on"
				}
				var val interface{}

				switch ch.Unit {
				case thingP.Unit_BOOL:
					val = ch.BoolValue
				case thingP.Unit_NUMBER:
					val = ch.FloatValue
				case thingP.Unit_DATA:
					val = ch.DataValue
				}

				queryRes.Payload.Devices[t.Id][st] = val
			}
			queryRes.Payload.Devices[t.Id]["online"] = true
		}

		log.Printf("Query Res: %+v", queryRes)

		return c.JSON(http.StatusOK, queryRes)
	}

	if gsr.Inputs[0].Intent == googlesmarthome.IntentExecute {
		req := &thingP.ProjectThingsRequest{
			Project: project,
		}
		res := &thingP.ProjectThingsResponse{}

		err := a.eb.RequestMessage(thingT.ProjectThings, req, res, lib.DefaultTimeout)
		if err != nil {
			log.Printf("error: %+v", err)
			return a.makeError(c, http.StatusInternalServerError, err)
		}

		if !res.GetSuccess() {
			return a.makeError(c, http.StatusUnauthorized, errors.New(res.GetError()))
		}

		if res.Things == nil {
			res.Things = []*thingP.Thing{}
		}

		execReq := &googlesmarthome.Execute{}
		err = json.Unmarshal(body, execReq)
		if err != nil {
			log.Printf("Execute error: %+v", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"message": err.Error(),
			})
		}

		log.Println("got", string(body))

		updateRes := &thingP.UpdateThingsChannelsRequest{
			Things: []*thingP.ThingChannels{},
		}

		thingsMap := map[string]*thingP.Thing{}
		for _, t := range res.Things {
			thingsMap[t.Id] = t
		}

		scenesToStart := &sceneP.ActivateScenesRequest{
			Ids:   []string{},
			Owner: claims["id"].(string),
		}

		for _, cmd := range execReq.Inputs[0].Payload.Commands {
			deviceCmds := []*thingP.Channel{}
			skip := false
			for _, exec := range cmd.Execution {
				var val = exec.Params[googlesmarthome.CommandConvertParam[exec.Command]]

				log.Println("val", val, "::type-", reflect.TypeOf(val))

				if val == nil {
					skip = true

					for _, dev := range cmd.Devices {
						scenesToStart.Ids = append(scenesToStart.Ids, dev.ID)
					}
					/* return a.sendSucess(c, googlesmarthome.ExecuteResponse{
						RequestID: execReq.RequestID,
						Payload: googlesmarthome.ExecuteResponsePayload{
							Commands: []googlesmarthome.ExecuteResponseCommand{{
								IDS:    []string{cmd.Devices[0].ID},
								Status: "SUCCESS",
								States: make(map[string]interface{}),
							}},
						},
					}) */

					break
				}

				ch := &thingP.Channel{}
				ch.Id = googlesmarthome.CommandMap[exec.Command]
				ch.Unit = thingP.Unit(int32(googlesmarthome.UnitMap[exec.Command]))

				switch ch.Unit {
				case thingP.Unit_BOOL:
					ch.BoolValue = val.(bool)
				case thingP.Unit_NUMBER:
					ch.FloatValue = val.(float64)
				case thingP.Unit_DATA:
					ch.DataValue = val.(string)
				}

				deviceCmds = append(deviceCmds, ch)
			}

			if skip {
				continue
			}

			for _, d := range cmd.Devices {
				updateRes.Things = append(updateRes.Things, &thingP.ThingChannels{
					Id:         d.ID,
					PhysicalId: thingsMap[d.ID].PhysicalId,
					Channels:   deviceCmds,
				})
			}
		}

		successIDs := []string{}
		failIDs := []string{}

		if len(updateRes.Things) > 0 {
			updateResp := &thingP.UpdateThingsChannelsResponse{}

			err = a.eb.RequestMessage(thingT.UpdateThingsChannels, updateRes, updateResp, lib.DefaultTimeout)
			if err == nil && updateResp.GetError() != "" {
				err = errors.New(updateResp.GetError())
			}
			if err != nil {
				log.Printf("Execute error: %+v", err)
				return c.JSON(http.StatusInternalServerError, map[string]string{
					"message": err.Error(),
				})
			}

			for _, t := range updateRes.Things {
				successIDs = append(successIDs, t.Id)
			}
		}

		if len(scenesToStart.Ids) > 0 {
			res := &sceneP.ActivateScenesResponse{}

			log.Printf("sending request %+v", *req)

			err := a.eb.RequestMessage(sceneT.ActivateScenes, scenesToStart, res, lib.DefaultTimeout)
			if err != nil {
				log.Printf("error: %+v", err)
				return a.makeError(c, http.StatusInternalServerError, err)
			}

			if !res.GetSuccess() {
				return a.makeError(c, http.StatusUnauthorized, errors.New(res.GetError()))
			}

			successIDs = append(successIDs, scenesToStart.GetIds()...)
		}

		execCmdRes := []googlesmarthome.ExecuteResponseCommand{}

		if len(successIDs) > 0 {
			execCmdRes = append(execCmdRes, googlesmarthome.ExecuteResponseCommand{
				IDS:    successIDs,
				Status: "SUCCESS",
			})
		}

		if len(failIDs) > 0 {
			execCmdRes = append(execCmdRes, googlesmarthome.ExecuteResponseCommand{
				IDS:       failIDs,
				Status:    "ERROR",
				ErrorCode: "deviceTurnedOff",
			})
		}

		return c.JSON(http.StatusOK, googlesmarthome.ExecuteResponse{
			RequestID: execReq.RequestID,
			Payload: googlesmarthome.ExecuteResponsePayload{
				Commands: execCmdRes,
			},
		})
	}

	return c.NoContent(http.StatusOK)
}
