package linkcleaner

import "net/url"

func Unescape(u *url.URL) (*url.URL, error) {
	s, err := url.PathUnescape(u.String())
	if err != nil {
		return nil, err
	}
	return url.Parse(s)
}

func CleanURL(u *url.URL) (*url.URL, error) {
	cu, err := RemoveRedirects(u)
	if err != nil {
		return nil, err
	}

	cu = RemoveTrackingParameters(cu)

	cu, err = Unescape(cu)
	if err != nil {
		return nil, err
	}

	return cu, nil
}

func CleanURLString(s string) (*url.URL, error) {
	u, err := url.Parse(s)
	if err != nil {
		return nil, err
	}
	return CleanURL(u)
}
