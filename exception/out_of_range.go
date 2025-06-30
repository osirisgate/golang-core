// Package exception provides a structured and standardized approach to error handling
// within the application. This file defines a specific exception type for
// errors related to values being out of an expected range, leveraging the core
// exception handling mechanisms.
package exception

import (
	// status "github.com/osirisgate/golang-core/enum" is expected to provide
	// the `status.BadRequest` constant for setting the default status code.
	status "github.com/osirisgate/golang-core/enum"
)

// OutOfRange is a specific exception type that signifies an error where
// a numerical or categorical value falls outside of its permissible range.
// This typically indicates an input validation error where the provided value
// is either too low, too high, or not within an allowed set.
// It embeds `CoreException` to inherit all its properties and methods,
// ensuring consistent error reporting and formatting.
type OutOfRange struct {
	CoreException // Embeds CoreException to inherit its fields and methods.
}

// NewOutOfRange creates and returns a new `OutOfRange` exception.
// It initializes the embedded `CoreException` with the provided error details
// and sets the default status code to `status.BadRequest`. This status code
// is appropriate for indicating that the client's request failed due to
// a value that is outside the acceptable range.
//
// Parameters:
//
//	errors: A map of string to interface{} containing detailed error information
//	        about the out-of-range value. This map can include a "message" key
//	        which will be used as the primary error message for the exception.
//
// Returns:
//
//	A pointer to a new `OutOfRange` instance.
func NewOutOfRange(errors map[string]interface{}) *OutOfRange {
	// Initialize the base CoreException with the given errors and a default
	// status of BadRequest, as out-of-range errors are typically client-side input validation issues.
	base := NewInstance(errors, status.BadRequest)
	return &OutOfRange{CoreException: *base}
}
