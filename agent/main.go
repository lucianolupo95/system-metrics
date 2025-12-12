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
	UptimeSeconds string  `json:"uptime_seconds"`
	MemoryUsedKB  int64   `json:"memory_used_kb"`
	CPUUsagePct   float64 `json:"cpu_usage_pct"`
	Timestamp     string  `json:"timestamp"`
}

func main() {
	data, err := os.ReadFile("/proc/uptime")
	if err != nil {
		panic(err)
	}

	fields := strings.Fields(string(data))

	msg := MetricsMessage{
		UptimeSeconds: fields[0],
		MemoryUsedKB:  readMemoryUsedKB(),
		CPUUsagePct:   readCPUUsagePercent(),
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
func readMemoryUsedKB() int64 {
	data, err := os.ReadFile("/proc/meminfo")
	if err != nil {
		panic(err)
	}

	var memTotal, memAvailable int64

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "MemTotal:") {
			fmt.Sscanf(line, "MemTotal: %d kB", &memTotal)
		}
		if strings.HasPrefix(line, "MemAvailable:") {
			fmt.Sscanf(line, "MemAvailable: %d kB", &memAvailable)
		}
	}

	return memTotal - memAvailable
}
func readCPUStat() (idle, total uint64) {
	data, err := os.ReadFile("/proc/stat")
	if err != nil {
		panic(err)
	}

	var cpu string
	var user, nice, system, idleTime, iowait, irq, softirq, steal uint64

	fmt.Sscanf(
		string(data),
		"%s %d %d %d %d %d %d %d %d",
		&cpu,
		&user,
		&nice,
		&system,
		&idleTime,
		&iowait,
		&irq,
		&softirq,
		&steal,
	)

	idle = idleTime + iowait
	total = user + nice + system + idleTime + iowait + irq + softirq + steal
	return
}
func readCPUUsagePercent() float64 {
	idle1, total1 := readCPUStat()
	time.Sleep(500 * time.Millisecond)
	idle2, total2 := readCPUStat()

	idleDelta := float64(idle2 - idle1)
	totalDelta := float64(total2 - total1)
	if totalDelta == 0 {
		return 0
	}

	return (1.0 - idleDelta/totalDelta) * 100.0
}
