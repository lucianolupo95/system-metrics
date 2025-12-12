package main

import (
	"encoding/json"
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MetricsMessage struct {
	UptimeSeconds string `json:"uptime_seconds"`
	Timestamp     string `json:"timestamp"`
}

func main() {
	opts := mqtt.NewClientOptions().
		AddBroker("tcp://localhost:1883").
		SetClientID("system-metrics-backend")

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	client.Subscribe("system/metrics", 0, func(c mqtt.Client, m mqtt.Message) {
		var msg MetricsMessage

		err := json.Unmarshal(m.Payload(), &msg)
		if err != nil {
			fmt.Println("Invalid message:", err)
			return
		}

		fmt.Printf(
			"Uptime: %s seconds | Timestamp: %s\n",
			msg.UptimeSeconds,
			msg.Timestamp,
		)
	})

	fmt.Println("Backend listening on system/metrics")
	select {}
}
