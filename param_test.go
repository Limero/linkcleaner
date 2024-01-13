package linkcleaner

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveTrackingParameters(t *testing.T) {
	t.Run("global param", func(t *testing.T) {
		in := getURL(t, "https://feber.se/bil/nu-borjar-tesla-med-privatleasing-av-model-y/458956/?utm_source=rss&utm_medium=feed")
		out := RemoveTrackingParameters(in)
		assert.Equal(t, "https://feber.se/bil/nu-borjar-tesla-med-privatleasing-av-model-y/458956/", out.String())
	})

	t.Run("site specific param", func(t *testing.T) {
		in := getURL(t, "https://www.imdb.com/title/tt1375666/?ref_=wl_li_i")
		out := RemoveTrackingParameters(in)
		assert.Equal(t, "https://www.imdb.com/title/tt1375666/", out.String())
	})

	t.Run("site specific param should not apply to other site", func(t *testing.T) {
		in := getURL(t, "https://example.com/title/tt1375666/?ref_=wl_li_i")
		out := RemoveTrackingParameters(in)
		assert.Equal(t, "https://example.com/title/tt1375666/?ref_=wl_li_i", out.String())
	})
}
