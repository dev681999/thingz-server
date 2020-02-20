package googlesmarthome

const (
	TraitOnOff      = "action.devices.traits.OnOff"
	TraitScene      = "action.devices.traits.Scene"
	TraitBrightness = "action.devices.traits.Brightness"
	TraitFanSpeed   = "action.devices.traits.FanSpeed"
)

var traitMap = map[string]string{
	"onOff":      TraitOnOff,
	"brightness": TraitBrightness,
	"Speed":      TraitFanSpeed,
}

func ConvertCommandToTraits(cmds []string) []string {
	res := []string{}
	for _, cmd := range cmds {
		res = append(res, traitMap[cmd])
	}

	return res
}
