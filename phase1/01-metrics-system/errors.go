package metrics

import "errors"

var (
	// ErrDuplicateMetric is returned when attempting to register a metric
	// with a name that is already registered in the registry.
	ErrDuplicateMetric = errors.New("metric with this name already exists")

	// ErrMetricNotFound is returned when attempting to retrieve or unregister
	// a metric that does not exist in the registry.
	ErrMetricNotFound = errors.New("metric not found")

	// ErrInvalidMetricName is returned when attempting to register a metric
	// with an empty or invalid name.
	ErrInvalidMetricName = errors.New("metric name cannot be empty")
)
