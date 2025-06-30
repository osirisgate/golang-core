package status_tests

import (
	status "github.com/osirisgate/golang-core/enum"
	"testing"
)

func TestStatusGetValue(t *testing.T) {
	tests := []struct {
		name          string
		inputStatus   status.Status
		expectedValue string
	}{
		{
			name:          "SUCCESS enum check",
			inputStatus:   status.SUCCESS,
			expectedValue: "success",
		},
		{
			name:          "ERROR enum check",
			inputStatus:   status.ERROR,
			expectedValue: "error",
		},
		{
			name:          "Custom enum check",
			inputStatus:   status.Status("pending"),
			expectedValue: "pending",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualValue := tt.inputStatus.GetValue()

			if actualValue != tt.expectedValue {
				t.Errorf("GetValue() for enum %q returned %q, but expected %q", tt.inputStatus, actualValue, tt.expectedValue)
			}
		})
	}
}

func TestNewStatus(t *testing.T) {
	customStatus := status.NewStatus("processing")

	expectedValue := "processing"
	actualValue := customStatus.GetValue()

	if actualValue != expectedValue {
		t.Errorf("NewStatus() created a enum whose GetValue() returned %q, but expected %q", actualValue, expectedValue)
	}

	if customStatus == status.SUCCESS {
		t.Error("The custom enum should not be equal to enum.SUCCESS.")
	}

	if customStatus == status.ERROR {
		t.Error("The custom enum should not be equal to enum.ERROR.")
	}
}
