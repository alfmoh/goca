package goca

import (
	"regexp"
	"strings"
)

var charSet = `[-_\\|]+`

func stringCase(str string, firstCharCase string, otherCharCase string) []string {
	slice := strings.Split(str, " ")
	var resSlice []string
	for _, val := range slice {
		var firstChar string
		var otherChars string

		if firstCharCase == "upper" {
			firstChar = strings.ToUpper(val[:1])
		} else {
			firstChar = strings.ToLower(val[:1])
		}

		if otherCharCase == "upper" {
			otherChars = strings.ToUpper(val[1:])
		} else {
			otherChars = strings.ToLower(val[1:])
		}

		newStr := firstChar + otherChars
		resSlice = append(resSlice, newStr)
	}
	return resSlice
}

func prepareString(str string) string {
	str = trim(str)
	str = regexp.MustCompile(charSet).ReplaceAllString(str, " ")
	return str
}

func prepareStringCharExclude(str string, strExclude ...string) string {
	if len(strExclude) > 0 {
		for _, strEx := range strExclude {
			if strings.Contains(charSet, strEx) {
				charSet = strings.ReplaceAll(charSet, strEx, "")
			}
		}
	}
	return prepareString(str)
}

func trim(str string) string {
	return strings.Trim(str, " ")
}

func strLen(str string) int {
	result := 0
	for range str {
		result++
	}
	return result
}
