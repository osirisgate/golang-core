package status_tests

import (
	"reflect"
	"testing"

	status "github.com/osirisgate/golang-core/enum"
)

func TestGetValue(t *testing.T) {
	tests := []struct {
		name     string
		input    status.StatusCode
		expected int
	}{
		{name: "OK", input: status.OK, expected: 200},
		{name: "NotFound", input: status.NotFound, expected: 404},
		{name: "InternalServerError", input: status.InternalServerError, expected: 500},
		{name: "Continue", input: status.Continue, expected: 100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.input.GetValue(); got != tt.expected {
				t.Errorf("GetValue() for %v returned %d, but expected %d", tt.input, got, tt.expected)
			}
		})
	}
}

func TestGetDescription(t *testing.T) {
	tests := []struct {
		name     string
		input    status.StatusCode
		expected string
	}{
		{name: "OK", input: status.OK, expected: "OK"},
		{name: "NotFound", input: status.NotFound, expected: "Not Found"},
		{name: "Unauthorized", input: status.Unauthorized, expected: "Unauthorized"},
		{name: "IMATeapot", input: status.IMATeapot, expected: "I'm a teapot"},
		{name: "UnknownStatus", input: status.StatusCode(999), expected: "Unknown Status Code"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.input.GetDescription(); got != tt.expected {
				t.Errorf("GetDescription() for %v returned %q, but expected %q", tt.input, got, tt.expected)
			}
		})
	}
}

func TestNewStatusCode(t *testing.T) {
	tests := []struct {
		name            string
		input           int
		expectedCode    status.StatusCode
		expectedSuccess bool
	}{
		{name: "ValidCode-200", input: 200, expectedCode: status.OK, expectedSuccess: true},
		{name: "ValidCode-404", input: 404, expectedCode: status.NotFound, expectedSuccess: true},
		{name: "InvalidCode", input: 999, expectedCode: status.StatusCode(0), expectedSuccess: false},
		{name: "ValidCode-500", input: 500, expectedCode: status.InternalServerError, expectedSuccess: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCode, gotSuccess := status.NewStatusCode(tt.input)

			if gotSuccess != tt.expectedSuccess {
				t.Errorf("NewStatusCode() for value %d returned success as %v, but expected %v", tt.input, gotSuccess, tt.expectedSuccess)
			}

			if gotCode != tt.expectedCode {
				t.Errorf("NewStatusCode() for value %d returned code %v, but expected %v", tt.input, gotCode, tt.expectedCode)
			}
		})
	}
}

func TestGetStatusTexts(t *testing.T) {
	statusTexts := status.GetStatusTexts()
	expectedSize := 62

	if len(statusTexts) != expectedSize {
		t.Errorf("The map has a size of %d, but %d was expected", len(statusTexts), expectedSize)
	}

	if statusTexts[status.OK] != "OK" {
		t.Errorf("The description for OK is incorrect: got %q, expected %q", statusTexts[status.OK], "OK")
	}

	statusTexts[status.OK] = "Not OK"

	originalStatusTexts := status.GetStatusTexts()
	if originalStatusTexts[status.OK] == "Not OK" {
		t.Error("The original map was modified, indicating that a copy was not returned.")
	}

	if reflect.DeepEqual(statusTexts, originalStatusTexts) {
		t.Error("The two maps are identical, which indicates a copy was not created.")
	}
}
