package main

import (
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func main() {
	opts := mqtt.NewClientOptions().
		AddBroker("tcp://localhost:1883").
		SetClientID("system-metrics-backend")

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	client.Subscribe("system/metrics", 0, func(c mqtt.Client, m mqtt.Message) {
		fmt.Println("Received:", string(m.Payload()))
	})

	fmt.Println("Backend listening on system/metrics")
	select {}
}
