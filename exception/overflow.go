// Package exception provides a structured and standardized approach to error handling
// within the application. This file defines a specific exception type for
// arithmetic or data overflow errors, leveraging the core exception handling
// mechanisms.
package exception

import (
	// status "github.com/osirisgate/golang-core/enum" is expected to provide
	// the `status.UnprocessableContent` constant for setting the default status code.
	status "github.com/osirisgate/golang-core/enum"
)

// Overflow is a specific exception type that indicates an error condition
// where a numerical operation or data storage attempt exceeds the maximum
// capacity of its data type or allocated space. This can occur, for example,
// when performing an arithmetic operation that results in a value too large
// to be represented, or when attempting to store data that exceeds a buffer size.
// It embeds `CoreException` to inherit all its properties and methods,
// ensuring consistent error reporting and formatting.
type Overflow struct {
	CoreException // Embeds CoreException to inherit its fields and methods.
}

// NewOverflow creates and returns a new `Overflow` exception.
// It initializes the embedded `CoreException` with the provided error details
// and sets the default status code to `status.UnprocessableContent`. This
// status code is appropriate when the server understands the content type
// of the request entity, but was unable to process the contained instructions
// due to semantic errors, which an overflow condition can often be classified as.
//
// Parameters:
//
//	errors: A map of string to interface{} containing detailed error information
//	        about the overflow condition. This map can include a "message" key
//	        which will be used as the primary error message for the exception.
//
// Returns:
//
//	A pointer to a new `Overflow` instance.
func NewOverflow(errors map[string]interface{}) *Overflow {
	// Initialize the base CoreException with the given errors and a default
	// status of UnprocessableContent, as overflow issues often relate to
	// semantically incorrect or excessively large data.
	base := NewInstance(errors, status.UnprocessableContent)
	return &Overflow{CoreException: *base}
}
