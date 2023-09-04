// Package milkvduo provides convenience mappings from Milk-V Duo pin names to
// GPIO chips and offsets.
package milkvduo

import (
	"errors"
	"regexp"
	"strconv"
)

type LineID struct {
	Chip   string
	Offset int
}

var GPIO_TO_CHIP = map[string]string{
	"GPIOA":    "gpiochip0",
	"GPIOB":    "gpiochip1",
	"GPIOC":    "gpiochip2",
	"GPIOD":    "gpiochip3",
	"PWR_GPIO": "gpiochip4",
}

// ErrInvalid indicates the pin name does not match a known pin.
var ErrInvalid = errors.New("invalid pin name/number")

// rangeCheck checks if the given integer is within the valid range.
//
// It takes an integer parameter 'p' and returns an integer and an error.
func rangeCheck(p int) (int, error) {
	if p < 2 || p > 29 {
		return 0, ErrInvalid
	}
	return p, nil
}

// PinLineID generates a LineID from a given string.
//
// Takes a string as input and returns a LineID and an error.
func PinLineID(s string) (LineID, error) {
	re := regexp.MustCompile("([A-Z_]+)([0-9]{1,2})")
	m := re.FindStringSubmatch(s)
	if m == nil {
		return LineID{}, ErrInvalid
	}
	chip, ok := GPIO_TO_CHIP[m[1]]
	if !ok {
		return LineID{}, ErrInvalid
	}
	offset, err := strconv.Atoi(m[2])
	if err != nil {
		return LineID{}, ErrInvalid
	}
	offset, err = rangeCheck(offset)
	if err != nil {
		return LineID{}, ErrInvalid
	}
	return LineID{
		Chip:   chip,
		Offset: offset,
	}, nil
}

// MustPinGpio generates a LineID from a string representation of a GPIO pin and panics if there is an error.
//
// s: the string representation of the GPIO pin.
// Returns: the LineID generated from the string representation of the GPIO pin.
func MustPinGpio(s string) LineID {
	v, err := PinLineID(s)
	if err != nil {
		panic(err)
	}
	return v
}
