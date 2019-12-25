package main

import (
	lib "github.com/dev681999/helperlibs"
)

func main() {
	s := &lib.Store{
		Address:  "localhost:27017",
		Username: "",
		Password: "",
		Database: "thingz-test",
	}

	err := s.Connect()
	if err != nil {
		panic(err)
	}

	defer s.Close()

	db, err := s.GetMongoSession()
	if err != nil {
		panic(err)
	}

	err = db.DB("").C("test").Insert(map[string]interface{}{
		"_id": "1234",
	})
	if err != nil {
		panic(err)
	}
}
