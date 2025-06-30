// Package exception provides a structured and standardized approach to error handling
// within the application. This file defines a specific exception type for
// arithmetic or data underflow errors, leveraging the core exception handling
// mechanisms.
package exception

import (
	// status "github.com/osirisgate/golang-core/enum" is expected to provide
	// the `status.InternalServerError` constant for setting the default status code.
	status "github.com/osirisgate/golang-core/enum"
)

// Underflow is a specific exception type that indicates an error condition
// where a numerical operation or data processing results in a value that is
// too small to be represented in the target data type, or falls below a minimum
// acceptable threshold. This can occur, for example, with floating-point calculations
// that produce a value closer to zero than the system can store.
// It embeds `CoreException` to inherit all its properties and methods,
// ensuring consistent error reporting and formatting.
type Underflow struct {
	CoreException // Embeds CoreException to inherit its fields and methods.
}

// NewUnderflow creates and returns a new `Underflow` exception.
// It initializes the embedded `CoreException` with the provided error details
// and sets the default status code to `status.InternalServerError`. This
// status code is appropriate when the underflow is an unexpected internal
// processing error, rather than a direct result of invalid client input.
//
// Parameters:
//
//	errors: A map of string to interface{} containing detailed error information
//	        about the underflow condition. This map can include a "message" key
//	        which will be used as the primary error message for the exception.
//
// Returns:
//
//	A pointer to a new `Underflow` instance.
func NewUnderflow(errors map[string]interface{}) *Underflow {
	// Initialize the base CoreException with the given errors and a default
	// status of InternalServerError, as underflow issues typically represent
	// internal computational errors.
	base := NewInstance(errors, status.InternalServerError)
	return &Underflow{CoreException: *base}
}
