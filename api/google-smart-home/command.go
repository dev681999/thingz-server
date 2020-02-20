package googlesmarthome

const (
	commandChangeColor        = "action.devices.commands.ChangeColor"
	commandOnOff              = "action.devices.commands.OnOff"
	commandActivateScene      = "action.devices.commands.ActivateScene"
	commandBrightness         = "action.devices.commands.Brightness"
	commandBrightnessAbsolute = "action.devices.commands.BrightnessAbsolute"
	commandFanSpeed           = "action.devices.commands.SetFanSpeed"
)

var CommandMap = map[string]string{
	commandBrightness:         "brightness",
	commandBrightnessAbsolute: "brightness",
	commandFanSpeed:           "fanSpeed",
	commandOnOff:              "onOff",
}

var UnitMap = map[string]int{
	commandBrightness:         1,
	commandBrightnessAbsolute: 1,
	commandFanSpeed:           1,
	commandOnOff:              0,
}

var CommandConvertParam = map[string]string{
	commandBrightness:         "brightness",
	commandBrightnessAbsolute: "brightness",
	commandChangeColor:        "color",
	commandFanSpeed:           "fanSpeed",
	commandOnOff:              "on",
}

/* var CommandConvertParam = map[string]string{
	"brightness": "brightness",
	"fanSpeed":   "fanSpeed",
	"onOff":      "on",
} */
