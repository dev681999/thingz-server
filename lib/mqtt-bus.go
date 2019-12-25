package lib

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/golang/protobuf/proto"
)

// Msg is Mesasge structure
type Msg struct {
	Topic string
	Msg   []byte
}

// MqttData is a special Data Type to auto decode/encode protobufs
type MqttData []byte

// Decode a byte array to protobuf message
func (m MqttData) Decode(data proto.Message) error {
	return proto.Unmarshal([]byte(m), data)
}

// Encode a byte array to protobuf message
func encode(data proto.Message) ([]byte, error) {
	return proto.Marshal(data)
}

// MqttProtocolBus is a MQTT protocol bus
type MqttProtocolBus struct {
	ServerURL         string
	ClientID          string
	Password          string
	UserName          string
	CaCert            string
	ClientCert        string
	ClientKey         string
	defaultMsgHandler func(topic string, payload []byte)

	mqttClient mqtt.Client
}

// NewMqttProtocolBus :
func NewMqttProtocolBus(serverURL, clientID string, defaultMsgHandler func(topic string, payload []byte)) *MqttProtocolBus {
	return &MqttProtocolBus{
		ClientID:          clientID,
		defaultMsgHandler: defaultMsgHandler,
		ServerURL:         serverURL,
	}
}

// NewMqttProtocolBusSecure :
func NewMqttProtocolBusSecure(serverURL, clientID, password, caCert, userName, clientCert, clientKey string, defaultMsgHandler func(topic string, payload []byte)) *MqttProtocolBus {
	return &MqttProtocolBus{
		CaCert:            caCert,
		ClientCert:        clientCert,
		ClientID:          clientID,
		ClientKey:         clientKey,
		defaultMsgHandler: defaultMsgHandler,
		ServerURL:         serverURL,
	}
}

// Connect to MQTT broker
func (p *MqttProtocolBus) Connect() error {
	opts := mqtt.NewClientOptions()

	opts.AddBroker(p.ServerURL)
	opts.SetClientID(p.ClientID)

	var f mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
		p.defaultMsgHandler(msg.Topic(), msg.Payload())

	}

	opts.SetDefaultPublishHandler(f)

	// Start the connection
	p.mqttClient = mqtt.NewClient(opts)
	if token := p.mqttClient.Connect(); token.Wait() && token.Error() != nil {
		return token.Error()
	}

	return nil
}

// ConnectSecure to MQTT broker
func (p *MqttProtocolBus) ConnectSecure(secure ...bool) error {
	opts := mqtt.NewClientOptions()

	opts.AddBroker(p.ServerURL)

	if len(secure) > 0 && secure[0] {
		certpool := x509.NewCertPool()
		pemCerts, err := ioutil.ReadFile(p.CaCert)
		if err == nil {
			certpool.AppendCertsFromPEM(pemCerts)
		}

		// Import client certificate/key pair
		cert, err := tls.LoadX509KeyPair(p.ClientCert, p.ClientKey)
		if err != nil {
			return (err)
		}

		// Just to print out the client certificate..
		cert.Leaf, err = x509.ParseCertificate(cert.Certificate[0])
		if err != nil {
			return (err)
		}

		// Create tls.Config with desired tls properties
		tls := &tls.Config{
			// RootCAs = certs used to verify server cert.
			RootCAs: certpool,
			// ClientAuth = whether to request cert from server.
			// Since the server is set up for SSL, this happens
			// anyways.
			ClientAuth: tls.NoClientCert,
			// ClientCAs = certs used to validate client cert.
			ClientCAs: nil,
			// InsecureSkipVerify = verify that cert contents
			// match server. IP matches what is in cert etc.
			InsecureSkipVerify: true,
			// Certificates = list of certs client sends to server.
			Certificates: []tls.Certificate{cert},
		}

		opts.SetClientID(p.ClientID).SetTLSConfig(tls)

		if len(secure) > 1 && secure[1] {
			opts.SetUsername(p.UserName)
			opts.SetPassword(p.Password)
		}
	} else {
		// opts.SetClientID(p.clientID)
	}

	var f mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
		p.defaultMsgHandler(msg.Topic(), msg.Payload())
	}

	opts.SetDefaultPublishHandler(f)

	// Start the connection
	p.mqttClient = mqtt.NewClient(opts)
	if token := p.mqttClient.Connect(); token.Wait() && token.Error() != nil {
		return token.Error()
	}

	return nil
}

// Close MQTT connection
func (p *MqttProtocolBus) Close() {
	p.mqttClient.Disconnect(100)
}

// RegisterListener too MQTT broker
func (p *MqttProtocolBus) RegisterListener(topic string, listener func(*Msg)) error {
	if token := p.mqttClient.Subscribe(topic, 2, func(client mqtt.Client, msg mqtt.Message) {
		listener(&Msg{
			Topic: msg.Topic(),
			Msg:   msg.Payload(),
		})
	}); token.Wait() && token.Error() != nil {
		return token.Error()
	}

	return nil
}

// SendMessage to specified topic
func (p *MqttProtocolBus) SendMessage(topic string, data []byte) error {
	/* data, err := proto.Marshal(payload)
	if err != nil {
		return err
	} */

	/* if token := ...  ;token.Wait() && token.Error() != nil {
		return token.Error()
	} */

	p.mqttClient.Publish(topic, byte(2), false, data)
	return nil
}

// DeregisterListener from broker
func (p *MqttProtocolBus) DeregisterListener(topic string) error {
	t := p.mqttClient.Unsubscribe(topic)
	t.Wait()
	return t.Error()
}
