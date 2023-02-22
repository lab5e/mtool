package regmap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeviceESM4450(t *testing.T) {
	errs := DeviceESM4450.Validate()
	assert.Empty(t, errs)
}
