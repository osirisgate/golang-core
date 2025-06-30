// Package exception provides a structured and standardized approach to error handling
// within the application. It defines an interface and a concrete type for representing
// and managing exceptions, including their messages, status codes, detailed errors,
// and stack traces.
package exception

import (
	// "github.com/osirisgate/golang-core/status" is expected to provide
	// the 'status.StatusCode' type and the 'status.ERROR' constant.
	"github.com/osirisgate/golang-core/enum"
	"runtime/debug" // Used for capturing stack trace information.
)

// CoreInterface defines the contract that any core exception type must satisfy.
// It extends the standard Go `error` interface, meaning any implementation
// must also provide an `Error()` string method.
type CoreInterface interface {
	error // Embeds the built-in error interface, requiring an Error() string method.

	// Format returns a map representation of the exception, typically used for
	// standardized API responses or structured logging. It includes general
	// error information like status, error code, and message.
	Format() map[string]interface{}

	// GetErrors returns the raw map of additional, specific error details
	// associated with the exception. This map can contain various key-value pairs
	// providing context about the error.
	GetErrors() map[string]interface{}

	// GetDetails extracts and returns a specific "details" map from the main
	// Errors map, if it exists and is of the correct type. This is useful for
	// retrieving nested error information. Returns an empty map if not found.
	GetDetails() map[string]interface{}

	// GetDetailsMessage attempts to extract a string value under the "error" key
	// within the "details" map. This can be used to retrieve a specific,
	// human-readable error message from nested details. Returns an empty string
	// if the key is not found or is not a string.
	GetDetailsMessage() string

	// GetStatusCode returns the integer value of the HTTP-like status code
	// associated with this exception.
	GetStatusCode() int

	// GetErrorsForLog returns a map containing comprehensive error information
	// specifically formatted for logging purposes. This includes the message,
	// status code, full errors map, and the stack trace.
	GetErrorsForLog() map[string]interface{}

	// GetStackTrace returns the full stack trace captured at the moment
	// the exception was created. This is crucial for debugging.
	GetStackTrace() string
}

// CoreException is the concrete implementation of the CoreInterface.
// It encapsulates all relevant information about an error, including its
// primary message, a status code, a flexible map for additional error details,
// and the execution stack trace.
type CoreException struct {
	Message    string                 // The primary human-readable message describing the exception.
	StatusCode status.StatusCode      // The HTTP-like status code associated with the exception (e.g., 400, 500).
	Errors     map[string]interface{} // A flexible map to hold additional, granular error information.
	StackTrace string                 // The stack trace captured when this exception was initialized.
}

// NewInstance creates and returns a new CoreException.
// This constructor handles the initialization of the exception, including
// setting the primary message and capturing the stack trace.
//
// Parameters:
//
//	errors: A map that can contain various error details. If this map includes
//	        a key "message" with a string value, that value will be used as
//	        the CoreException's main Message, and the "message" key will be
//	        removed from the `errors` map itself.
//	defaultStatusCode: The default `status.StatusCode` to use if no explicit
//	                   message is provided within the `errors` map. Its
//	                   description will be used as the message in such cases.
//
// Returns:
//
//	A pointer to a newly created CoreException instance.
func NewInstance(errors map[string]interface{}, defaultStatusCode status.StatusCode) *CoreException {
	message, ok := errors["message"].(string)
	if !ok || message == "" {
		// If no message is provided in the errors map, or it's empty,
		// use the description of the default status code as the message.
		message = defaultStatusCode.GetDescription()
	} else {
		// If a message was provided in the errors map, remove it to avoid redundancy
		// in the `Errors` field, as it's now the main `Message`.
		delete(errors, "message")
	}

	return &CoreException{
		Message:    message,
		StatusCode: defaultStatusCode,
		Errors:     errors,
		// Capture the current goroutine's stack trace at the point of exception creation.
		StackTrace: string(debug.Stack()),
	}
}

// Error implements the `error` interface for CoreException.
// It returns the primary message of the exception.
func (e CoreException) Error() string {
	return e.Message
}

// GetStatusCode returns the integer representation of the exception's
// `StatusCode`.
func (e CoreException) GetStatusCode() int {
	return e.StatusCode.GetValue()
}

// GetErrors returns the map containing additional error details associated
// with the exception.
func (e CoreException) GetErrors() map[string]interface{} {
	return e.Errors
}

// GetDetails attempts to retrieve a sub-map named "details" from the `Errors` map.
// This is commonly used for more granular, structured error information.
// Returns an empty map if "details" is not present or is not a map[string]interface{}.
func (e CoreException) GetDetails() map[string]interface{} {
	if details, ok := e.Errors["details"].(map[string]interface{}); ok {
		return details
	}
	return map[string]interface{}{} // Return an empty map if details are not found or not of the expected type.
}

// GetDetailsMessage attempts to extract a string message from the "error" key
// within the "details" map. This provides a way to get a specific error message
// that might be nested within the error details.
// Returns an empty string if the "details" map or the "error" key within it
// is not found or not a string.
func (e CoreException) GetDetailsMessage() string {
	details := e.GetDetails()
	if msg, ok := details["error"].(string); ok {
		return msg
	}
	return "" // Return an empty string if the specific error message is not found.
}

// GetErrorsForLog returns a map specifically formatted for logging purposes.
// This map includes the main message, the status code, the full `Errors` map,
// and the `StackTrace`, providing a complete context for logging systems.
func (e CoreException) GetErrorsForLog() map[string]interface{} {
	return map[string]interface{}{
		"message":     e.Message,
		"status_code": e.StatusCode.GetValue(),
		"errors":      e.Errors,
		"stack_trace": e.StackTrace,
	}
}

// GetStackTrace returns the complete stack trace string associated with
// the exception. This is invaluable for debugging and pinpointing the
// origin of the error.
func (e CoreException) GetStackTrace() string {
	return e.StackTrace
}

// Format returns a map representation of the exception, designed for
// standardized output, such as API responses. It includes a general "status"
// (assumed to be a constant like `status.ERROR`), an "error_code"
// corresponding to the status code, and the primary "message". Any additional
// key-value pairs from the `Errors` map are flattened directly into this
// formatted output.
func (e CoreException) Format() map[string]interface{} {
	formatted := map[string]interface{}{
		"status":     status.ERROR, // Assumed to be a constant indicating a general error status.
		"error_code": e.StatusCode.GetValue(),
		"message":    e.Message,
	}

	// If there are additional errors in the `Errors` map, merge them
	// into the top level of the formatted output.
	if e.Errors != nil {
		for key, value := range e.Errors {
			formatted[key] = value
		}
	}

	return formatted
}
