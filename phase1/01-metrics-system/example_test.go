package metrics_test

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/Aryagorjipour/uber-go-guide-projects/phase1/01-metrics-system"
)

// ExampleCounter demonstrates basic counter usage.
func ExampleCounter() {
	counter := metrics.NewCounter("requests")

	// Increment the counter
	counter.Inc()
	counter.Inc()

	// Add multiple at once
	counter.Add(5)

	fmt.Println(counter.Load())
	// Output: 7
}

// ExampleGauge demonstrates basic gauge usage.
func ExampleGauge() {
	gauge := metrics.NewGauge("temperature")

	// Set to a specific value
	gauge.Set(20.5)

	// Increment and decrement
	gauge.Inc() // 21.5
	gauge.Dec() // 20.5

	fmt.Println(gauge.Load())
	// Output: 20.5
}

// ExampleRegistry demonstrates registry usage.
func ExampleRegistry() {
	// Create a registry with capacity hint
	registry := metrics.NewRegistry(10)

	// Create and register metrics
	requests := metrics.NewCounter("http_requests_total")
	errors := metrics.NewCounter("http_errors_total")
	temperature := metrics.NewGauge("cpu_temperature")

	registry.Register(requests)
	registry.Register(errors)
	registry.Register(temperature)

	// Update metrics
	requests.Add(100)
	errors.Add(5)
	temperature.Set(65.3)

	// Get a snapshot
	snapshot := registry.Snapshot()

	fmt.Printf("Total requests: %v\n", snapshot["http_requests_total"])
	fmt.Printf("Total errors: %v\n", snapshot["http_errors_total"])
	fmt.Printf("CPU temp: %v\n", snapshot["cpu_temperature"])
	// Output:
	// Total requests: 100
	// Total errors: 5
	// CPU temp: 65.3
}

// ExampleRegistry_concurrent demonstrates thread-safe concurrent usage.
func ExampleRegistry_concurrent() {
	registry := metrics.NewRegistry(10)
	counter := metrics.NewCounter("concurrent_counter")
	registry.Register(counter)

	// Simulate concurrent updates from multiple goroutines
	var wg sync.WaitGroup
	workers := 10
	increments := 100

	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < increments; j++ {
				counter.Inc()
			}
		}()
	}

	wg.Wait()

	snapshot := registry.Snapshot()
	fmt.Println(snapshot["concurrent_counter"])
	// Output: 1000
}

// ExampleCounter_zeroValue demonstrates zero-value usability.
func ExampleCounter_zeroValue() {
	// Counter can be used without explicit initialization
	var counter metrics.Counter

	counter.Inc()
	counter.Inc()

	fmt.Println(counter.Load())
	// Output: 2
}

// ExampleRegistry_boundaryProtection demonstrates defensive copying.
func ExampleRegistry_boundaryProtection() {
	registry := metrics.NewRegistry(10)
	counter := metrics.NewCounter("test")
	counter.Add(100)
	registry.Register(counter)

	// Get first snapshot
	snapshot1 := registry.Snapshot()
	fmt.Printf("Snapshot 1: %v\n", snapshot1["test"])

	// Modify the original counter
	counter.Add(50)

	// Get second snapshot
	snapshot2 := registry.Snapshot()
	fmt.Printf("Snapshot 2: %v\n", snapshot2["test"])

	// First snapshot is unchanged (boundary protection)
	fmt.Printf("Snapshot 1 (unchanged): %v\n", snapshot1["test"])
	// Output:
	// Snapshot 1: 100
	// Snapshot 2: 150
	// Snapshot 1 (unchanged): 100
}

// ExampleRegistry_errorHandling demonstrates error handling.
func ExampleRegistry_errorHandling() {
	registry := metrics.NewRegistry(10)
	counter := metrics.NewCounter("my_counter")

	// First registration succeeds
	if err := registry.Register(counter); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Registered successfully")
	}

	// Duplicate registration fails
	duplicate := metrics.NewCounter("my_counter")
	if err := registry.Register(duplicate); err != nil {
		fmt.Println("Error: duplicate metric")
	}

	// Output:
	// Registered successfully
	// Error: duplicate metric
}

// ExampleGauge_monitoring demonstrates a realistic monitoring scenario.
func ExampleGauge_monitoring() {
	registry := metrics.NewRegistry(10)

	// Create metrics for a monitoring system
	cpuUsage := metrics.NewGauge("cpu_usage_percent")
	memoryUsage := metrics.NewGauge("memory_usage_mb")
	activeConnections := metrics.NewGauge("active_connections")

	registry.Register(cpuUsage)
	registry.Register(memoryUsage)
	registry.Register(activeConnections)

	// Simulate monitoring updates
	cpuUsage.Set(45.2)
	memoryUsage.Set(2048.5)
	activeConnections.Set(150)

	// Simulate changes
	cpuUsage.Add(5.3)       // CPU increased
	activeConnections.Dec() // One connection closed
	activeConnections.Dec() // Another connection closed

	snapshot := registry.Snapshot()
	fmt.Printf("CPU: %.1f%%\n", snapshot["cpu_usage_percent"])
	fmt.Printf("Memory: %.1f MB\n", snapshot["memory_usage_mb"])
	fmt.Printf("Connections: %.0f\n", snapshot["active_connections"])
	// Output:
	// CPU: 50.5%
	// Memory: 2048.5 MB
	// Connections: 148
}

// ExampleCounter_webServer demonstrates HTTP request counting.
func ExampleCounter_webServer() {
	registry := metrics.NewRegistry(10)

	// Create metrics for different HTTP status codes
	requests2xx := metrics.NewCounter("http_2xx_total")
	requests4xx := metrics.NewCounter("http_4xx_total")
	requests5xx := metrics.NewCounter("http_5xx_total")

	registry.Register(requests2xx)
	registry.Register(requests4xx)
	registry.Register(requests5xx)

	// Simulate handling requests
	requests2xx.Add(1500) // Successful requests
	requests4xx.Add(50)   // Client errors
	requests5xx.Add(5)    // Server errors

	// Print metrics
	snapshot := registry.Snapshot()
	total := snapshot["http_2xx_total"].(int64) +
		snapshot["http_4xx_total"].(int64) +
		snapshot["http_5xx_total"].(int64)

	fmt.Printf("Total requests: %d\n", total)
	fmt.Printf("Success rate: %.1f%%\n",
		float64(snapshot["http_2xx_total"].(int64))/float64(total)*100)
	// Output:
	// Total requests: 1555
	// Success rate: 96.5%
}

// ExampleRegistry_snapshot demonstrates snapshot functionality.
func ExampleRegistry_Snapshot() {
	registry := metrics.NewRegistry(5)

	// Register multiple metrics
	counter1 := metrics.NewCounter("counter1")
	counter2 := metrics.NewCounter("counter2")
	gauge1 := metrics.NewGauge("gauge1")

	counter1.Add(10)
	counter2.Add(20)
	gauge1.Set(30.5)

	registry.Register(counter1)
	registry.Register(counter2)
	registry.Register(gauge1)

	// Get snapshot - returns a defensive copy
	snapshot := registry.Snapshot()

	// Snapshot can be safely modified
	snapshot["counter1"] = int64(999)
	snapshot["new_key"] = "won't affect registry"

	// Original registry is unchanged
	if metric, found := registry.Get("counter1"); found {
		fmt.Println(metric.Value()) // Still 10
	}

	fmt.Println(len(registry.Snapshot())) // Still 3 metrics
	// Output:
	// 10
	// 3
}

// Benchmark examples

// BenchmarkCounter_Inc benchmarks counter increment operations.
func BenchmarkCounter_Inc(b *testing.B) {
	counter := metrics.NewCounter("bench")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		counter.Inc()
	}
}

// BenchmarkCounter_Concurrent benchmarks concurrent counter operations.
func BenchmarkCounter_Concurrent(b *testing.B) {
	counter := metrics.NewCounter("bench")
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			counter.Inc()
		}
	})
}

// BenchmarkGauge_Set benchmarks gauge set operations.
func BenchmarkGauge_Set(b *testing.B) {
	gauge := metrics.NewGauge("bench")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gauge.Set(float64(i))
	}
}

// BenchmarkRegistry_Register benchmarks metric registration.
func BenchmarkRegistry_Register(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		registry := metrics.NewRegistry(100)
		counter := metrics.NewCounter(fmt.Sprintf("counter_%d", i))
		b.StartTimer()

		registry.Register(counter)
	}
}

// BenchmarkRegistry_Snapshot benchmarks snapshot operations.
func BenchmarkRegistry_Snapshot(b *testing.B) {
	registry := metrics.NewRegistry(100)
	for i := 0; i < 50; i++ {
		counter := metrics.NewCounter(fmt.Sprintf("counter_%d", i))
		registry.Register(counter)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = registry.Snapshot()
	}
}

// Example demonstrating real-world usage pattern
func Example_realWorld() {
	// Create a global registry (typically initialized at startup)
	registry := metrics.NewRegistry(20)

	// Create application metrics
	httpRequests := metrics.NewCounter("http_requests_total")
	httpDuration := metrics.NewGauge("http_request_duration_ms")
	activeUsers := metrics.NewGauge("active_users")

	// Register all metrics
	registry.Register(httpRequests)
	registry.Register(httpDuration)
	registry.Register(activeUsers)

	// Simulate application activity
	start := time.Now()

	// Process a request
	httpRequests.Inc()
	activeUsers.Inc()

	// Simulate request processing
	time.Sleep(1 * time.Millisecond)

	// Record duration
	duration := time.Since(start).Milliseconds()
	httpDuration.Set(float64(duration))

	// Request complete
	activeUsers.Dec()

	// Export metrics (e.g., for Prometheus scraping)
	snapshot := registry.Snapshot()
	fmt.Printf("Requests: %v\n", snapshot["http_requests_total"])
	fmt.Printf("Duration: %.0fms\n", snapshot["http_request_duration_ms"])
	fmt.Printf("Active users: %.0f\n", snapshot["active_users"])
	// Output:
	// Requests: 1
	// Duration: 1ms
	// Active users: 0
}
