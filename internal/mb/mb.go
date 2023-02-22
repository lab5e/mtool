// Package mb is our modbus wrapper.
package mb

import (
	"errors"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/goburrow/modbus"
)

type MB struct {
	mu      sync.Mutex
	handler *modbus.RTUClientHandler
}

type Config struct {
	Baud     int
	DataBits int
	Parity   string
	StopBits int
	Timeout  time.Duration
	Debug    bool
}

// DefaultConfig is the default configuration for connection if no
// config parameters are given.
var DefaultConfig = Config{
	Baud:     9600,
	DataBits: 8,
	Parity:   "N",
	StopBits: 1,
	Timeout:  5 * time.Second,
	Debug:    false,
}

// errors
var (
	ErrConnect = errors.New("error connecting to serial device")
)

// Connect to modbus interface.  You can optionally include a Config if the
// default values in DefaultConfig do not fit your needs.
func Connect(serialDevice string, c ...Config) (*MB, error) {
	config := DefaultConfig
	if len(c) > 0 {
		config = c[0]
	}

	handler := modbus.NewRTUClientHandler(serialDevice)
	handler.BaudRate = config.Baud
	handler.DataBits = config.DataBits
	handler.Parity = config.Parity
	handler.StopBits = config.StopBits
	handler.Timeout = config.Timeout
	handler.SlaveId = 100

	if config.Debug {
		handler.Logger = log.New(os.Stderr, "modbus", log.LstdFlags)
	}

	err := handler.Connect()
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrConnect, err)
	}

	return &MB{handler: handler}, err
}

func (m *MB) readHoldingRegister(deviceID byte, addr uint16) ([]byte, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.handler.SlaveId = deviceID

	client := modbus.NewClient(m.handler)
	return client.ReadHoldingRegisters(addr-1, 1)
}

func (m *MB) Close() error {
	m.mu.Lock()
	defer m.mu.Unlock()
	err := m.handler.Close()
	if err != nil {
		return err
	}
	m.handler = nil // trigger nil errors if operating on closed instance
	return nil
}
