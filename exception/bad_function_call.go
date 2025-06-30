// Package exception provides a structured and standardized approach to error handling
// within the application. This specific file defines an exception type for
// indicating bad function calls, leveraging the core exception handling
// mechanisms.
package exception

import (
	// status "github.com/osirisgate/golang-core/enum" is expected to provide
	// the `status.BadRequest` constant for setting the default status code.
	status "github.com/osirisgate/golang-core/enum"
)

// BadFunctionCall is a specific exception type that indicates an error
// related to an invalid or improperly used function call.
// It embeds `CoreException` to inherit all its properties and methods,
// ensuring consistent error reporting and formatting.
type BadFunctionCall struct {
	CoreException // Embeds CoreException to inherit its fields and methods.
}

// NewBadFunctionCall creates and returns a new `BadFunctionCall` exception.
// It initializes the embedded `CoreException` with the provided error details
// and sets the default status code to `status.BadRequest`, which is appropriate
// for client-side issues related to how a function was invoked.
//
// Parameters:
//
//	errors: A map of string to interface{} containing detailed error information.
//	        This map can include a "message" key which will be used as the
//	        primary error message for the exception.
//
// Returns:
//
//	A pointer to a new `BadFunctionCall` instance.
func NewBadFunctionCall(errors map[string]interface{}) *BadFunctionCall {
	// Initialize the base CoreException with the given errors and a default
	// status of BadRequest, as this exception typically signifies a client-side
	// issue with a function call.
	base := NewInstance(errors, status.BadRequest)
	return &BadFunctionCall{CoreException: *base}
}
