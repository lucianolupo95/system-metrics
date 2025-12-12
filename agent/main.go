package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MetricsMessage struct {
	UptimeSeconds string `json:"uptime_seconds"`
	Timestamp     string `json:"timestamp"`
}

func main() {
	data, err := os.ReadFile("/proc/uptime")
	if err != nil {
		panic(err)
	}

	fields := strings.Fields(string(data))

	msg := MetricsMessage{
		UptimeSeconds: fields[0],
		Timestamp:     time.Now().UTC().Format(time.RFC3339),
	}

	jsonData, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	opts := mqtt.NewClientOptions().
		AddBroker("tcp://localhost:1883").
		SetClientID("system-metrics-agent")

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	token := client.Publish("system/metrics", 0, false, jsonData)
	token.Wait()

	fmt.Println("Published:", string(jsonData))

	client.Disconnect(250)
}
