package main

import (
	"errors"
	"log"
	"net/http"
	"thingz-server/lib"
	userP "thingz-server/user/proto"
	userT "thingz-server/user/topics"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func (a *app) login(c echo.Context) error {
	u := &userP.User{}
	err := c.Bind(u)
	if err != nil {
		log.Println("bind error")
		return a.makeError(c, http.StatusBadRequest, err)
	}

	req := &userP.VerifyUserRequest{
		Email:    u.GetEmail(),
		Password: u.GetPassword(),
	}

	res := &userP.VerifyUserResponse{}

	err = a.eb.RequestMessage(userT.VerifyUser, req, res, lib.DefaultTimeout)
	if err != nil {
		log.Printf("error: %+v", err)
		return a.makeError(c, http.StatusInternalServerError, err)
	}

	if !res.GetSuccess() {
		return a.makeError(c, http.StatusUnauthorized, errors.New(res.GetError()))
	}

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = res.User.Name
	claims["id"] = res.User.Id
	// claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(a.c.JwtSecret))
	if err != nil {
		return err
	}

	return a.sendSucess(c, map[string]string{
		"token": t,
		"name":  res.User.Name,
	})
}

func (a *app) register(c echo.Context) error {
	u := &userP.User{}
	err := c.Bind(u)
	if err != nil {
		log.Println("bind error")
		return a.makeError(c, http.StatusBadRequest, err)
	}

	req := &userP.CreateUserRequest{
		User: u,
	}

	res := &userP.CreateUserResponse{}

	err = a.eb.RequestMessage(userT.CreateUser, req, res, lib.DefaultTimeout)
	if err != nil {
		log.Printf("error: %+v", err)
		return a.makeError(c, http.StatusInternalServerError, err)
	}

	if !res.GetSuccess() {
		return a.makeError(c, http.StatusBadRequest, errors.New(res.GetError()))
	}

	return a.sendSucess(c, map[string]string{
		"msg": "ok",
	})
}

func getUserFromContext(c echo.Context) jwt.MapClaims {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	return claims
}
