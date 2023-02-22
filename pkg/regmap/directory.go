package regmap

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

// Directory contains a directory of devices.
type Directory struct {
	Devices map[string]Device `json:"devices" db:"-"`
}

var (
	ErrDuplicateDevice  = errors.New("duplicate device")
	ErrDeviceNotFound   = errors.New("device not found")
	ErrLoadingDirectory = errors.New("error loading directory")
)

// New directory.
func NewDirectory() *Directory {
	return &Directory{
		Devices: map[string]Device{},
	}
}

// Add device to directory.
func (d *Directory) Add(device Device) error {
	_, ok := d.Devices[device.Name]
	if ok {
		return ErrDuplicateDevice
	}
	d.Devices[device.Name] = device
	return nil
}

// Delete device from directory.
func (d *Directory) Delete(name string) error {
	_, ok := d.Devices[name]
	if !ok {
		return ErrDeviceNotFound
	}
	delete(d.Devices, name)
	return nil
}

// Get device from directory.
func (d *Directory) Get(name string) (Device, error) {
	dev, ok := d.Devices[name]
	if !ok {
		return Device{}, ErrDeviceNotFound
	}
	return dev, nil
}

// LoadDirectory directory from io.Reader.
func LoadDirectory(r io.Reader) (*Directory, error) {
	decoder := json.NewDecoder(r)

	dir := NewDirectory()
	err := decoder.Decode(dir)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrLoadingDirectory, err)
	}
	return dir, nil
}

// Save directory to writer.
func (d *Directory) Save(w io.Writer) error {
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ")
	return encoder.Encode(d)
}
