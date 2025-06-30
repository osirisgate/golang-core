// Package exception provides a structured and standardized approach to error handling
// within the application. This file defines a specific exception type for
// domain-related errors, leveraging the core exception handling mechanisms.
package exception

import (
	// status "github.com/osirisgate/golang-core/enum" is expected to provide
	// the `status.BadRequest` constant for setting the default status code.
	status "github.com/osirisgate/golang-core/enum"
)

// Domain is a specific exception type that represents an error occurring
// within the domain logic or business rules of the application.
// It embeds `CoreException` to inherit all its properties and methods,
// ensuring consistent error reporting and formatting. Domain errors are
// typically indicative of issues with input data or business constraints.
type Domain struct {
	CoreException // Embeds CoreException to inherit its fields and methods.
}

// NewDomain creates and returns a new `Domain` exception.
// It initializes the embedded `CoreException` with the provided error details
// and sets the default status code to `status.BadRequest`, which is often
// appropriate for domain-level validation failures or business rule violations
// caused by client input.
//
// Parameters:
//
//	errors: A map of string to interface{} containing detailed error information
//	        relevant to the domain context. This map can include a "message" key
//	        which will be used as the primary error message for the exception.
//
// Returns:
//
//	A pointer to a new `Domain` instance.
func NewDomain(errors map[string]interface{}) *Domain {
	// Initialize the base CoreException with the given errors and a default
	// status of BadRequest, as domain errors often stem from invalid client input.
	base := NewInstance(errors, status.BadRequest)
	return &Domain{CoreException: *base}
}
