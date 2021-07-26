package main

import (
	"fmt"
	"task3/helper"
	"time"

	emitter "github.com/emitter-io/go/v2"
)

var fileName = "../config.yml"
var emitterHostPort = "tcp://localhost:8080"

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

	for i := 1; i <= config.MessagesToEmit; i++ {
		err = conn.Publish(config.SecretKey, config.Channel, config.Message)
		if err != nil {
			panic(err)
			break
		}
		fmt.Printf("publishing message %s serial number %d\n", config.Message, i)
		time.Sleep(1 * time.Second)
	}
}
