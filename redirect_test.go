package linkcleaner

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRemoveRedirects(t *testing.T) {
	t.Run("steam link redirect", func(t *testing.T) {
		in := getURL(t, "https://steamcommunity.com/linkfilter/?u=https%3A%2F%2Fstore.epicgames.com%2Fp%2Fjitsu-squad-af3f2f")
		out, err := RemoveRedirects(in)
		require.NoError(t, err)
		assert.Equal(t, "https%3A%2F%2Fstore.epicgames.com%2Fp%2Fjitsu-squad-af3f2f", out.String())
	})
}
