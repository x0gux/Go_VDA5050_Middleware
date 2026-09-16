package mqtt

import (
	"log"
	"strings"
)

type HandlerFunc func(vendor, serial, msgType string, payload []byte)

type Router struct {
	handlers map[string]HandlerFunc
}

func NewRouter() *Router {
	return &Router{
		handlers: make(map[string]HandlerFunc),
	}
}

func (r *Router) Register(vendor, msgType string, handler HandlerFunc) {
	key := vendor + "/" + msgType
	r.handlers[key] = handler
}

func (r *Router) Dispatch(topic string, payload []byte) {
	parts := strings.Split(topic, "/")
	if len(parts) < 5 {
		log.Printf("[Router Warning] Invalid topic format: %s", topic)
		return
	}

	vendor := parts[2]
	serial := parts[3]
	msgType := parts[4]

	key := vendor + "/" + msgType
	handler, exists := r.handlers[key]
	if !exists {
		log.Printf("[Router Warning] No handler registered for key: %s (Topic: %s)", key, topic)
		return
	}

	handler(vendor, serial, msgType, payload)
}
