package section01

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestChangeValue(t *testing.T) {
	t.Run("should change value to 20", func(t *testing.T) {
		require := require.New(t)
		value := 10

		expected := 20
		ChangeValue(&value)

		require.NotEqual(value, 10)
		require.Equal(expected, value)
	})
}
