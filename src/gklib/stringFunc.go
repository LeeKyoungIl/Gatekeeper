package gklib

import "strings"

func TrimSuffix(s, suffix string) string {
	if strings.HasSuffix(s, suffix) {
		s = s[:len(s)-len(suffix)]
	}
	return s
}

func AddSuffix(s, suffix string) string {
	if !strings.HasSuffix(s, suffix) {
		s = s + suffix
	}
	return s
}
