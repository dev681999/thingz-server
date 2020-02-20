package main

import (
	"encoding/json"
	"log"
	proto "thingz-server/thing/proto"
)

type thingInfo struct {
	Name     string
	IsSensor bool
	Unit     proto.Unit
	Channels []*proto.Channel
}

const (
	thingTypeLightFan = 100
)

func init() {
	things := []proto.Thing{}

	json.Unmarshal([]byte(`[
		{
			"id": "SENALCO0001",
			"name": "Alcohol Detector",
			"key": "1234",
			"project": "",
			"type": 1,
			"channels": [
				{
					"id": "sensor",
					"name": "sensor",
					"unit": 1,
					"isSensor": true,
					"floatValue": 0
				}
			]
		},
		{
			"id": "SENACCL0001",
			"name": "Accelerometer",
			"key": "1234",
			"project": "",
			"type": 2,
			"channels": [
				{
					"id": "sensor",
					"name": "sensor",
					"unit": 1,
					"isSensor": true,
					"floatValue": 0
				}
			]
		},
		{
			"id": "SENCURR0001",
			"name": "Current meter",
			"key": "1234",
			"project": "",
			"type": 3,
			"channels": [
				{
					"id": "sensor",
					"name": "sensor",
					"unit": 1,
					"isSensor": true,
					"floatValue": 0
				}
			]
		},
		{
			"id": "SENSDHT0001",
			"name": "Temprature and humidity",
			"key": "1234",
			"project": "",
			"type": 4,
			"channels": [
				{
					"id": "sensor",
					"name": "sensor",
					"unit": 1,
					"isSensor": true,
					"floatValue": 0
				}
			]
		},
		{
			"id": "SENRFID0001",
			"name": "EM- 18 Rfid",
			"key": "1234",
			"project": "",
			"type": 5,
			"channels": [
				{
					"id": "sensor",
					"name": "sensor",
					"unit": 1,
					"isSensor": true,
					"floatValue": 0
				}
			]
		},
		{
			"id": "SENFLME0001",
			"name": "Flame Detector",
			"key": "1234",
			"project": "",
			"type": 6,
			"channels": [
				{
					"id": "sensor",
					"name": "sensor",
					"unit": 1,
					"isSensor": true,
					"floatValue": 0
				}
			]
		},
		{
			"id": "SENGYRO0001",
			"name": "Gyroscope",
			"key": "1234",
			"project": "",
			"type": 7,
			"channels": [
				{
					"id": "sensor",
					"name": "sensor",
					"unit": 1,
					"isSensor": true,
					"floatValue": 0
				}
			]
		},
		{
			"id": "SENUGPS0001",
			"name": "GPS",
			"key": "1234",
			"project": "",
			"type": 8,
			"channels": [
				{
					"id": "sensor",
					"name": "sensor",
					"unit": 1,
					"isSensor": true,
					"floatValue": 0
				}
			]
		},
		{
			"id": "SENIRPRX0001",
			"name": "Infrared Proximity",
			"key": "1234",
			"project": "",
			"type": 9,
			"channels": [
				{
					"id": "sensor",
					"name": "sensor",
					"unit": 1,
					"isSensor": true,
					"floatValue": 0
				}
			]
		},
		{
			"id": "SENIMSW0001",
			"name": "Impact switch",
			"key": "1234",
			"project": "",
			"type": 10,
			"channels": [
				{
					"id": "sensor",
					"name": "sensor",
					"unit": 1,
					"isSensor": true,
					"floatValue": 0
				}
			]
		},
		{
			"id": "SENRLDR0001",
			"name": "LDR",
			"key": "1234",
			"project": "",
			"type": 11,
			"channels": [
				{
					"id": "sensor",
					"name": "sensor",
					"unit": 1,
					"isSensor": true,
					"floatValue": 0
				}
			]
		},
		{
			"id": "SENDS180001",
			"name": "Liquid temperature",
			"key": "1234",
			"project": "",
			"type": 12,
			"channels": [
				{
					"id": "sensor",
					"name": "sensor",
					"unit": 1,
					"isSensor": true,
					"floatValue": 0
				}
			]
		},
		{
			"id": "SENSMIC0001",
			"name": "Mic",
			"key": "1234",
			"project": "",
			"type": 13,
			"channels": [
				{
					"id": "sensor",
					"name": "sensor",
					"unit": 1,
					"isSensor": true,
					"floatValue": 0
				}
			]
		},
		{
			"id": "SENSLMS0001",
			"name": "Soil Moisture Proportion",
			"key": "1234",
			"project": "",
			"type": 14,
			"channels": [
				{
					"id": "sensor",
					"name": "sensor",
					"unit": 1,
					"isSensor": true,
					"floatValue": 0
				}
			]
		},
		{
			"id": "SENSPIR0001",
			"name": "PIR",
			"key": "1234",
			"project": "",
			"type": 15,
			"channels": [
				{
					"id": "sensor",
					"name": "sensor",
					"unit": 0,
					"isSensor": true,
					"boolValue": false
				}
			]
		},
		{
			"id": "SENSPOT0001",
			"name": "Potentiometer",
			"key": "1234",
			"project": "",
			"type": 16,
			"channels": [
				{
					"id": "sensor",
					"name": "sensor",
					"unit": 1,
					"isSensor": true,
					"floatValue": 0
				}
			]
		},
		{
			"id": "SENPRES0001",
			"name": "Pressure sense",
			"key": "1234",
			"project": "",
			"type": 17,
			"channels": [
				{
					"id": "sensor",
					"name": "sensor",
					"unit": 1,
					"isSensor": true,
					"floatValue": 0
				}
			]
		},
		{
			"id": "SENSPH0001",
			"name": "PH SENSOR",
			"key": "1234",
			"project": "",
			"type": 18,
			"channels": [
				{
					"id": "sensor",
					"name": "sensor",
					"unit": 1,
					"isSensor": true,
					"floatValue": 0
				}
			]
		},
		{
			"id": "SENPUSH0001",
			"name": "Push Button",
			"key": "1234",
			"project": "",
			"type": 19,
			"channels": [
				{
					"id": "sensor",
					"name": "sensor",
					"unit": 0,
					"isSensor": true,
					"boolValue": false
				}
			]
		},
		{
			"id": "SENSMKE0001",
			"name": "Smoke detector",
			"key": "1234",
			"project": "",
			"type": 20,
			"channels": [
				{
					"id": "sensor",
					"name": "sensor",
					"unit": 1,
					"isSensor": true,
					"floatValue": 0
				}
			]
		},
		{
			"id": "SENSWBU0001",
			"name": "Switch Button",
			"key": "1234",
			"project": "",
			"type": 21,
			"channels": [
				{
					"id": "sensor",
					"name": "sensor",
					"unit": 0,
					"isSensor": true,
					"boolValue": false
				}
			]
		},
		{
			"id": "SENTHBU0001",
			"name": "Touch Button",
			"key": "1234",
			"project": "",
			"type": 22,
			"channels": [
				{
					"id": "sensor",
					"name": "sensor",
					"unit": 0,
					"isSensor": true,
					"boolValue": false
				}
			]
		},
		{
			"id": "SENULTR0001",
			"name": "Ultrasonic",
			"key": "1234",
			"project": "",
			"type": 23,
			"channels": [
				{
					"id": "sensor",
					"name": "sensor",
					"unit": 1,
					"isSensor": true,
					"floatValue": 0
				}
			]
		},
		{
			"id": "SENVIBR0001",
			"name": "Vibration",
			"key": "1234",
			"project": "",
			"type": 24,
			"channels": [
				{
					"id": "sensor",
					"name": "sensor",
					"unit": 1,
					"isSensor": true,
					"floatValue": 0
				}
			]
		},
		{
			"id": "SENVOLT0001",
			"name": "Voltmeter",
			"key": "1234",
			"project": "",
			"type": 25,
			"channels": [
				{
					"id": "sensor",
					"name": "sensor",
					"unit": 1,
					"isSensor": true,
					"floatValue": 0
				}
			]
		},
		{
			"id": "SENIRBL0001",
			"name": "IR Blaster",
			"key": "1234",
			"project": "",
			"type": 26,
			"channels": [
				{
					"id": "sensor",
					"name": "sensor",
					"unit": 1,
					"isSensor": true,
					"floatValue": 0
				}
			]
		},
		{
			"id": "SENJOYS0001",
			"name": "Joystick",
			"key": "1234",
			"project": "",
			"type": 27,
			"channels": [
				{
					"id": "sensor",
					"name": "sensor",
					"unit": 1,
					"isSensor": true,
					"floatValue": 0
				}
			]
		},
		{
			"id": "SENHRTB0001",
			"name": "Heartbeat",
			"key": "1234",
			"project": "",
			"type": 28,
			"channels": [
				{
					"id": "sensor",
					"name": "sensor",
					"unit": 1,
					"isSensor": true,
					"floatValue": 0
				}
			]
		},
		{
			"id": "ACTELVA0001",
			"name": "Electronic Valve",
			"key": "1234",
			"project": "",
			"type": 29,
			"channels": [
				{
					"id": "value",
					"name": "Status",
					"unit": 0,
					"isSensor": false,
					"boolValue": true
				}
			]
		},
		{
			"id": "ACTSVMO0001",
			"name": "Servo Motor",
			"key": "1234",
			"project": "",
			"type": 30,
			"channels": [
				{
					"id": "value",
					"name": "Status",
					"unit": 0,
					"isSensor": false,
					"boolValue": true
				}
			]
		},
		{
			"id": "ACTSTMO0001",
			"name": "Stepper Motor",
			"key": "1234",
			"project": "",
			"type": 31,
			"channels": [
				{
					"id": "value",
					"name": "Status",
					"unit": 0,
					"isSensor": false,
					"boolValue": true
				}
			]
		},
		{
			"id": "ACTRLSW0001",
			"name": "Relay switch",
			"key": "1234",
			"project": "",
			"type": 32,
			"channels": [
				{
					"id": "value",
					"name": "Status",
					"unit": 0,
					"isSensor": false,
					"boolValue": true
				}
			]
		},
		{
			"id": "ACTDLED0001",
			"name": "LED",
			"key": "1234",
			"project": "",
			"type": 33,
			"channels": [
				{
					"id": "value",
					"name": "Status",
					"unit": 0,
					"isSensor": false,
					"boolValue": true
				}
			]
		},
		{
			"id": "ACTBUZZ0001",
			"name": "Buzzer",
			"key": "1234",
			"project": "",
			"type": 34,
			"channels": [
				{
					"id": "value",
					"name": "Status",
					"unit": 0,
					"isSensor": false,
					"boolValue": true
				}
			]
		},
		{
			"id": "ACTBUZZ0001",
			"name": "HUB",
			"key": "1234",
			"project": "",
			"type": 35,
			"channels": [
				{
					"id": "sensor_1_1",
					"name": "Sensor 1",
					"unit": 1,
					"isSensor": true
				},
				{
					"id": "sensor_1_2",
					"name": "Sensor 1",
					"unit": 0,
					"isSensor": true
				},
				{
					"id": "sensor_2",
					"name": "Sensor 2",
					"unit": 0,
					"isSensor": true
				},
				{
					"id": "sensor_3",
					"name": "Sensor 3",
					"unit": 0,
					"isSensor": true
				},
				{
					"id": "sensor_4",
					"name": "Sensor 4",
					"unit": 0,
					"isSensor": true
				},
				{
					"id": "actuator_1",
					"name": "Actuator 1",
					"unit": 0
				},
				{
					"id": "actuator_2",
					"name": "Actuator 2",
					"unit": 0
				},
				{
					"id": "actuator_3",
					"name": "Actuator 3",
					"unit": 0
				},
				{
					"id": "actuator_4",
					"name": "Actuator 4",
					"unit": 0
				},
				{
					"id": "analog_1_1",
					"name": "Analog",
					"unit": 1,
					"isSensor": true
				},
				{
					"id": "analog_1_2",
					"name": "Analog",
					"unit": 1
				}
			]
		},
		{
			"id": "",
			"name": "GPS",
			"type": 36,
			"channels": [
				{
					"id": "value",
					"name": "Coordinates",
					"unit": 2,
					"isSensor": true
				}
			]
		},
		{
			"id": "",
			"name": "RFID",
			"type": 37,
			"channels": [
				{
					"id": "value",
					"name": "Tag ID",
					"unit": 2,
					"isSensor": true
				}
			]
		},
		{
			"id": "",
			"name": "Light",
			"type": 38,
			"channels": [
				{
					"id": "onOff",
					"name": "OnOff",
					"unit": 0
				},
				{
					"id": "brightness",
					"name": "Brightness",
					"unit": 1
				}
			]
		},
		{
			"id": "",
			"name": "Fan",
			"type": 39,
			"channels": [
				{
					"id": "onOff",
					"name": "OnOff",
					"unit": 0
				},
				{
					"id": "Speed",
					"name": "Speed",
					"unit": 1
				}
			]
		},
		{
			"id": "",
			"name": "Motion Sensor",
			"type": 40,
			"channels": [
				{
					"id": "onOff",
					"name": "OnOff",
					"unit": 0,
					"isSensor": true
				}
			]
		}
	]`), &things)

	thingTypes = []thingInfo{}

	for idx, t := range things {
		t.Type = int32(idx)
		// fmt.Printf("%+v\n", t)
		thingTypes = append(thingTypes, thingInfo{
			Name: t.Name,
			// IsSensor: t.Channels[0].IsSensor,
			// Unit:     t.Channels[0].Unit,
			Channels: t.Channels,
		})
		log.Println("ID:", t.Type, "Name:", t.Name)
	}
}
