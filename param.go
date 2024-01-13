package linkcleaner

import "net/url"

func RemoveParameters(u *url.URL, parameters ...string) *url.URL {
	q := u.Query()
	for _, param := range parameters {
		q.Del(param)
	}
	u.RawQuery = q.Encode()
	return u
}

func GetSiteSpecificTrackingParameters(u *url.URL) []string {
	siteSpecificTrackingParams := map[string][]string{
		"duckduckgo.com": {
			"ia",
			"t",
		},
		"imdb.com": {
			"ref_",
		},
		"store.steampowered.com": {
			"snr",
		},
	}

	hostname := TrimPrefix(u.Hostname())

	if params, ok := siteSpecificTrackingParams[hostname]; ok {
		return params
	}

	return []string{}
}

func RemoveTrackingParameters(u *url.URL) *url.URL {
	globalTrackingParams := []string{
		"utm_source",
		"utm_medium",
	}
	siteSpecificTrackingParams := GetSiteSpecificTrackingParameters(u)
	allTrackingParams := append(globalTrackingParams, siteSpecificTrackingParams...)

	return RemoveParameters(u, allTrackingParams...)
}
