package misc

import (
	"encoding/base64"
	"math/rand"
	"strings"
)

func ShuffleRuneArray(input []rune) {
	length := len(input)
	for index, _ := range input {
		swapIndex := rand.Intn(length)
		temp := input[index]
		input[index] = input[swapIndex]
		input[swapIndex] = temp
	}
}

func ObfuscateString(s string) string {
	a := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789+-*/=")
	b := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789+-*/=")
	ShuffleRuneArray(a)
	ShuffleRuneArray(b)
	input := []rune(s)
	result := ""
	for _, r := range input {
		aIndex := strings.IndexRune(string(a), r)
		if aIndex < 0 {
			result += string(r)
		} else {
			result += string(b[aIndex])
		}
	}
	return result + string(a) + string(b)
}

func DecodeObfuscatedString(s string) string {
	length := len(s)
	strA := s[length-67*2 : length-67]
	strB := s[length-67 : length]
	str := s[0 : length-67*2]
	result := ""
	for i := 0; i < len(str); i++ {
		bIndex := strings.Index(strB, s[i:i+1])
		if bIndex < 0 {
			result += s[i : i+1]
		} else {
			result += strA[bIndex : bIndex+1]
		}
	}
	return result
}

const base64Prefix = "JSHEBFKS"
const base64Suffix = "LY"

func Base64RandEncode(src string) string {
	dist := base64.URLEncoding.EncodeToString([]byte(src))
	l := len(dist)
	return base64Prefix + dist[:l-4] + base64Suffix + dist[l-4:]
}

func Base64RandDecode(src string) string {
	l := len(src)
	if l < 10 {
		return ""
	}
	dist := src[8:l-6] + src[l-4:]
	str, _ := base64.URLEncoding.DecodeString(dist)
	return string(str)
}
