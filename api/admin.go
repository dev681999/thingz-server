package main

import (
	"errors"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"thingz-server/lib"
	ruleP "thingz-server/rule/proto"
	ruleT "thingz-server/rule/topics"
	thingP "thingz-server/thing/proto"
	thingT "thingz-server/thing/topics"

	"github.com/jung-kurt/gofpdf"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
	"github.com/skip2/go-qrcode"
)

func (a *app) adminGetAllThings(c echo.Context) error {
	req := &thingP.GetThingsByIDsRequest{}
	res := &thingP.GetThingsByIDsResponse{}

	err := a.eb.RequestMessage(thingT.GetThingsByIDs, req, res, lib.DefaultTimeout)
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

	return a.sendSucess(c, echo.Map{
		"msg":    "ok",
		"things": res.GetThings(),
	})
}

func (a *app) adminCreateThings(c echo.Context) error {
	things := []int32{}
	err := c.Bind(&things)
	if err != nil {
		log.Printf("error: %+v", err)
		return a.makeError(c, http.StatusInternalServerError, err)
	}

	log.Printf("creating devices: %+v", things)

	req := &thingP.CreateThingsRequest{
		Things: []*thingP.Thing{},
	}

	for _, t := range things {
		req.Things = append(req.Things, &thingP.Thing{
			Type: t,
		})
	}

	res := &thingP.CreateThingsResponse{}
	err = a.eb.RequestMessage(thingT.CreateThings, req, res, lib.DefaultTimeout)
	if err != nil {
		log.Printf("error: %+v", err)
		return a.makeError(c, http.StatusInternalServerError, err)
	}
	if !res.GetSuccess() {
		log.Printf("error: %+v", res.GetError())
		return a.makeError(c, http.StatusUnauthorized, errors.New(res.GetError()))
	}

	for _, t := range res.Things {
		trReq := &ruleP.CreateThingLinkRequest{
			Thing:    t.GetId(),
			Channels: t.Channels,
		}
		trRes := &ruleP.CreateThingLinkResponse{}

		err = a.eb.RequestMessage(ruleT.CreateThingLink, trReq, trRes, lib.DefaultTimeout)
		if err != nil {
			log.Printf("error: %+v", err)
			return a.makeError(c, http.StatusInternalServerError, err)
		}
		if !res.GetSuccess() {
			log.Printf("error: %+v", res.GetError())
			return a.makeError(c, http.StatusUnauthorized, errors.New(res.GetError()))
		}
	}

	return a.sendSucess(c, echo.Map{
		"msg": "ok",
	})
}

func (a *app) adminCreateThing(c echo.Context) error {
	typeOfThing, _ := strconv.Atoi(c.QueryParam("type"))
	req := &thingP.CreateThingRequest{
		Thing: &thingP.Thing{
			Type: int32(typeOfThing),
			Name: c.QueryParam("name"),
		},
	}
	res := &thingP.CreateThingResponse{}

	err := a.eb.RequestMessage(thingT.CreateThing, req, res, lib.DefaultTimeout)
	if err != nil {
		log.Printf("error: %+v", err)
		return a.makeError(c, http.StatusInternalServerError, err)
	}
	if !res.GetSuccess() {
		log.Printf("error: %+v", res.GetError())
		return a.makeError(c, http.StatusUnauthorized, errors.New(res.GetError()))
	}

	trReq := &ruleP.CreateThingLinkRequest{
		Thing:    res.GetId(),
		Channels: res.Channels,
	}
	trRes := &ruleP.CreateThingLinkResponse{}

	err = a.eb.RequestMessage(ruleT.CreateThingLink, trReq, trRes, lib.DefaultTimeout)
	if err != nil {
		log.Printf("error: %+v", err)
		return a.makeError(c, http.StatusInternalServerError, err)
	}
	if !res.GetSuccess() {
		log.Printf("error: %+v", res.GetError())
		return a.makeError(c, http.StatusUnauthorized, errors.New(res.GetError()))
	}

	return a.sendSucess(c, echo.Map{
		"msg": "ok",
		"id":  res.GetId(),
	})
}

func (a *app) adminGetAllThingsPDF(c echo.Context) error {
	req := &thingP.GetThingsByIDsRequest{}
	res := &thingP.GetThingsByIDsResponse{}

	err := a.eb.RequestMessage(thingT.GetThingsByIDs, req, res, lib.DefaultTimeout)
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

	fileName := lib.NewObjectID()
	pdf := gofpdf.New("P", "mm", "A4", "")

	pdf.SetCompression(false)
	pdf.AddPage()

	y := 10.0
	_, ph := pdf.GetPageSize()

	for _, thing := range res.Things {
		if y >= ph-40 {
			pdf.AddPage()
			y = 10
		}

		err = qrcode.WriteFile(thing.Id, qrcode.Medium, 256, "./pdfs/"+fileName+thing.Id+".png")
		if err != nil {
			log.Printf("error: %+v", err)
			return a.makeError(c, http.StatusInternalServerError, err)
		}

		pdf.Image("./pdfs/"+fileName+thing.Id+".png", pdf.GetX(), y, 30, 0, false, "", 0, "")
		pdf.SetFont("Arial", "B", 16)
		pdf.Text(45, y+10, thing.Id)
		pdf.Text(120, y+10, thing.Name)
		pdf.Rect(10, y, 30, 30, "")
		pdf.Rect(10+30, y, 70, 30, "")
		pdf.Rect(10+30+70, y, 80, 30, "")
		y = y + 30
	}

	err = pdf.OutputFileAndClose("./pdfs/" + fileName + ".pdf")
	if err != nil {
		log.Printf("error: %+v", err)
		return a.makeError(c, http.StatusInternalServerError, err)
	}

	files, err := filepath.Glob("./pdfs/" + fileName + "*.png")
	if err != nil {
		log.Printf("error: %+v", err)
		return a.makeError(c, http.StatusInternalServerError, err)
	}
	for _, f := range files {
		if err := os.Remove(f); err != nil {
			panic(err)
		}
	}

	return a.sendSucess(c, echo.Map{
		"msg": "ok",
		// "things": res.GetThings(),
		"file": fileName,
	})
}

func (a *app) adminGetThingTypes(c echo.Context) error {
	req := &thingP.GetThingTypesRequest{}
	res := &thingP.GetThingTypesResponse{}

	err := a.eb.RequestMessage(thingT.GetThingTypes, req, res, lib.DefaultTimeout)
	if err != nil {
		log.Printf("error: %+v", err)
		return a.makeError(c, http.StatusInternalServerError, err)
	}

	if !res.GetSuccess() {
		return a.makeError(c, http.StatusUnauthorized, errors.New(res.GetError()))
	}

	return a.sendSucess(c, echo.Map{
		"msg":   "ok",
		"types": res.Types,
	})
}
