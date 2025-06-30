// Package exception provides a structured and standardized approach to error handling
// within the application. This file defines a specific exception type for
// errors related to encountering an unexpected value, leveraging the core
// exception handling mechanisms.
package exception

import (
	// status "github.com/osirisgate/golang-core/enum" is expected to provide
	// the `status.UnprocessableContent` constant for setting the default status code.
	status "github.com/osirisgate/golang-core/enum"
)

// UnexpectedValue is a specific exception type that signifies an error where
// a value encountered during processing is valid in its type but is not one
// of the expected or allowed discrete values. This often indicates a semantic
// error in client input or an unexpected state in internal data.
// It embeds `CoreException` to inherit all its properties and methods,
// ensuring consistent error reporting and formatting.
type UnexpectedValue struct {
	CoreException // Embeds CoreException to inherit its fields and methods.
}

// NewUnexpectedValue creates and returns a new `UnexpectedValue` exception.
// It initializes the embedded `CoreException` with the provided error details
// and sets the default status code to `status.UnprocessableContent`. This
// status code is appropriate when the server understands the content type
// of the request entity, but was unable to process the contained instructions
// due to semantic errors, such as a value that is valid in type but unexpected in context.
//
// Parameters:
//
//	errors: A map of string to interface{} containing detailed error information
//	        about the unexpected value. This map can include a "message" key
//	        which will be used as the primary error message for the exception.
//
// Returns:
//
//	A pointer to a new `UnexpectedValue` instance.
func NewUnexpectedValue(errors map[string]interface{}) *UnexpectedValue {
	// Initialize the base CoreException with the given errors and a default
	// status of UnprocessableContent, as unexpected values often relate to
	// semantically incorrect input data that cannot be processed.
	base := NewInstance(errors, status.UnprocessableContent)
	return &UnexpectedValue{CoreException: *base}
}
