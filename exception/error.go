// Package exception provides a structured and standardized approach to error handling
// within the application. This file defines a generic `Error` exception type,
// intended for unclassified or general application errors, leveraging the
// core exception handling mechanisms.
package exception

import (
	// status "github.com/osirisgate/golang-core/enum" is expected to provide
	// the `status.InternalServerError` constant for setting the default status code.
	status "github.com/osirisgate/golang-core/enum"
)

// Error is a generic exception type that can be used for general or
// uncategorized errors within the application. It serves as a fallback
// for situations where a more specific exception type (like `Domain`,
// `NotFound`, etc.) is not applicable or necessary.
// It embeds `CoreException` to inherit all its properties and methods,
// ensuring consistent error reporting and formatting.
type Error struct {
	CoreException // Embeds CoreException to inherit its fields and methods.
}

// NewError creates and returns a new `Error` exception.
// It initializes the embedded `CoreException` with the provided error details
// and sets the default status code to `status.InternalServerError`. This
// default indicates that the error is typically a server-side issue
// not directly attributable to client input.
//
// Parameters:
//
//	errors: A map of string to interface{} containing detailed error information.
//	        This map can include a "message" key which will be used as the
//	        primary error message for the exception.
//
// Returns:
//
//	A pointer to a new `Error` instance.
func NewError(errors map[string]interface{}) *Error {
	// Initialize the base CoreException with the given errors and a default
	// status of InternalServerError, as this is a generic error often indicating
	// a server-side problem.
	base := NewInstance(errors, status.InternalServerError)
	return &Error{CoreException: *base}
}
