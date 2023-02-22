package mb

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMB(t *testing.T) {
	m, err := Connect("/dev/cu.usbserial-A50285BI")
	require.NoError(t, err)
	defer m.Close()

	data, err := m.readHoldingRegister(1, 1)
	require.NoError(t, err)

	fmt.Println(data)
}
