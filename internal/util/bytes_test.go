package util

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBytesToInt16(t *testing.T) {
	s, err := BytesToInt16StringArray([]byte{99, 0, 50, 55}, 10)
	fmt.Println(">>>", s)
	require.NoError(t, err)
}
