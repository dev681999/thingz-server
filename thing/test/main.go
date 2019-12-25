package main

import (
	"math/rand"
	"os"
	"path/filepath"

	"github.com/jung-kurt/gofpdf"
	"github.com/skip2/go-qrcode"
)

func loremList() []string {
	return []string{
		"Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod " +
			"tempor incididunt ut labore et dolore magna aliqua.",
		"Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut " +
			"aliquip ex ea commodo consequat.",
		"Duis aute irure dolor in reprehenderit in voluptate velit esse cillum " +
			"dolore eu fugiat nulla pariatur.",
		"Excepteur sint occaecat cupidatat non proident, sunt in culpa qui " +
			"officia deserunt mollit anim id est laborum.",
	}
}

const letterBytes = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

var (
	thingTypes = [][]string{
		{"Alcohol Detector", "true"},
		{"Accelerometer", "true"},
		{"Current meter", "true"},
		{"Temprature and humidity", "true"},
		{"EM- 18 Rfid", "true"},
		{"Flame Detector", "true"},
		{"Gyroscope", "true"},
		{"GPS", "true"},
		{"Infrared Proximity", "true"},
		{"Impact switch", "true"},
		{"LDR", "true"},
		{"Liquid temperature", "true"},
		{"Mic", "true"},
		{"Soil Moisture Proportion", "true"},
		{"PIR", "true"},
		{"Potentiometer", "true"},
		{"Pressure sense", "true"},
		{"PH SENSOR", "true"},
		{"Push Button", "true"},
		{"Smoke detector", "true"},
		{"Switch Button", "true"},
		{"Touch Button", "true"},
		{"Ultrasonic", "true"},
		{"Vibration", "true"},
		{"Voltmeter", "true"},
		{"IR Blaster", "true"},
		{"Joystick", "true"},
		{"Heartbeat", "true"},
		{"Electronic Valve", "false"},
		{"Servo Motor", "false"},
		{"Stepper Motor", "false"},
		{"Relay switch", "false"},
		{"LED", "false"},
		{"Buzzer", "false"},
		{"HUB", ""},
	}
	thingTypeToInt = map[string]int32{}
)

func init() {
	for idx, ty := range thingTypes {
		thingTypeToInt[ty[0]] = int32(idx)
	}
}

func main() {
	pdf := gofpdf.New("P", "mm", "A4", "")

	pdf.SetCompression(false)
	pdf.AddPage()

	y := 10.0
	_, ph := pdf.GetPageSize()

	for i := 0; i < 50; i++ {
		if y >= ph-40 {
			pdf.AddPage()
			y = 10
		}

		id := randStringBytes(8)

		qrcode.WriteFile(id, qrcode.Medium, 256, id+".png")

		pdf.Image(id+".png", pdf.GetX(), y, 30, 0, false, "", 0, "")
		pdf.SetFont("Arial", "B", 16)
		pdf.Text(45, y+10, id)
		pdf.Text(120, y+10, "Relay")
		pdf.Rect(10, y, 30, 30, "")
		pdf.Rect(10+30, y, 70, 30, "")
		pdf.Rect(10+30+70, y, 60, 30, "")
		y = y + 30
	}

	// pdf.UseTemplate(template)

	err := pdf.OutputFileAndClose("qr_table_test.pdf")
	if err != nil {
		panic(err)
	}

	files, err := filepath.Glob("*.png")
	if err != nil {
		panic(err)
	}
	for _, f := range files {
		if err := os.Remove(f); err != nil {
			panic(err)
		}
	}
}
