#!/bin/bash

cd mqtt/proto
make proto

echo MQTT 

cd ../../api/proto
make proto

echo api

cd ../../project/proto
make proto

echo Project

cd ../../thing/proto
make proto

echo Thing

cd ../../user/proto
make proto

echo User 

cd ../../rule/proto
make proto

echo RULE
