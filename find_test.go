package linkcleaner

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFindURLs(t *testing.T) {
	t.Run("find in text", func(t *testing.T) {
		in := `
	lorem ipsum https://example.com ipsum lorem https://example.com/test?a=123
	`
		out := FindURLs(in)
		assert.Len(t, out, 2)
		assert.Equal(t, "https://example.com", out[0].String())
		assert.Equal(t, "https://example.com/test?a=123", out[1].String())
	})

	t.Run("find individual", func(t *testing.T) {
		tt := map[string]string{
			"<a href=\"https://example.com\">":                        "https://example.com",
			"https://example.com/?utm_source=rss&amp;utm_medium=feed": "https://example.com/?utm_source=rss&amp;utm_medium=feed",
			"https://example.com/tag/Hello+World/":                    "https://example.com/tag/Hello+World",
		}

		for in, expected := range tt {
			out := FindURLs(in)
			require.Len(t, out, 1)
			assert.Equal(t, expected, out[0].String())
		}
	})
}

func TestFindURLsWithPos(t *testing.T) {
	in := `
	lorem ipsum https://example.com ipsum lorem https://example.com/test?a=123
	`
	out := FindURLsWithPos(in)
	assert.Len(t, out, 2)

	assert.Equal(t, "https://example.com", out[0].URL.String())
	assert.Equal(t, 14, out[0].Start)
	assert.Equal(t, len("https://example.com"), out[0].Length)

	assert.Equal(t, "https://example.com/test?a=123", out[1].URL.String())
	assert.Equal(t, 46, out[1].Start)
	assert.Equal(t, len("https://example.com/test?a=123"), out[1].Length)
}
