// Package exception provides a structured and standardized approach to error handling
// within the application. This file defines a specific exception type for
// errors related to values being outside of a valid range, leveraging the core
// exception handling mechanisms.
package exception

import (
	// status "github.com/osirisgate/golang-core/enum" is expected to provide
	// the `status.UnprocessableContent` constant for setting the default status code.
	status "github.com/osirisgate/golang-core/enum"
)

// Range is a specific exception type that indicates an error where a value
// provided is outside the expected or allowed boundaries of a specific range.
// This is often encountered during validation of numerical or ordered data.
// It embeds `CoreException` to inherit all its properties and methods,
// ensuring consistent error reporting and formatting.
type Range struct {
	CoreException // Embeds CoreException to inherit its fields and methods.
}

// NewRange creates and returns a new `Range` exception.
// It initializes the embedded `CoreException` with the provided error details
// and sets the default status code to `status.UnprocessableContent`. This
// status code is appropriate when the server understands the content type
// of the request entity, but was unable to process the contained instructions
// due to semantic errors, such as a value falling outside a valid range.
//
// Parameters:
//
//	errors: A map of string to interface{} containing detailed error information
//	        about the range violation. This map can include a "message" key
//	        which will be used as the primary error message for the exception.
//
// Returns:
//
//	A pointer to a new `Range` instance.
func NewRange(errors map[string]interface{}) *Range {
	// Initialize the base CoreException with the given errors and a default
	// status of UnprocessableContent, as range issues often relate to
	// semantically incorrect input data.
	base := NewInstance(errors, status.UnprocessableContent)
	return &Range{CoreException: *base}
}
