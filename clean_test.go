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
	t.Run("raw text", func(t *testing.T) {
		in := `
	hello https://duckduckgo.com/?t=ffab&q=hello&ia=web there
	https://store.steampowered.com/app/413150/Stardew_Valley?snr=1_7_15__13
	`

		expected := `
	hello https://duckduckgo.com/?q=hello there
	https://store.steampowered.com/app/413150/Stardew_Valley
	`
		assert.Equal(t, expected, CleanAllURLsInString(in))
	})

	t.Run("html", func(t *testing.T) {
		in := `
			<a href="https://feber.se/film/amerikansk-remake-av-en-runda-till-pa-gang/461760/?utm_source=rss&amp;utm_medium=feed">https://feber.se/film/amerikansk-remake-av-en-runda-till-pa-gang/461760/</a><br /><br />

		L&auml;s mer om <a href="https://feber.se/tag/en+runda+till/">en runda till</a>, <a href="https://feber.se/tag/druk/">druk</a>, <a href="https://feber.se/tag/another+round/">another round</a>, <a href="https://feber.se/tag/remake/">remake</a>, <a href="https://feber.se/tag/Chris+Rock/">Chris Rock</a>, <a href="https://feber.se/tag/Thomas+Vinterberg/">Thomas Vinterberg</a>, <a href="https://feber.se/tag/Mads+Mikkelsen/">Mads Mikkelsen</a>`

		// TODO: This shouldn't be the expected output
		expected := `
			<a href="https://feber.se/film/amerikansk-remake-av-en-runda-till-pa-gang/461760/ /><br />

		L&auml;s mer om <a href="https://feber.se/tag/en+runda+till/%22%3Een runda till</a>, <a href="https://feber.se/tag/druk/%22%3Edruk%3C/a>, <a href="https://feber.se/tag/another+round/%22%3Eanother round</a>, <a href="https://feber.se/tag/remake/%22%3Eremake%3C/a>, <a href="https://feber.se/tag/Chris+Rock/%22%3EChris Rock</a>, <a href="https://feber.se/tag/Thomas+Vinterberg/%22%3EThomas Vinterberg</a>, <a href="https://feber.se/tag/Mads+Mikkelsen/%22%3EMads Mikkelsen</a>`
		assert.Equal(t, expected, CleanAllURLsInString(in))
	})
}

func getURL(t *testing.T, s string) *url.URL {
	u, err := url.Parse(s)
	require.NoError(t, err)
	return u
}
