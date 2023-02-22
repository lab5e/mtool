package regmap

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDirectory(t *testing.T) {
	// Create a directory and add a device to it
	dir := NewDirectory()
	require.NoError(t, dir.Add(deviceESM4450))
	require.ErrorIs(t, ErrDuplicateDevice, dir.Add(deviceESM4450))

	// Make sure we can retrieve device
	dev, err := dir.Get(deviceESM4450.Name)
	require.NoError(t, err)
	require.Equal(t, deviceESM4450, dev)

	// Make sure we can save and load directory
	buf := bytes.Buffer{}
	require.NoError(t, dir.Save(&buf))
	newDir, err := LoadDirectory(&buf)
	require.NoError(t, err)
	require.Equal(t, dir, newDir)

	// Make sure delete works
	require.NoError(t, dir.Delete(deviceESM4450.Name))

	_, err = dir.Get(deviceESM4450.Name)
	require.ErrorIs(t, ErrDeviceNotFound, err)
	require.ErrorIs(t, ErrDeviceNotFound, dir.Delete(deviceESM4450.Name))
}
