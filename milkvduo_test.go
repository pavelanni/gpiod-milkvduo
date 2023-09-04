package milkvduo

import (
	"reflect"
	"testing"
)

func TestPinLineID(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		expectedLineID LineID
		expectedError  error
	}{
		{
			name:           "Valid input",
			input:          "GPIOA14",
			expectedLineID: LineID{Chip: "gpiochip0", Offset: 14},
			expectedError:  nil,
		},
		{
			name:           "Invalid input - no match",
			input:          "invalid",
			expectedLineID: LineID{},
			expectedError:  ErrInvalid,
		},
		{
			name:           "Invalid input - unknown GPIO chip",
			input:          "GPIOH1",
			expectedLineID: LineID{},
			expectedError:  ErrInvalid,
		},
		{
			name:           "Invalid input - invalid offset",
			input:          "GPIO_100",
			expectedLineID: LineID{},
			expectedError:  ErrInvalid,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lineID, err := PinLineID(tt.input)
			if !reflect.DeepEqual(lineID, tt.expectedLineID) {
				t.Errorf("got %v, want %v", lineID, tt.expectedLineID)
			}
			if err != tt.expectedError {
				t.Errorf("got %v, want %v", err, tt.expectedError)
			}
		})
	}
}
