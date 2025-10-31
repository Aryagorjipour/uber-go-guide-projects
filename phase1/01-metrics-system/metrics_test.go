package metrics

import (
	"errors"
	"sync"
	"testing"
)

// TestMetricType_String tests the String method of MetricType.
func TestMetricType_String(t *testing.T) {
	tests := []struct {
		name       string
		metricType MetricType
		want       string
	}{
		{
			name:       "counter type",
			metricType: TypeCounter,
			want:       "counter",
		},
		{
			name:       "gauge type",
			metricType: TypeGauge,
			want:       "gauge",
		},
		{
			name:       "unknown type",
			metricType: MetricType(999),
			want:       "unknown",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.metricType.String(); got != tt.want {
				t.Errorf("MetricType.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestCounter tests the Counter implementation.
func TestCounter(t *testing.T) {
	t.Run("zero value is usable", func(t *testing.T) {
		var c Counter
		c.name = "test"

		if got := c.Load(); got != 0 {
			t.Errorf("zero value Counter.Load() = %v, want 0", got)
		}

		c.Inc()
		if got := c.Load(); got != 1 {
			t.Errorf("after Inc(), Counter.Load() = %v, want 1", got)
		}
	})

	t.Run("NewCounter creates counter with name", func(t *testing.T) {
		c := NewCounter("test_counter")

		if got := c.Name(); got != "test_counter" {
			t.Errorf("Counter.Name() = %v, want test_counter", got)
		}

		if got := c.Type(); got != TypeCounter {
			t.Errorf("Counter.Type() = %v, want %v", got, TypeCounter)
		}

		if got := c.Load(); got != 0 {
			t.Errorf("new Counter.Load() = %v, want 0", got)
		}
	})

	t.Run("Inc increments by 1", func(t *testing.T) {
		c := NewCounter("test")
		c.Inc()
		c.Inc()
		c.Inc()

		if got := c.Load(); got != 3 {
			t.Errorf("after 3 Inc(), Counter.Load() = %v, want 3", got)
		}
	})

	t.Run("Add increments by delta", func(t *testing.T) {
		tests := []struct {
			name  string
			delta int64
			want  int64
		}{
			{
				name:  "add positive value",
				delta: 10,
				want:  10,
			},
			{
				name:  "add zero",
				delta: 0,
				want:  0,
			},
			{
				name:  "add negative (treated as 0)",
				delta: -5,
				want:  0,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				c := NewCounter("test")
				c.Add(tt.delta)

				if got := c.Load(); got != tt.want {
					t.Errorf("Counter.Add(%v) resulted in %v, want %v", tt.delta, got, tt.want)
				}
			})
		}
	})

	t.Run("Value returns interface{}", func(t *testing.T) {
		c := NewCounter("test")
		c.Add(42)

		value := c.Value()
		got, ok := value.(int64)
		if !ok {
			t.Errorf("Counter.Value() returned type %T, want int64", value)
		}
		if got != 42 {
			t.Errorf("Counter.Value() = %v, want 42", got)
		}
	})

	t.Run("implements Metric interface", func(t *testing.T) {
		var _ Metric = (*Counter)(nil)
	})
}

// TestCounter_Concurrent tests counter operations under concurrent access.
func TestCounter_Concurrent(t *testing.T) {
	c := NewCounter("concurrent_test")
	const goroutines = 100
	const increments = 1000

	var wg sync.WaitGroup
	wg.Add(goroutines)

	for i := 0; i < goroutines; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < increments; j++ {
				c.Inc()
			}
		}()
	}

	wg.Wait()

	expected := int64(goroutines * increments)
	if got := c.Load(); got != expected {
		t.Errorf("concurrent Counter.Load() = %v, want %v", got, expected)
	}
}

// TestGauge tests the Gauge implementation.
func TestGauge(t *testing.T) {
	t.Run("zero value is usable", func(t *testing.T) {
		var g Gauge
		g.name = "test"

		if got := g.Load(); got != 0.0 {
			t.Errorf("zero value Gauge.Load() = %v, want 0.0", got)
		}

		g.Inc()
		if got := g.Load(); got != 1.0 {
			t.Errorf("after Inc(), Gauge.Load() = %v, want 1.0", got)
		}
	})

	t.Run("NewGauge creates gauge with name", func(t *testing.T) {
		g := NewGauge("test_gauge")

		if got := g.Name(); got != "test_gauge" {
			t.Errorf("Gauge.Name() = %v, want test_gauge", got)
		}

		if got := g.Type(); got != TypeGauge {
			t.Errorf("Gauge.Type() = %v, want %v", got, TypeGauge)
		}

		if got := g.Load(); got != 0.0 {
			t.Errorf("new Gauge.Load() = %v, want 0.0", got)
		}
	})

	t.Run("Set sets value", func(t *testing.T) {
		g := NewGauge("test")
		g.Set(42.5)

		if got := g.Load(); got != 42.5 {
			t.Errorf("after Set(42.5), Gauge.Load() = %v, want 42.5", got)
		}

		g.Set(-10.3)
		if got := g.Load(); got != -10.3 {
			t.Errorf("after Set(-10.3), Gauge.Load() = %v, want -10.3", got)
		}
	})

	t.Run("Inc and Dec", func(t *testing.T) {
		g := NewGauge("test")
		g.Inc()
		g.Inc()

		if got := g.Load(); got != 2.0 {
			t.Errorf("after 2 Inc(), Gauge.Load() = %v, want 2.0", got)
		}

		g.Dec()
		if got := g.Load(); got != 1.0 {
			t.Errorf("after Dec(), Gauge.Load() = %v, want 1.0", got)
		}
	})

	t.Run("Add with positive and negative values", func(t *testing.T) {
		tests := []struct {
			name  string
			delta float64
			want  float64
		}{
			{
				name:  "add positive",
				delta: 10.5,
				want:  10.5,
			},
			{
				name:  "add negative",
				delta: -5.25,
				want:  -5.25,
			},
			{
				name:  "add zero",
				delta: 0,
				want:  0,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				g := NewGauge("test")
				g.Add(tt.delta)

				if got := g.Load(); got != tt.want {
					t.Errorf("Gauge.Add(%v) resulted in %v, want %v", tt.delta, got, tt.want)
				}
			})
		}
	})

	t.Run("Value returns interface{}", func(t *testing.T) {
		g := NewGauge("test")
		g.Set(3.14)

		value := g.Value()
		got, ok := value.(float64)
		if !ok {
			t.Errorf("Gauge.Value() returned type %T, want float64", value)
		}
		if got != 3.14 {
			t.Errorf("Gauge.Value() = %v, want 3.14", got)
		}
	})

	t.Run("implements Metric interface", func(t *testing.T) {
		var _ Metric = (*Gauge)(nil)
	})
}

// TestGauge_Concurrent tests gauge operations under concurrent access.
func TestGauge_Concurrent(t *testing.T) {
	g := NewGauge("concurrent_test")
	const goroutines = 50

	var wg sync.WaitGroup
	wg.Add(goroutines * 2)

	// Half increment, half decrement
	for i := 0; i < goroutines; i++ {
		go func() {
			defer wg.Done()
			g.Inc()
		}()
		go func() {
			defer wg.Done()
			g.Dec()
		}()
	}

	wg.Wait()

	// Should be back to 0
	if got := g.Load(); got != 0.0 {
		t.Errorf("concurrent Gauge.Load() = %v, want 0.0", got)
	}
}

// TestRegistry tests the Registry implementation.
func TestRegistry(t *testing.T) {
	t.Run("zero value is usable", func(t *testing.T) {
		var r Registry

		c := NewCounter("test")
		if err := r.Register(c); err != nil {
			t.Errorf("zero value Registry.Register() failed: %v", err)
		}

		if got := r.Len(); got != 1 {
			t.Errorf("zero value Registry.Len() = %v, want 1", got)
		}
	})

	t.Run("NewRegistry creates registry with capacity", func(t *testing.T) {
		r := NewRegistry(10)

		if got := r.Len(); got != 0 {
			t.Errorf("new Registry.Len() = %v, want 0", got)
		}
	})

	t.Run("Register adds metric", func(t *testing.T) {
		r := NewRegistry(0)
		c := NewCounter("test_counter")

		if err := r.Register(c); err != nil {
			t.Errorf("Registry.Register() error = %v, want nil", err)
		}

		if got := r.Len(); got != 1 {
			t.Errorf("Registry.Len() = %v, want 1", got)
		}
	})

	t.Run("Register rejects duplicate metric", func(t *testing.T) {
		r := NewRegistry(0)
		c1 := NewCounter("duplicate")
		c2 := NewCounter("duplicate")

		if err := r.Register(c1); err != nil {
			t.Fatalf("first Register() failed: %v", err)
		}

		err := r.Register(c2)
		if err == nil {
			t.Error("Register() with duplicate name should return error")
		}

		if !errors.Is(err, ErrDuplicateMetric) {
			t.Errorf("Register() error = %v, want ErrDuplicateMetric", err)
		}
	})

	t.Run("Register rejects nil metric", func(t *testing.T) {
		r := NewRegistry(0)

		err := r.Register(nil)
		if err == nil {
			t.Error("Register(nil) should return error")
		}
	})

	t.Run("Register rejects empty name", func(t *testing.T) {
		r := NewRegistry(0)
		c := NewCounter("")

		err := r.Register(c)
		if err == nil {
			t.Error("Register() with empty name should return error")
		}

		if !errors.Is(err, ErrInvalidMetricName) {
			t.Errorf("Register() error = %v, want ErrInvalidMetricName", err)
		}
	})

	t.Run("Get retrieves metric", func(t *testing.T) {
		r := NewRegistry(0)
		c := NewCounter("test")
		c.Add(42)

		if err := r.Register(c); err != nil {
			t.Fatalf("Register() failed: %v", err)
		}

		metric, ok := r.Get("test")
		if !ok {
			t.Error("Get() returned false, want true")
		}

		counter, ok := metric.(*Counter)
		if !ok {
			t.Fatalf("Get() returned type %T, want *Counter", metric)
		}

		if got := counter.Load(); got != 42 {
			t.Errorf("retrieved Counter.Load() = %v, want 42", got)
		}
	})

	t.Run("Get returns false for non-existent metric", func(t *testing.T) {
		r := NewRegistry(0)

		_, ok := r.Get("nonexistent")
		if ok {
			t.Error("Get() for non-existent metric returned true, want false")
		}
	})

	t.Run("Unregister removes metric", func(t *testing.T) {
		r := NewRegistry(0)
		c := NewCounter("test")

		if err := r.Register(c); err != nil {
			t.Fatalf("Register() failed: %v", err)
		}

		if err := r.Unregister("test"); err != nil {
			t.Errorf("Unregister() error = %v, want nil", err)
		}

		if got := r.Len(); got != 0 {
			t.Errorf("after Unregister(), Registry.Len() = %v, want 0", got)
		}
	})

	t.Run("Unregister returns error for non-existent metric", func(t *testing.T) {
		r := NewRegistry(0)

		err := r.Unregister("nonexistent")
		if err == nil {
			t.Error("Unregister() for non-existent metric should return error")
		}

		if !errors.Is(err, ErrMetricNotFound) {
			t.Errorf("Unregister() error = %v, want ErrMetricNotFound", err)
		}
	})

	t.Run("Snapshot returns defensive copy", func(t *testing.T) {
		r := NewRegistry(0)
		c := NewCounter("counter")
		g := NewGauge("gauge")
		c.Add(10)
		g.Set(3.14)

		if err := r.Register(c); err != nil {
			t.Fatalf("Register(counter) failed: %v", err)
		}
		if err := r.Register(g); err != nil {
			t.Fatalf("Register(gauge) failed: %v", err)
		}

		snapshot := r.Snapshot()

		// Verify snapshot contains correct values
		if len(snapshot) != 2 {
			t.Errorf("Snapshot() length = %v, want 2", len(snapshot))
		}

		counterVal, ok := snapshot["counter"].(int64)
		if !ok {
			t.Errorf("snapshot['counter'] type = %T, want int64", snapshot["counter"])
		}
		if counterVal != 10 {
			t.Errorf("snapshot['counter'] = %v, want 10", counterVal)
		}

		gaugeVal, ok := snapshot["gauge"].(float64)
		if !ok {
			t.Errorf("snapshot['gauge'] type = %T, want float64", snapshot["gauge"])
		}
		if gaugeVal != 3.14 {
			t.Errorf("snapshot['gauge'] = %v, want 3.14", gaugeVal)
		}

		// Modify snapshot and verify original is unchanged
		snapshot["counter"] = int64(999)
		snapshot["new_key"] = "should not affect registry"

		// Get new snapshot to verify independence
		newSnapshot := r.Snapshot()
		if len(newSnapshot) != 2 {
			t.Errorf("after modifying first snapshot, new Snapshot() length = %v, want 2", len(newSnapshot))
		}
	})

	t.Run("Snapshot returns empty map for zero value registry", func(t *testing.T) {
		var r Registry

		snapshot := r.Snapshot()
		if snapshot == nil {
			t.Error("Snapshot() returned nil, want empty map")
		}
		if len(snapshot) != 0 {
			t.Errorf("Snapshot() length = %v, want 0", len(snapshot))
		}
	})

	t.Run("Clear removes all metrics", func(t *testing.T) {
		r := NewRegistry(0)
		c1 := NewCounter("counter1")
		c2 := NewCounter("counter2")

		if err := r.Register(c1); err != nil {
			t.Fatalf("Register(c1) failed: %v", err)
		}
		if err := r.Register(c2); err != nil {
			t.Fatalf("Register(c2) failed: %v", err)
		}

		r.Clear()

		if got := r.Len(); got != 0 {
			t.Errorf("after Clear(), Registry.Len() = %v, want 0", got)
		}
	})
}

// TestRegistry_Concurrent tests registry operations under concurrent access.
func TestRegistry_Concurrent(t *testing.T) {
	r := NewRegistry(100)
	const goroutines = 50

	var wg sync.WaitGroup
	wg.Add(goroutines * 3)

	// Concurrent registrations
	for i := 0; i < goroutines; i++ {
		i := i // capture loop variable
		go func() {
			defer wg.Done()
			c := NewCounter("counter_" + string(rune(i)))
			_ = r.Register(c)
		}()
	}

	// Concurrent reads
	for i := 0; i < goroutines; i++ {
		go func() {
			defer wg.Done()
			_ = r.Snapshot()
		}()
	}

	// Concurrent gets
	for i := 0; i < goroutines; i++ {
		i := i // capture loop variable
		go func() {
			defer wg.Done()
			_, _ = r.Get("counter_" + string(rune(i)))
		}()
	}

	wg.Wait()

	// Verify registry is in a consistent state
	if got := r.Len(); got <= 0 {
		t.Errorf("after concurrent operations, Registry.Len() = %v, want > 0", got)
	}
}

// TestRegistry_BoundaryProtection tests that snapshot provides boundary protection.
func TestRegistry_BoundaryProtection(t *testing.T) {
	r := NewRegistry(0)
	c := NewCounter("test")
	c.Add(100)

	if err := r.Register(c); err != nil {
		t.Fatalf("Register() failed: %v", err)
	}

	// Get first snapshot
	snapshot1 := r.Snapshot()
	val1 := snapshot1["test"].(int64)

	// Modify the original metric
	c.Add(50)

	// Get second snapshot
	snapshot2 := r.Snapshot()
	val2 := snapshot2["test"].(int64)

	// First snapshot should not be affected
	if val1 != 100 {
		t.Errorf("first snapshot value = %v, want 100", val1)
	}

	// Second snapshot should have new value
	if val2 != 150 {
		t.Errorf("second snapshot value = %v, want 150", val2)
	}

	// Verify snapshots are independent
	if val1 == val2 {
		t.Error("snapshots should be independent, but have same value")
	}
}
