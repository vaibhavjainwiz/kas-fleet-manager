/*
 * Kafka Service Fleet Manager
 *
 * Kafka Service Fleet Manager APIs that are used by internal services e.g kas-fleetshard operators.
 *
 * API version: 1.7.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package private

// WatchEvent struct for WatchEvent
type WatchEvent struct {
	Type   string                  `json:"type"`
	Error  Error                   `json:"error,omitempty"`
	Object *map[string]interface{} `json:"object,omitempty"`
}
