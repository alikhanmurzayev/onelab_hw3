package reverse

import (
	"unicode/utf8"
)

//Reverse returns reversed string, works for different types of alphabets
func Reverse(str string) string {
	var bytes, res = []rune(str), []rune(str)
	var j = utf8.RuneCountInString(str) - 1
	for i, _ := range bytes {
		res[i] = bytes[j]
		j--
	}
	return string(res)
}

