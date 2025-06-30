// Package exception provides a structured and standardized approach to error handling
// within the application. This file defines a specific exception type for
// runtime errors, leveraging the core exception handling mechanisms.
package exception

import (
	// status "github.com/osirisgate/golang-core/enum" is expected to provide
	// the `status.InternalServerError` constant for setting the default status code.
	status "github.com/osirisgate/golang-core/enum"
)

// Runtime is a specific exception type that signifies an error occurring
// during the execution of the program that is not typically due to invalid
// client input or domain logic, but rather an unexpected condition or fault
// within the application's operational environment or underlying systems.
// It embeds `CoreException` to inherit all its properties and methods,
// ensuring consistent error reporting and formatting.
type Runtime struct {
	CoreException // Embeds CoreException to inherit its fields and methods.
}

// NewRuntime creates and returns a new `Runtime` exception.
// It initializes the embedded `CoreException` with the provided error details
// and sets the default status code to `status.InternalServerError`. This
// status code is appropriate for errors that are internal to the server
// and are not directly the fault of the client's request.
//
// Parameters:
//
//	errors: A map of string to interface{} containing detailed error information
//	        about the runtime issue. This map can include a "message" key
//	        which will be used as the primary error message for the exception.
//
// Returns:
//
//	A pointer to a new `Runtime` instance.
func NewRuntime(errors map[string]interface{}) *Runtime {
	// Initialize the base CoreException with the given errors and a default
	// status of InternalServerError, as runtime errors are typically server-side issues.
	base := NewInstance(errors, status.InternalServerError)
	return &Runtime{CoreException: *base}
}
