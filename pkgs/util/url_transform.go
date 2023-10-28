package util

import "strings"

// Very basic URL Transform util
func TransformURL(url string) string {
	if !strings.HasPrefix(url, "https") || !strings.HasPrefix(url, "http") {
		url = "https://" + url
	}
	return url
}
