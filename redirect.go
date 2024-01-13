package linkcleaner

import (
	"net/url"
	"strings"
)

func RemoveRedirects(u *url.URL) (*url.URL, error) {
	s := u.String()

	s, _ = strings.CutPrefix(s, "https://steamcommunity.com/linkfilter/?u=")

	return url.Parse(s)
}
