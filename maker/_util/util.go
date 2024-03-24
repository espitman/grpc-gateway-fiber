package _util

import (
	"time"
	"unicode"
)

func Time() string {
	currentTime := time.Now()
	currentTimeString := currentTime.Format("2006-01-02 15:04:05")
	return currentTimeString
}

func CapitalizeFirstChar(str string) string {
	if len(str) == 0 {
		return str // Return the original string if it's empty
	}

	firstChar := []rune(str)[0]
	upperFirstChar := unicode.ToUpper(firstChar)

	return string(upperFirstChar) + str[1:]
}
