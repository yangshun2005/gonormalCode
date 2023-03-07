package normalcodeblock

import (
	"strings"
	"unicode"
)

func CodeSplic(s string) []string {
	//GOTO
	// s := "We are humans. We are social animals."
	return strings.FieldsFunc(s, codeFunc)

}

func codeFunc(s rune) bool {
	// if s == ',' {
	// 	return true
	// }
	// return false

	return unicode.IsSpace(s) || s == '.'
}
