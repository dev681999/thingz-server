package channel

// Type represents type of channel
type Type int32

const (
	TypeHumidity     = Type(iota)
	TypeTemperatrure = Type(iota)
	TypeMoisture     = Type(iota)
	TypeSwitch       = Type(iota)
	TypeButton       = Type(iota)
	TypeMotion       = Type(iota)
)
