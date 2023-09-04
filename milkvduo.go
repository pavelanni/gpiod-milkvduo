// Package milkvduo provides convenience mappings from Milk-V Duo pin names to
// GPIO chips and offsets.
package milkvduo

import (
	"errors"
	"regexp"
	"strconv"
)

type GPIO struct {
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
var ErrInvalid = errors.New("invalid pin number")

func rangeCheck(p int) (int, error) {
	if p < 2 || p > 29 {
		return 0, ErrInvalid
	}
	return p, nil
}

func PinGpio(s string) (GPIO, error) {
	re := regexp.MustCompile("([A-Z_]+)([0-9]{1,2})")
	m := re.FindStringSubmatch(s)
	if m == nil {
		return GPIO{}, ErrInvalid
	}
	chip, ok := GPIO_TO_CHIP[m[1]]
	if !ok {
		return GPIO{}, ErrInvalid
	}
	offset, err := strconv.Atoi(m[2])
	if err != nil {
		return GPIO{}, ErrInvalid
	}
	offset, err = rangeCheck(offset)
	if err != nil {
		return GPIO{}, ErrInvalid
	}
	return GPIO{
		Chip:   chip,
		Offset: offset,
	}, nil
}

// MustPinGpio converts the string to the corresponding pin number or panics if that
// is not possible.
func MustPinGpio(s string) GPIO {
	v, err := PinGpio(s)
	if err != nil {
		panic(err)
	}
	return v
}
