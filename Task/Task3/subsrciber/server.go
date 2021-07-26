package main

import (
	"fmt"
	"task3/helper"

	emitter "github.com/emitter-io/go/v2"
)

var fileName = "../config.yml"
var emitterHostPort = "tcp://localhost:8080"
var test bool

func main() {
	config, err := helper.LoadConfig(fileName)
	if err != nil {
		fmt.Println("Unable to read configs", err)
		return
	}

	conn, err := emitter.Connect(emitterHostPort, func(_ *emitter.Client, msg emitter.Message) {
		fmt.Printf("Topic = %s Payload = %s\n", msg.Topic(), msg.Payload())
	})
	if err != nil {
		fmt.Println("wrong emitter address")
		return
	}
	fmt.Println("Connected with emitter server. Waiting to get messages")
	for {
		err = conn.Subscribe(config.SecretKey, config.Channel, func(_ *emitter.Client, msg emitter.Message) {
			fmt.Printf("Topic = %s Payload = %s\n", msg.Topic(), msg.Payload())
		})
		if err != nil {
			panic(err)
		}

		if test {
			break
		}
	}
}
