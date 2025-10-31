// Package metrics provides a thread-safe metrics collection system with
// support for counters and gauges using atomic operations.
package metrics

// Metric represents a metric that can be collected and reported.
// All metric implementations must be safe for concurrent use by multiple goroutines.
type Metric interface {
	// Name returns the unique identifier for this metric.
	Name() string

	// Type returns the type of this metric (counter, gauge, etc.).
	Type() MetricType

	// Value returns the current value of this metric.
	// The concrete type returned depends on the metric type.
	Value() interface{}
}

// MetricType represents the type of a metric.
type MetricType int

const (
	// TypeCounter represents a monotonically increasing counter metric.
	// Counters start at 1 to avoid zero-value ambiguity.
	TypeCounter MetricType = iota + 1

	// TypeGauge represents a gauge metric that can increase or decrease.
	TypeGauge
)

// String returns a human-readable string representation of the metric type.
func (t MetricType) String() string {
	switch t {
	case TypeCounter:
		return "counter"
	case TypeGauge:
		return "gauge"
	default:
		return "unknown"
	}
}
