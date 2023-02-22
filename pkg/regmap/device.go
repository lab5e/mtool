package regmap

import (
	"errors"
	"fmt"
	"net/url"
)

// Device represents a Modbus device.
type Device struct {
	Name        string     `json:"name" db:"name"`
	Description string     `json:"description" db:"description"`
	URL         string     `json:"url" db:"url"`
	Registers   []Register `json:"registers" db:"-"`
}

// errors for device
var (
	ErrDuplicateRegister  = errors.New("duplicate register")
	ErrNameMissing        = errors.New("name is missing")
	ErrDescriptionMissing = errors.New("description is missing")
	ErrMalformedURL       = errors.New("malformed URL")
	ErrAccessUnspecified  = errors.New("access mode unspecified")
)

// Validate a device entry.
func (d *Device) Validate() []error {
	errors := []error{}

	// Make sure we have a name.
	if d.Name == "" {
		errors = append(errors, ErrNameMissing)
	}

	// Make sure we have a description.
	if d.Description == "" {
		errors = append(errors, ErrDescriptionMissing)
	}

	// If we have an URL, make sure the URL is correctly formatted.
	if d.URL != "" {
		_, err := url.Parse(d.URL)
		if err != nil {
			errors = append(errors, fmt.Errorf("%w url=[%s]: %v", ErrMalformedURL, d.URL, err))
		}
	}

	// Check registers
	m := map[uint16]Register{}
	for _, reg := range d.Registers {
		// Make sure that there are no duplicate registers
		if _, ok := m[reg.Address]; ok {
			errors = append(errors, fmt.Errorf("%w address=[%d] symbol=[%s]", ErrDuplicateRegister, reg.Address, reg.Symbol))
		}

		// Make sure the access mode of the register is known
		if reg.Access == ModeUnspecified {
			errors = append(errors, fmt.Errorf("%w address=[%d] symbol=[%s]", ErrAccessUnspecified, reg.Address, reg.Symbol))
		}
	}
	return errors
}
