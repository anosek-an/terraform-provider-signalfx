/*
 * Metrics Metadata API
 *
 * API for creating, retrieving, updating, and deleting metric names and MTS metadata.<br> **NOTE:*() Although you can't set custom properties or tags for a metric, you *can* retrieve them for metrics and metric time series (**MTS**).
 *
 * API version: 3.2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package metrics_metadata

// The results of the metrics metadata request
type RetrieveMetricMetadataResponseModel struct {
	// An array of metrics metadata results
	Result []*Metric `json:"result,omitempty"`
	// Number of result objects returned. This value is the same as the size of the `result` value array.
	Count int32 `json:"count,omitempty"`
}
