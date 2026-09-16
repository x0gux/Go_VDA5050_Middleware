package mqtt

import (
	"fmt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

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
