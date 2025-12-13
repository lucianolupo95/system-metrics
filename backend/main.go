package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	amqp "github.com/rabbitmq/amqp091-go"
)

type MetricsMessage struct {
	UptimeSeconds string  `json:"uptime_seconds"`
	MemoryUsedKB  int64   `json:"memory_used_kb"`
	CPUUsagePct   float64 `json:"cpu_usage_pct"`
	Timestamp     string  `json:"timestamp"`
}

var lastMetrics *MetricsMessage

func main() {
	opts := mqtt.NewClientOptions().
		AddBroker("tcp://localhost:1883").
		SetClientID("system-metrics-backend")

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	queue, err := ch.QueueDeclare(
		"metrics_internal",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		panic(err)
	}

	client.Subscribe("system/metrics", 0, func(c mqtt.Client, m mqtt.Message) {
		var msg MetricsMessage

		if err := json.Unmarshal(m.Payload(), &msg); err != nil {
			fmt.Println("Invalid message:", err)
			return
		}

		lastMetrics = &msg

		fmt.Printf(
			"Uptime: %s | RAM: %d KB | CPU: %.2f%%\n",
			msg.UptimeSeconds,
			msg.MemoryUsedKB,
			msg.CPUUsagePct,
		)
		ch.Publish(
			"",
			queue.Name,
			false,
			false,
			amqp.Publishing{
				ContentType: "application/json",
				Body:        m.Payload(),
			},
		)

	})

	fmt.Println("Subscribed to MQTT topic system/metrics")

	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		// --- CORS ---
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		// ------------

		if lastMetrics == nil {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(lastMetrics)
	})

	go func() {
		fmt.Println("HTTP server listening on :8080")
		http.ListenAndServe(":8080", nil)
	}()

	select {}
}
