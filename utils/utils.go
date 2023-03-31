package utils

import (
	"math/rand"
	"strconv"
	"strings"
)

func IsStrEqual(str1, str2 string) bool {
	return strings.EqualFold(strings.ToLower(str1), strings.ToLower(str2))
}

func IsArrEmpty(arr []string) bool {
	return len(arr) == 0
}

func Generate(n int) string {
	var chars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0987654321")
	str := make([]rune, n)
	for i := range str {
		str[i] = chars[rand.Intn(len(chars))]
	}
	x := rand.Intn(3)

	return string(str) + strconv.Itoa(x)
}
