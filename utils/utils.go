package utils

import "strings"

func IsStrEqual(str1, str2 string) bool {
	return strings.EqualFold(strings.ToLower(str1), strings.ToLower(str2))
}

func IsArrEmpty(arr []string) bool {
	return len(arr) == 0
}
