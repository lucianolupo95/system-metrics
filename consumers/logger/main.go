package main

import (
	"encoding/json"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type MetricsMessage struct {
	UptimeSeconds string  `json:"uptime_seconds"`
	MemoryUsedKB  int64   `json:"memory_used_kb"`
	CPUUsagePct   float64 `json:"cpu_usage_pct"`
	Timestamp     string  `json:"timestamp"`
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
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
		log.Fatal(err)
	}

	msgs, err := ch.Consume(
		queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("RabbitMQ logger consumer started")

	for msg := range msgs {
		var metrics MetricsMessage
		if err := json.Unmarshal(msg.Body, &metrics); err != nil {
			log.Println("Invalid message:", err)
			continue
		}

		log.Printf(
			"[LOGGER] CPU: %.2f%% | RAM: %d KB | Uptime: %s\n",
			metrics.CPUUsagePct,
			metrics.MemoryUsedKB,
			metrics.UptimeSeconds,
		)
	}
}
