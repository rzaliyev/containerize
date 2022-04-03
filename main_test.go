package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEcho(t *testing.T) {

	t.Run("testing happy paht", func(t *testing.T) {
		err := echo([]string{"bin-name", "hello", "world!"})
		require.NoError(t, err)
	})

	t.Run("testing no args error", func(t *testing.T) {
		err := echo([]string{})
		require.Error(t, err)
	})

}
