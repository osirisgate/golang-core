// Package exception provides a structured and standardized approach to error handling
// within the application. This file defines a specific exception type for
// logic-related errors, leveraging the core exception handling mechanisms.
package exception

import (
	// status "github.com/osirisgate/golang-core/enum" is expected to provide
	// the `status.BadRequest` constant for setting the default status code.
	status "github.com/osirisgate/golang-core/enum"
)

// Logic is a specific exception type that signifies an error resulting from
// a logical flaw or inconsistency in the application's business rules or
// internal processing. While often triggered by invalid client input,
// it specifically points to an issue in the application's expected logical flow.
// It embeds `CoreException` to inherit all its properties and methods,
// ensuring consistent error reporting and formatting.
type Logic struct {
	CoreException // Embeds CoreException to inherit its fields and methods.
}

// NewLogic creates and returns a new `Logic` exception.
// It initializes the embedded `CoreException` with the provided error details
// and sets the default status code to `status.BadRequest`. This status code
// is generally appropriate for logic errors that are ultimately caused by
// client-provided data that violates business logic.
//
// Parameters:
//
//	errors: A map of string to interface{} containing detailed error information
//	        about the logical inconsistency. This map can include a "message" key
//	        which will be used as the primary error message for the exception.
//
// Returns:
//
//	A pointer to a new `Logic` instance.
func NewLogic(errors map[string]interface{}) *Logic {
	// Initialize the base CoreException with the given errors and a default
	// status of BadRequest, as logic errors often manifest due to invalid input
	// that breaches business rules.
	base := NewInstance(errors, status.BadRequest)
	return &Logic{CoreException: *base}
}
