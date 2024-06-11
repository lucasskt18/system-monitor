package main

import (
	"fmt"
	"log"
	"time"

	"github.com/rcrowley/go-metrics"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
)

type SystemStats struct {
	CPUUsage    float64
	MemoryUsage float64
	DiskUsage   float64
}

func collectSystemStats() (*SystemStats, error) {
	cpuPercent, err := cpu.Percent(time.Second, false)
	if err != nil {
		return nil, err
	}

	v, err := mem.VirtualMemory()
	if err != nil {
		return nil, err
	}

	d, err := disk.Usage("/")
	if err != nil {
		return nil, err
	}

	stats := &SystemStats{
		CPUUsage:    cpuPercent[0],
		MemoryUsage: v.UsedPercent,
		DiskUsage:   d.UsedPercent,
	}

	return stats, nil
}

func main() {

	cpuMetric := metrics.NewGaugeFloat64()
	memoryMetric := metrics.NewGaugeFloat64()
	diskMetric := metrics.NewGaugeFloat64()

	metrics.Register("cpu", cpuMetric)
	metrics.Register("memory", memoryMetric)
	metrics.Register("disk", diskMetric)

	go metrics.Log(metrics.DefaultRegistry, 10*time.Second, log.New(log.Writer(), "metrics: ", log.Lmicroseconds))

	for {
		stats, err := collectSystemStats()
		if err != nil {
			log.Fatalf("Error collecting system stats: %v", err)
		}

		cpuMetric.Update(stats.CPUUsage)
		memoryMetric.Update(stats.MemoryUsage)
		diskMetric.Update(stats.DiskUsage)

		fmt.Printf("CPU Usage: %.2f%%\n", stats.CPUUsage)
		fmt.Printf("Memory Usage: %.2f%%\n", stats.MemoryUsage)
		fmt.Printf("Disk Usage: %.2f%%\n", stats.DiskUsage)

		time.Sleep(10 * time.Second) 
	}
}
