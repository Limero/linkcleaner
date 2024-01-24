package linkcleaner

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindURLs(t *testing.T) {
	in := `
	lorem ipsum https://example.com ipsum lorem https://example.com/test?a=123
	`
	out := FindURLs(in)
	assert.Len(t, out, 2)
	assert.Equal(t, "https://example.com", out[0].String())
	assert.Equal(t, "https://example.com/test?a=123", out[1].String())
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
