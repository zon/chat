package main

import (
	"encoding/json"

	"github.com/nats-io/nats.go"
)

var nc *nats.Conn 

func Connect() error {
	if nc != nil && nc.IsConnected() {
		return nil
	}
	var err error
	nc, err = nats.Connect(config.Nats.String())
	return err
}

func Publish(subject string, value any) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return nc.Publish(subject, data)
}
