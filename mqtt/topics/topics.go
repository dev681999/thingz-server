package topics

const (
	Base        = "server.Mqtt."
	CreateMqtt  = Base + "CreateMqtt"
	UserMqtts   = Base + "UserMqtts"
	UpdateThing = Base + "UpdateThing"
)

const (
	BaseMQTT      = "server"
	UpdateChannel = BaseMQTT + "/channel/update"
	UpdateDevice  = BaseMQTT + "/device/update/channel"
	// UpdateChannel = "test"
	// UserMqtts  = Base + "UserMqtts"
)
