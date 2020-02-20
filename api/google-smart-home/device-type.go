package googlesmarthome

const (
	typeFan    = "action.devices.types.FAN"
	typeLight  = "action.devices.types.LIGHT"
	typeOutlet = "action.devices.types.OUTLET"
	TypeScene  = "action.devices.types.SCENE"
)

var DeviceTypeMap = map[int]string{
	38: typeFan,
	37: typeLight,
	// lib.DeviceTypeOutlet: typeOutlet,
}
