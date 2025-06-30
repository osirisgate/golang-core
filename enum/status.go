// Package status provides core constants and a type for representing
// the general status of an operation, typically used in API responses
// or internal process results.
package status

const (
	SUCCESS = "success" // SUCCESS represents a successful operation status.
	ERROR   = "error"   // ERROR represents an erroneous or failed operation status.
)

// Status is a custom string type used to define and manage
// the various states of an operation (e.g., "success", "error").
type Status string

// GetValue returns the underlying string value of the Status type.
func (s Status) GetValue() string {
	return string(s)
}

// NewStatus is a constructor function that creates and returns a new Status type
// from a given string value.
//
// Parameters:
//
//	value: The string representing the desired status (e.g., "success", "error").
//
// Returns:
//
//	A new Status type initialized with the provided string value.
func NewStatus(value string) Status {
	return Status(value)
}
