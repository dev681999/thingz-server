package lib

import (
	"time"

	nats "github.com/nats-io/nats.go"
	penc "github.com/nats-io/nats.go/encoders/protobuf"
)

func init() {
	nats.RegisterEncoder(penc.PROTOBUF_ENCODER, &penc.ProtobufEncoder{})
}

const (
	// JSONEnc is JSON Encoded NATS connection
	JSONEnc = nats.JSON_ENCODER

	// ProtobufEnc is protobuf Encoded NATS connection
	ProtobufEnc = penc.PROTOBUF_ENCODER
)

// Listener is a listener to a topic on event bus
type Listener struct {
	Topic string      `json:"topic,omitempty"`
	Func  interface{} `json:"func,omitempty"`
}

// EventBus is used to send & receive events from various services in the syste
type EventBus struct {
	nc         *nats.Conn
	enc        *nats.EncodedConn
	clientCert string
	clientKey  string
	caCert     string
	serverURL  string
}

// MsgHandler is a callback function that processes messages
type MsgHandler func(subject, reply string, msg interface{})

// NewEventBus return a new EventBus handler
func NewEventBus(serverURL, clientCert, clientKey, caCert string) *EventBus {
	return &EventBus{
		serverURL:  serverURL,
		clientCert: clientCert,
		clientKey:  clientKey,
		caCert:     caCert,
	}
}

// NewEventBusUnsecure return a new EventBus handler without security
func NewEventBusUnsecure(serverURL string) *EventBus {
	return &EventBus{
		serverURL:  serverURL,
		clientCert: "",
		clientKey:  "",
		caCert:     "",
	}
}

// SecureConnect connects to the event bus securly and sets up encodded conenction with TLS
func (b *EventBus) SecureConnect(encoder string) error {
	cert := nats.ClientCert(b.clientCert, b.clientKey)
	var err error

	b.nc, err = nats.Connect(b.serverURL, cert, nats.RootCAs(b.caCert))
	if err != nil {
		return err
	}

	b.enc, err = nats.NewEncodedConn(b.nc, encoder)
	return err
}

// Connect connects to the event bus securly and sets up encodded conenction without TLS
func (b *EventBus) Connect(encoder string) error {
	var err error
	b.nc, err = nats.Connect(b.serverURL)

	if err != nil {
		return err
	}

	b.enc, err = nats.NewEncodedConn(b.nc, encoder)
	return err
}

// Close closes the connection to the event bus
func (b *EventBus) Close() {
	if b.nc != nil && b.enc != nil && !b.nc.IsClosed() {
		b.enc.Close()
		b.nc.Close()
	}
}

// RegisterListeners registers an array of message handler for specified topics
func (b *EventBus) RegisterListeners(listeners []Listener) error {
	for _, listener := range listeners {
		if err := b.RegisterListener(listener.Topic, listener.Func); err != nil {
			return err
		}
	}

	return nil
}

// SendMessage sends a message on the event bus on specified topic
func (b *EventBus) SendMessage(topic string, msg interface{}) error {
	return b.enc.Publish(topic, msg)
}

// RequestMessage requests a message on the event bus and waits for reply until specified duration
func (b *EventBus) RequestMessage(topic string, req interface{}, rep interface{}, timeout time.Duration) error {
	return b.enc.Request(topic, req, rep, timeout)
}

// SendByteMessage sends a message on the event bus on specified topic
func (b *EventBus) SendByteMessage(topic string, msg []byte) error {
	return b.nc.Publish(topic, msg)
}

// RequestByteMessage requests a message on the event bus and waits for reply until specified duration
func (b *EventBus) RequestByteMessage(topic string, req []byte, timeout time.Duration) ([]byte, error) {
	m, err := b.nc.Request(topic, req, timeout)
	if m == nil {
		return []byte(""), err
	}
	return m.Data, err
}

// RegisterListener registers a message handler for specified topic
func (b *EventBus) RegisterListener(topic string, h interface{}) error {
	_, err := b.enc.Subscribe(topic, h)
	return err
}

// Ping pings the NATS server
func (b *EventBus) Ping() error {
	return b.nc.Flush()
}

// PingTimeout pings the NATS server waits until timeout
func (b *EventBus) PingTimeout(duration time.Duration) error {
	return b.nc.FlushTimeout(duration)
}
