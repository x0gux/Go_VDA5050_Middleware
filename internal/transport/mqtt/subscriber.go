package mqtt

import (
	"fmt"
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

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
