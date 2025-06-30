package exception_test

import (
	"errors"
	status "github.com/osirisgate/golang-core/enum"
	"github.com/osirisgate/golang-core/exception"
	"reflect"
	"testing"
)

func TestNewInstance(t *testing.T) {
	t.Run("WithExplicitMessage", func(t *testing.T) {
		errorsMap := map[string]interface{}{
			"message": "This is a custom message.",
			"details": map[string]interface{}{
				"field": "username",
			},
		}
		e := exception.NewInstance(errorsMap, status.BadRequest)

		if e.Message != "This is a custom message." {
			t.Errorf("Expected message to be set from map, got %q", e.Message)
		}
		if _, ok := e.Errors["message"]; ok {
			t.Error("Message key was not removed from the Errors map")
		}
		if e.StatusCode != status.BadRequest {
			t.Errorf("Expected status code %v, got %v", status.BadRequest, e.StatusCode)
		}
		if details, ok := e.Errors["details"].(map[string]interface{}); !ok || details["field"] != "username" {
			t.Error("Expected 'details' map with 'field' key to be in Errors map")
		}
	})

	t.Run("WithDefaultMessage", func(t *testing.T) {
		errorsMap := map[string]interface{}{
			"field": "password",
		}
		e := exception.NewInstance(errorsMap, status.NotFound)

		expectedMessage := status.NotFound.GetDescription()
		if e.Message != expectedMessage {
			t.Errorf("Expected message to be default description %q, got %q", expectedMessage, e.Message)
		}

		if _, ok := e.Errors["field"]; !ok {
			t.Error("Expected 'field' to be in Errors map")
		}
	})
}

func TestCoreExceptionInterfaceMethods(t *testing.T) {
	errorsMap := map[string]interface{}{
		"message": "Validation failed.",
		"errors": map[string]interface{}{
			"email": "invalid_email",
		},
		"details": map[string]interface{}{
			"error": "email_format_error",
			"code":  123,
		},
		"extra_data": "some_value",
	}

	coreException := exception.NewInstance(errorsMap, status.BadRequest)

	t.Run("Error", func(t *testing.T) {
		expected := "Validation failed."
		if coreException.Error() != expected {
			t.Errorf("Error() returned %q, expected %q", coreException.Error(), expected)
		}
	})

	t.Run("GetErrors", func(t *testing.T) {
		expected := map[string]interface{}{
			"errors":     map[string]interface{}{"email": "invalid_email"},
			"details":    map[string]interface{}{"error": "email_format_error", "code": 123},
			"extra_data": "some_value",
		}
		if !reflect.DeepEqual(coreException.GetErrors(), expected) {
			t.Errorf("GetErrors() returned unexpected map:\n got %+v,\n expected %+v", coreException.GetErrors(), expected)
		}
	})

	t.Run("GetDetails", func(t *testing.T) {
		expected := map[string]interface{}{"error": "email_format_error", "code": 123}
		if !reflect.DeepEqual(coreException.GetDetails(), expected) {
			t.Errorf("GetDetails() returned unexpected map:\n got %+v,\n expected %+v", coreException.GetDetails(), expected)
		}
	})

	t.Run("GetDetailsMessage", func(t *testing.T) {
		expected := "email_format_error"
		if coreException.GetDetailsMessage() != expected {
			t.Errorf("GetDetailsMessage() returned %q, expected %q", coreException.GetDetailsMessage(), expected)
		}
	})

	t.Run("GetErrorsForLog", func(t *testing.T) {
		expected := map[string]interface{}{
			"message":     "Validation failed.",
			"status_code": 400,
			"errors": map[string]interface{}{
				"errors":     map[string]interface{}{"email": "invalid_email"},
				"details":    map[string]interface{}{"error": "email_format_error", "code": 123},
				"extra_data": "some_value",
			},
			"stack_trace": coreException.StackTrace,
		}

		got := coreException.GetErrorsForLog()

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("GetErrorsForLog() returned unexpected map:\n got %+v,\n expected %+v", got, expected)
		}
	})

	t.Run("Format", func(t *testing.T) {
		expected := map[string]interface{}{
			"status":     status.ERROR,
			"error_code": 400,
			"message":    "Validation failed.",
			"errors":     map[string]interface{}{"email": "invalid_email"},
			"details":    map[string]interface{}{"error": "email_format_error", "code": 123},
			"extra_data": "some_value",
		}
		got := coreException.Format()
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("Format() returned unexpected map:\n got %+v,\n expected %+v", got, expected)
		}
	})
}

func TestNewBadFunctionCall(t *testing.T) {
	errorsMap := map[string]interface{}{
		"message": "Missing required argument.",
		"details": map[string]interface{}{
			"error": "argument_count_mismatch",
		},
	}

	err := exception.NewBadFunctionCall(errorsMap)

	var badFunctionCall *exception.BadFunctionCall
	ok := errors.As(err, &badFunctionCall)
	if !ok {
		t.Fatal("NewBadFunctionCall did not return a *BadFunctionCall type")
	}

	if err.StatusCode != status.BadRequest {
		t.Errorf("The status code of BadFunctionCall is %d, expected %d", err.StatusCode, status.BadRequest)
	}

	if err.Error() != "Missing required argument." {
		t.Errorf("The error message is incorrect. Got: '%s'", err.Error())
	}

	if msg := err.GetDetailsMessage(); msg != "argument_count_mismatch" {
		t.Errorf("GetDetailsMessage returned '%s', expected 'argument_count_mismatch'", msg)
	}
}
