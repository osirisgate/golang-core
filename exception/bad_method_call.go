// Package exception provides a structured and standardized approach to error handling
// within the application. This file defines a specific exception type for
// indicating bad method calls, leveraging the core exception handling
// mechanisms.
package exception

import (
	// status "github.com/osirisgate/golang-core/enum" is expected to provide
	// the `status.BadRequest` constant for setting the default status code.
	status "github.com/osirisgate/golang-core/enum"
)

// BadMethodCall is a specific exception type that indicates an error
// related to an invalid or improperly used method call on an object.
// It embeds `CoreException` to inherit all its properties and methods,
// ensuring consistent error reporting and formatting.
type BadMethodCall struct {
	CoreException // Embeds CoreException to inherit its fields and methods.
}

// NewBadMethodCall creates and returns a new `BadMethodCall` exception.
// It initializes the embedded `CoreException` with the provided error details
// and sets the default status code to `status.BadRequest`, which is appropriate
// for client-side issues related to how a method was invoked.
//
// Parameters:
//
//	errors: A map of string to interface{} containing detailed error information.
//	        This map can include a "message" key which will be used as the
//	        primary error message for the exception.
//
// Returns:
//
//	A pointer to a new `BadMethodCall` instance.
func NewBadMethodCall(errors map[string]interface{}) *BadMethodCall {
	// Initialize the base CoreException with the given errors and a default
	// status of BadRequest, as this exception typically signifies a client-side
	// issue with a method call.
	base := NewInstance(errors, status.BadRequest)
	return &BadMethodCall{CoreException: *base}
}
