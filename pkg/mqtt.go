package pkg

import (
	"fmt"
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MessageHandler func(topic string, payload []byte)

type Client struct {
	rawClient mqtt.Client
}

func NewClient(broker, clientID, username, password string) (*Client, error) {
	opts := mqtt.NewClientOptions().
		AddBroker(broker).
		SetClientID(clientID).
		SetUsername(username).
		SetPassword(password).
		SetAutoReconnect(true).
		SetKeepAlive(60 * time.Second)

	c := mqtt.NewClient(opts)
	token := c.Connect()
	if token.Wait() && token.Error() != nil {
		return nil, fmt.Errorf("MQTT connection failed [%s]: %w", broker, token.Error())
	}

	return &Client{rawClient: c}, nil
}

// Subscribe: 토픽 구독 및 범용 OnMessage 콜백 연결
func (c *Client) Subscribe(topic string, qos byte, handler MessageHandler) error {
	token := c.rawClient.Subscribe(topic, qos, func(_ mqtt.Client, msg mqtt.Message) {
		handler(msg.Topic(), msg.Payload())
	})

	if token.Wait() && token.Error() != nil {
		return fmt.Errorf("subscribe failed [%s]: %w", topic, token.Error())
	}
	log.Printf("Subscribed to topic: %s", topic)
	return nil
}
