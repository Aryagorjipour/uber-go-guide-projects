package metrics

import (
	"fmt"
	"sync"
)

// Registry is a thread-safe collection of metrics.
// The zero value is ready to use.
type Registry struct {
	mu      sync.RWMutex
	metrics map[string]Metric
}

// NewRegistry creates a new metrics registry with the specified initial capacity.
// If capacity is 0, a default capacity is used.
func NewRegistry(capacity int) *Registry {
	if capacity <= 0 {
		capacity = 16 // default capacity hint
	}
	return &Registry{
		metrics: make(map[string]Metric, capacity),
	}
}

// Register adds a metric to the registry.
// It returns ErrDuplicateMetric if a metric with the same name already exists.
// It returns ErrInvalidMetricName if the metric name is empty.
func (r *Registry) Register(metric Metric) error {
	if metric == nil {
		return fmt.Errorf("cannot register nil metric")
	}

	name := metric.Name()
	if name == "" {
		return ErrInvalidMetricName
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	// Initialize map on first use (zero-value usability)
	if r.metrics == nil {
		r.metrics = make(map[string]Metric, 16)
	}

	if _, exists := r.metrics[name]; exists {
		return fmt.Errorf("%w: %s", ErrDuplicateMetric, name)
	}

	r.metrics[name] = metric
	return nil
}

// Unregister removes a metric from the registry by name.
// It returns ErrMetricNotFound if the metric does not exist.
func (r *Registry) Unregister(name string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.metrics == nil {
		return fmt.Errorf("%w: %s", ErrMetricNotFound, name)
	}

	if _, exists := r.metrics[name]; !exists {
		return fmt.Errorf("%w: %s", ErrMetricNotFound, name)
	}

	delete(r.metrics, name)
	return nil
}

// Get retrieves a metric by name.
// It returns the metric and true if found, or nil and false if not found.
func (r *Registry) Get(name string) (Metric, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if r.metrics == nil {
		return nil, false
	}

	metric, exists := r.metrics[name]
	return metric, exists
}

// Snapshot returns a copy of all metrics and their current values.
// This is a defensive copy to prevent external mutation of the internal state.
// The returned map is safe to modify by the caller.
func (r *Registry) Snapshot() map[string]interface{} {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if r.metrics == nil {
		return make(map[string]interface{})
	}

	// Create a defensive copy with capacity hint
	snapshot := make(map[string]interface{}, len(r.metrics))
	for name, metric := range r.metrics {
		snapshot[name] = metric.Value()
	}

	return snapshot
}

// Len returns the number of registered metrics.
func (r *Registry) Len() int {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if r.metrics == nil {
		return 0
	}

	return len(r.metrics)
}

// Clear removes all metrics from the registry.
func (r *Registry) Clear() {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.metrics != nil {
		r.metrics = make(map[string]Metric, 16)
	}
}
