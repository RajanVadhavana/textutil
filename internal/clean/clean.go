package clean

import "strings"

func Normalize(s string) (string) {
	return (strings.ToLower(strings.TrimSpace(s)))
}