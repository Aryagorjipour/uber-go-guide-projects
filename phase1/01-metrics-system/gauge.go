package metrics

import "go.uber.org/atomic"

// Gauge is a metric that can increase or decrease and is safe for
// concurrent use by multiple goroutines. The zero value is ready to use.
type Gauge struct {
	name  string
	value atomic.Float64
}

// Compile-time verification that Gauge implements Metric interface.
var _ Metric = (*Gauge)(nil)

// NewGauge creates a new gauge metric with the given name.
// The gauge starts at 0.0 and can be set to any value.
func NewGauge(name string) *Gauge {
	return &Gauge{
		name: name,
	}
}

// Name returns the name of this gauge metric.
func (g *Gauge) Name() string {
	return g.name
}

// Type returns TypeGauge, indicating this is a gauge metric.
func (g *Gauge) Type() MetricType {
	return TypeGauge
}

// Value returns the current value of the gauge as an interface{}.
// The underlying type is float64.
func (g *Gauge) Value() interface{} {
	return g.value.Load()
}

// Set sets the gauge to the given value.
// This operation is atomic and safe for concurrent use.
func (g *Gauge) Set(value float64) {
	g.value.Store(value)
}

// Inc increments the gauge by 1.
// This operation is atomic and safe for concurrent use.
func (g *Gauge) Inc() {
	g.value.Add(1.0)
}

// Dec decrements the gauge by 1.
// This operation is atomic and safe for concurrent use.
func (g *Gauge) Dec() {
	g.value.Add(-1.0)
}

// Add adds the given delta to the gauge.
// Delta can be positive or negative.
// This operation is atomic and safe for concurrent use.
func (g *Gauge) Add(delta float64) {
	g.value.Add(delta)
}

// Load returns the current value of the gauge.
// This is a convenience method that returns float64 directly.
func (g *Gauge) Load() float64 {
	return g.value.Load()
}
