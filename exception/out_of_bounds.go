// Package exception provides a structured and standardized approach to error handling
// within the application. This file defines a specific exception type for
// errors related to out-of-bounds access or values, leveraging the core
// exception handling mechanisms.
package exception

import (
	// status "github.com/osirisgate/golang-core/enum" is expected to provide
	// the `status.UnprocessableContent` constant for setting the default status code.
	status "github.com/osirisgate/golang-core/enum"
)

// OutOfBounds is a specific exception type that signifies an error where
// an operation attempts to access data outside of its valid boundaries (e.g.,
// array index out of bounds, value outside an allowed range). This often
// indicates an issue with client-provided data or an internal logical error
// that leads to invalid data access.
// It embeds `CoreException` to inherit all its properties and methods,
// ensuring consistent error reporting and formatting.
type OutOfBounds struct {
	CoreException // Embeds CoreException to inherit its fields and methods.
}

// NewOutOfBounds creates and returns a new `OutOfBounds` exception.
// It initializes the embedded `CoreException` with the provided error details
// and sets the default status code to `status.UnprocessableContent`. This
// status code is appropriate when the server understands the content type
// of the request entity, but was unable to process the contained instructions
// due to semantic errors, which an out-of-bounds value can often be classified as.
//
// Parameters:
//
//	errors: A map of string to interface{} containing detailed error information
//	        about the out-of-bounds condition. This map can include a "message" key
//	        which will be used as the primary error message for the exception.
//
// Returns:
//
//	A pointer to a new `OutOfBounds` instance.
func NewOutOfBounds(errors map[string]interface{}) *OutOfBounds {
	// Initialize the base CoreException with the given errors and a default
	// status of UnprocessableContent, as out-of-bounds issues often relate
	// to semantically incorrect data.
	base := NewInstance(errors, status.UnprocessableContent)
	return &OutOfBounds{CoreException: *base}
}
