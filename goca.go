package goca

import (
	"encoding/hex"
	"math"
	"strings"
)

// CamelCase converts a string into camelCase.
func CamelCase(str string) string {

	str = prepareString(str)

	if str == "" || strings.Index(str, " ") < 0 {
		return str
	}
	slice := strings.Split(str, " ")
	var resSlice []string

	for i, val := range slice {
		if i != 0 {
			newStr := strings.ToUpper(val[:1]) + strings.ToLower(val[1:])
			resSlice = append(resSlice, newStr)
		} else {
			resSlice = append(resSlice, strings.ToLower(val))
		}
	}
	return strings.Join(resSlice[:], "")
}

// Capitalize the first letters of a string to Upper case. Rest of the strings are in Lower case
func Capitalize(str string) string {
	str = prepareString(str)
	if str == "" {
		return str
	}
	resSlice := stringCase(str, "upper", "lower")
	return strings.Join(resSlice[:], " ")
}

// Decapitalize the first letters of a string. Rest of the strings are in Upper case
func Decapitalize(str string) string {
	str = prepareString(str)
	if str == "" {
		return str
	}
	resSlice := stringCase(str, "lower", "upper")
	return strings.Join(resSlice[:], " ")
}

// KebabCase converts the string to kebab-case
func KebabCase(str string, separator ...string) string {
	str = prepareString(str)
	if str == "" {
		return str
	}
	sep := "-"
	if len(separator) > 0 {
		sep = separator[0]
	}
	resSlice := stringCase(str, "lower", "lower")
	return strings.Join(resSlice[:], sep)
}

// SwapCase swaps case in a string for the opposite case
func SwapCase(str string) string {
	str = strings.Trim(str, " ")
	if str == "" {
		return str
	}
	res := ""
	for _, char := range str {
		charStr := string(char)
		if strings.ToUpper(charStr) == charStr {
			res += strings.ToLower(charStr)
		} else if strings.ToLower(charStr) == charStr {
			res += strings.ToUpper(charStr)
		} else {
			res += charStr
		}
	}
	return res
}

// TitleCase converts the string expect the specified noSplit string to Title case
func TitleCase(str string, noSplit ...string) string {
	if str == "" {
		return str
	}
	if len(noSplit) < 1 {
		return Capitalize(str)
	}
	str = prepareStringCharExclude(str, noSplit...)
	resSlice := stringCase(str, "upper", "lower")
	return strings.Join(resSlice[:], " ")
}

// Chain converts a strong using the passed in functions
// func Chain(str string, funcs ...func(s string) string) string {
// 	res := str
// 	if str == "" || strLen(funcs) < 1 {
// 		return str
// 	}
// 	for _, fun := range funcs {
// 		res = fun(res)
// 	}
// 	return res
// }

// CharAt returns a character from a string using it's index
func CharAt(str string, index uint64) string {
	if trim(str) == "" || int(index) > strLen(str)-1 {
		return str
	}
	return string([]rune(str)[index])
}

// HexAt returns a hex value of the string
func HexAt(str string) string {
	return hex.EncodeToString([]byte(str))
}

// First returns the specified number of characters from the beginning of a string
func First(str string, charNum uint64) string {
	if trim(str) == "" || int(charNum) > strLen(str) {
		return str
	}

	return string([]rune(str)[:charNum])
}

// Last returns the specified number of characters from the end of a string
func Last(str string, charNum uint64) string {
	if trim(str) == "" || int(charNum) > strLen(str) {
		return str
	}

	return string([]rune(str)[strLen(str)-int(charNum):])
}

// Prune truncates a string and appends using a specified length and returns a string not greater than the specified length with three dots or a specified string appended to the end
func Prune(str string, length uint64, end ...string) string {
	str = prepareString(str)
	totalWordsNum := 0
	ending := "..."
	if len(end) > 0 {
		ending = end[0]
	}

	if str == "" || length == 0 {
		return str
	}

	if strings.Index(str, " ") < 0 {
		return str + ending
	}

	slice := strings.Split(str, " ")
	var resSlice []string

	for _, val := range slice {
		totalWordsNum += len(val)
		if totalWordsNum < int(length) {
			resSlice = append(resSlice, val)
		} else {
			break
		}
	}
	return strings.Trim(strings.Join(resSlice, " "), " ") + ending
}

// Slice returns a part of a string using the provided 'start' and 'end' index. The 'end' index string character is excluded
func Slice(str string, start float64, end ...uint64) string {
	startAbs := math.Abs(start)
	endVal := end[0]
	if int(endVal) > strLen(str) {
		endVal = uint64(strLen(str))
	}
	if startAbs > float64(strLen(str)) {
		return str
	}
	if start < 0 {
		return Last(str, uint64(startAbs))
	}
	if start > -1 && endVal == 0 {
		return Last(str, uint64(strLen(str)-int(start)))
	}
	return string([]rune(str)[int(start):int(endVal)])
}

// Count returns the number of characters in a string
func Count(str string) int {
	return strLen(str)
}

// CountSubStrings returns the number of specified substrings in the string
func CountSubStrings(str string, substr string) int {
	if substr == "" {
		return 0
	}
	counter := 0
	subStringLength := strLen(substr)
	stringLength := strLen(str)
	firstSubstrChar := string([]rune(substr)[0:1])
	strArr := strings.Split(str, "")
	for i := 0; i < len(strArr); i++ {
		chaStr := string(strArr[i])
		if chaStr != firstSubstrChar {
			continue
		}
		lastCharIndex := subStringLength + i
		if lastCharIndex > stringLength {
			lastCharIndex = stringLength
		}
		if strings.Join(strArr[i:lastCharIndex], "") == substr {
			counter++
		}
	}
	return counter
}

// CountWhere ....
// func CountWhere(str string, f func(funcString string, params ...string)) string {

// }
