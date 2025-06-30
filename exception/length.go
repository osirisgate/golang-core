// Package exception provides a structured and standardized approach to error handling
// within the application. This file defines a specific exception type for
// errors related to invalid lengths, leveraging the core exception handling
// mechanisms.
package exception

import (
	// status "github.com/osirisgate/golang-core/enum" is expected to provide
	// the `status.BadRequest` constant for setting the default status code.
	status "github.com/osirisgate/golang-core/enum"
)

// Length is a specific exception type that signifies an error where
// a provided length (e.g., of a string, array, or file size) does not
// meet the required criteria (e.g., too short, too long, incorrect exact length).
// It typically indicates a client-side input validation error.
// It embeds `CoreException` to inherit all its properties and methods,
// ensuring consistent error reporting and formatting.
type Length struct {
	CoreException // Embeds CoreException to inherit its fields and methods.
}

// NewLength creates and returns a new `Length` exception.
// It initializes the embedded `CoreException` with the provided error details
// and sets the default status code to `status.BadRequest`. This status code
// is appropriate for indicating that the client's request failed due to
// an invalid length specification.
//
// Parameters:
//
//	errors: A map of string to interface{} containing detailed error information
//	        about the length constraint violation. This map can include a
//	        "message" key which will be used as the primary error message for
//	        the exception.
//
// Returns:
//
//	A pointer to a new `Length` instance.
func NewLength(errors map[string]interface{}) *Length {
	// Initialize the base CoreException with the given errors and a default
	// status of BadRequest, as length errors are typically client-side input validation issues.
	base := NewInstance(errors, status.BadRequest)
	return &Length{CoreException: *base}
}
