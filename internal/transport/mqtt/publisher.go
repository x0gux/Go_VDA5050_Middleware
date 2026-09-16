package mqtt

import (
	"fmt"
	"log"
)

func (c *Client) Publish(topic string, qos byte, msg []byte) error {
	token := c.rawClient.Publish(topic, qos, false, msg)
	if token.Wait() && token.Error() != nil {
		return fmt.Errorf("Publish failed [%s]: %w", topic, token.Error())
	}
	log.Printf("Publish to topic: %s", topic)
	return nil
}
