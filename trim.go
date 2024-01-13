package linkcleaner

import "strings"

func TrimPrefix(hostname string) string {
	prefixes := []string{
		"m.",
		"www.",
	}
	for _, prefix := range prefixes {
		hostname = strings.TrimPrefix(hostname, prefix)
	}
	return hostname
}
