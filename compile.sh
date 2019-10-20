#!/bin/bash

cd api
go build -o ../build/api

cd ../mqtt
go build -o ../build/mqtt

cd ../project
go build -o ../build/project

cd ../thing
go build -o ../build/thing

cd ../user
go build -o ../build/user

cd ../rule
go build -o ../build/rule