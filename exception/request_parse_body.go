// Package exception provides a structured and standardized approach to error handling
// within the application. This file defines a specific exception type for
// errors encountered during the parsing of a request body, leveraging the
// core exception handling mechanisms.
package exception

import (
	// status "github.com/osirisgate/golang-core/enum" is expected to provide
	// the `status.BadRequest` constant for setting the default status code.
	status "github.com/osirisgate/golang-core/enum"
)

// RequestParseBody is a specific exception type that signifies an error
// occurring when the application attempts to parse the body of an incoming
// request (e.g., JSON parsing error, XML malformation). This typically
// indicates that the client has sent a request body that is syntactically
// incorrect or otherwise unparseable by the server.
// It embeds `CoreException` to inherit all its properties and methods,
// ensuring consistent error reporting and formatting.
type RequestParseBody struct {
	CoreException // Embeds CoreException to inherit its fields and methods.
}

// NewRequestParseBody creates and returns a new `RequestParseBody` exception.
// It initializes the embedded `CoreException` with the provided error details
// and sets the default status code to `status.BadRequest`. This status code
// is highly appropriate for indicating that the client's request could not be
// understood by the server due to malformed input in the request body.
//
// Parameters:
//
//	errors: A map of string to interface{} containing detailed error information
//	        about the parsing failure. This map can include a "message" key
//	        which will be used as the primary error message for the exception.
//
// Returns:
//
//	A pointer to a new `RequestParseBody` instance.
func NewRequestParseBody(errors map[string]interface{}) *RequestParseBody {
	// Initialize the base CoreException with the given errors and a default
	// status of BadRequest, as body parsing errors are typically due to
	// malformed client requests.
	base := NewInstance(errors, status.BadRequest)
	return &RequestParseBody{CoreException: *base}
}
