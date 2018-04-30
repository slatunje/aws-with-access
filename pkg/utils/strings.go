package utils

import "strings"

func ToUpper(s []string) (keys []string) {
	for _, k := range s {
		keys = append(keys, strings.ToUpper(k))
	}
	return
}
