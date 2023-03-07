package main

import (
	"strings"
)

func main() {
	//GOTO
	strings.FieldsFunc(TOPIC["data"], CodeSplit)
}

func CodeSplit(s rune) bool {
	if s == ',' {
		return true
	}
	return false
}
