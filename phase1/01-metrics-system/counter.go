package metrics

import "go.uber.org/atomic"

// Counter is a monotonically increasing counter metric that is safe for
// concurrent use by multiple goroutines. The zero value is ready to use.
type Counter struct {
	name  string
	value atomic.Int64
}

// Compile-time verification that Counter implements Metric interface.
var _ Metric = (*Counter)(nil)

// NewCounter creates a new counter metric with the given name.
// The counter starts at 0 and can only be incremented.
func NewCounter(name string) *Counter {
	return &Counter{
		name: name,
	}
}

// Name returns the name of this counter metric.
func (c *Counter) Name() string {
	return c.name
}

// Type returns TypeCounter, indicating this is a counter metric.
func (c *Counter) Type() MetricType {
	return TypeCounter
}

// Value returns the current value of the counter as an interface{}.
// The underlying type is int64.
func (c *Counter) Value() interface{} {
	return c.value.Load()
}

// Inc increments the counter by 1.
// This operation is atomic and safe for concurrent use.
func (c *Counter) Inc() {
	c.value.Add(1)
}

// Add increments the counter by the given delta.
// Delta must be non-negative. Negative values are treated as 0.
// This operation is atomic and safe for concurrent use.
func (c *Counter) Add(delta int64) {
	if delta < 0 {
		delta = 0
	}
	c.value.Add(delta)
}

// Load returns the current value of the counter.
// This is a convenience method that returns int64 directly.
func (c *Counter) Load() int64 {
	return c.value.Load()
}
