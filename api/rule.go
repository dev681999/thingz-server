package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"thingz-server/lib"
	ruleP "thingz-server/rule/proto"
	ruleT "thingz-server/rule/topics"
	thingP "thingz-server/thing/proto"
	thingT "thingz-server/thing/topics"

	log "github.com/sirupsen/logrus"

	"github.com/labstack/echo"
)

func (a *app) projectRules(c echo.Context) error {
	project := c.Param("id")
	req := &ruleP.ProjectRulesRequest{
		Project: project,
	}
	res := &ruleP.ProjectRulesResponse{}

	err := a.eb.RequestMessage(ruleT.ProjectRules, req, res, lib.DefaultTimeout)
	if err != nil {
		log.Printf("error: %+v", err)
		return a.makeError(c, http.StatusInternalServerError, err)
	}

	if !res.GetSuccess() {
		return a.makeError(c, http.StatusUnauthorized, errors.New(res.GetError()))
	}

	if res.Rules == nil {
		res.Rules = []*ruleP.Rule{}
	}

	return a.sendSucess(c, echo.Map{
		"msg":   "ok",
		"rules": res.GetRules(),
	})
}

func (a *app) createRule(c echo.Context) error {
	rule := &ruleP.Rule{}

	err := c.Bind(rule)
	if err != nil {
		log.Println("bind error")
		return a.makeError(c, http.StatusBadRequest, err)
	}

	ids := rule.GetAllThings()
	thingsMapReq := &thingP.GetThingsByIDsRequest{
		Ids: ids,
	}

	thingsMapRes := &thingP.GetThingsByIDsResponse{}

	err = a.eb.RequestMessage(thingT.GetThingsByIDs, thingsMapReq, thingsMapRes, lib.DefaultTimeout)
	if err != nil {
		log.Printf("error get things: %+v", err)
		return a.makeError(c, http.StatusInternalServerError, err)
	}

	thingsMap := map[string]*thingP.Thing{}
	for _, t := range thingsMapRes.Things {
		thingsMap[t.Id] = t
	}

	b, _ := json.Marshal(thingsMap)

	req := &ruleP.CreateRuleRequest{
		Rule:      rule,
		ThingsMap: string(b),
	}

	res := &ruleP.CreateRuleResponse{}

	err = a.eb.RequestMessage(ruleT.CreateRule, req, res, lib.DefaultTimeout)
	if err != nil {
		log.Printf("error create rule: %+v", err)
		return a.makeError(c, http.StatusInternalServerError, err)
	}

	if !res.GetSuccess() {
		return a.makeError(c, http.StatusUnauthorized, errors.New(res.GetError()))
	}
	return a.sendSucess(c, echo.Map{
		"msg": "ok",
		"id":  res.GetId(),
	})
}

func (a *app) deleteRule(c echo.Context) error {
	rule := c.Param("id")
	req := &ruleP.DeleteRuleRequest{
		Rule: rule,
	}
	res := &ruleP.DeleteRuleResponse{}

	err := a.eb.RequestMessage(ruleT.DeleteRule, req, res, lib.DefaultTimeout)
	if err != nil {
		log.Printf("error: %+v", err)
		return a.makeError(c, http.StatusInternalServerError, err)
	}

	if !res.GetSuccess() {
		return a.makeError(c, http.StatusUnauthorized, errors.New(res.GetError()))
	}

	return a.sendSucess(c, echo.Map{
		"msg": "ok",
	})
}

func (a *app) fireRule(c echo.Context) error {
	r := &ruleP.Rule{}
	c.Bind(r)
	req := &ruleP.CheckThingRuleRequest{
		Update: r,
	}

	res := &ruleP.CheckThingRuleResponse{}

	err := a.eb.RequestMessage(ruleT.CheckThingRule, req, res, lib.DefaultTimeout)
	if err != nil {
		log.Printf("error create rule: %+v", err)
		return a.makeError(c, http.StatusInternalServerError, err)
	}

	if !res.GetSuccess() {
		return a.makeError(c, http.StatusUnauthorized, errors.New(res.GetError()))
	}
	return a.sendSucess(c, echo.Map{
		"msg": "ok",
		"res": res,
	})
}
