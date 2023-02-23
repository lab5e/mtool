package mb

import (
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

const testDevice = "/dev/cu.usbserial-A50285BI"

func TestMB(t *testing.T) {
	// if the device doesn't exist we skip the test
	if _, err := os.Stat(testDevice); errors.Is(err, os.ErrNotExist) {
		return
	}

	m, err := Connect(testDevice)
	require.NoError(t, err)
	defer m.Close()

	data, err := m.readHoldingRegister(1, 1)
	require.NoError(t, err)

	fmt.Println(data)
}
