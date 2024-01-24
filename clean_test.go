package linkcleaner

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnescape(t *testing.T) {
	in := getURL(t, "https%3A%2F%2Fstore.epicgames.com%2Fp%2Fjitsu-squad-af3f2f")
	out, err := Unescape(in)
	require.NoError(t, err)
	assert.Equal(t, "https://store.epicgames.com/p/jitsu-squad-af3f2f", out.String())
}

func TestCleanURLString(t *testing.T) {
	tt := map[string]string{
		"https://steamcommunity.com/linkfilter/?u=https%3A%2F%2Fstore.epicgames.com%2Fp%2Fjitsu-squad-af3f2f":      "https://store.epicgames.com/p/jitsu-squad-af3f2f",
		"https://feber.se/bil/nu-borjar-tesla-med-privatleasing-av-model-y/458956/?utm_source=rss&utm_medium=feed": "https://feber.se/bil/nu-borjar-tesla-med-privatleasing-av-model-y/458956/",
		"https://duckduckgo.com/?t=ffab&q=hello&ia=web":                                                            "https://duckduckgo.com/?q=hello",
		"https://store.steampowered.com/app/413150/Stardew_Valley?snr=1_7_15__13":                                  "https://store.steampowered.com/app/413150/Stardew_Valley",
	}
	for in, expected := range tt {
		out, err := CleanURLString(in)
		require.NoError(t, err)
		assert.Equal(t, expected, out.String())
	}
}

func TestCleanAllURLsInString(t *testing.T) {
	in := `
	hello https://duckduckgo.com/?t=ffab&q=hello&ia=web there
	https://store.steampowered.com/app/413150/Stardew_Valley?snr=1_7_15__13
	`

	expected := `
	hello https://duckduckgo.com/?q=hello there
	https://store.steampowered.com/app/413150/Stardew_Valley
	`
	assert.Equal(t, expected, CleanAllURLsInString(in))
}

func getURL(t *testing.T, s string) *url.URL {
	u, err := url.Parse(s)
	require.NoError(t, err)
	return u
}
