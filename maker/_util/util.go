package _util

import "unicode"

func CapitalizeFirstChar(str string) string {
	if len(str) == 0 {
		return str // Return the original string if it's empty
	}

	firstChar := []rune(str)[0]
	upperFirstChar := unicode.ToUpper(firstChar)

	return string(upperFirstChar) + str[1:]
}
