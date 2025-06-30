// Package exception provides a structured and standardized approach to error handling
// within the application. This file defines a specific exception type for
// indicating invalid arguments, leveraging the core exception handling
// mechanisms.
package exception

import (
	// status "github.com/osirisgate/golang-core/enum" is expected to provide
	// the `status.BadRequest` constant for setting the default status code.
	status "github.com/osirisgate/golang-core/enum"
)

// InvalidArgument is a specific exception type that signifies that one or more
// arguments provided to a function or method were invalid. This typically
// indicates a client-side error where the input does not meet expected criteria.
// It embeds `CoreException` to inherit all its properties and methods,
// ensuring consistent error reporting and formatting.
type InvalidArgument struct {
	CoreException // Embeds CoreException to inherit its fields and methods.
}

// NewInvalidArgument creates and returns a new `InvalidArgument` exception.
// It initializes the embedded `CoreException` with the provided error details
// and sets the default status code to `status.BadRequest`. This status code
// is highly appropriate for errors arising from malformed or incorrect input.
//
// Parameters:
//
//	errors: A map of string to interface{} containing detailed error information
//	        about which arguments were invalid and why. This map can include a
//	        "message" key which will be used as the primary error message for
//	        the exception.
//
// Returns:
//
//	A pointer to a new `InvalidArgument` instance.
func NewInvalidArgument(errors map[string]interface{}) *InvalidArgument {
	// Initialize the base CoreException with the given errors and a default
	// status of BadRequest, as invalid arguments are typically client-side input errors.
	base := NewInstance(errors, status.BadRequest)
	return &InvalidArgument{CoreException: *base}
}
