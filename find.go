package linkcleaner

import (
	"net/url"
	"regexp"
)

var urlRegex = regexp.MustCompile(`(?i)\b(?:https?|ftp):\/\/\S+\b`)

type URLPos struct {
	URL    *url.URL
	Start  int
	Length int
}

func FindURLs(s string) []*url.URL {
	matches := urlRegex.FindAllString(s, -1)
	urls := make([]*url.URL, 0, len(matches))
	for _, match := range matches {
		parsedURL, err := url.Parse(match)
		if err != nil {
			continue
		}
		urls = append(urls, parsedURL)
	}
	return urls
}

func FindURLsWithPos(s string) []URLPos {
	matches := urlRegex.FindAllStringIndex(s, -1)
	urlPositions := make([]URLPos, 0, len(matches))

	for _, match := range matches {
		start, end := match[0], match[1]
		urlString := s[start:end]

		parsedURL, err := url.Parse(urlString)
		if err != nil {
			continue
		}

		urlPositions = append(urlPositions, URLPos{
			URL:    parsedURL,
			Start:  start,
			Length: end - start,
		})
	}

	return urlPositions
}
