package mqtt

import mqtt "github.com/eclipse/paho.mqtt.golang"

type MessageHandler func(topic string, payload []byte)

type Client struct {
	rawClient mqtt.Client
}
