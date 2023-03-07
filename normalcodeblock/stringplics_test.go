package normalcodeblock

import (
	"fmt"
	"testing"
)

func TestStrSpl(t *testing.T) {
	s := "We are humans. We are social animals."
	strarr := CodeSplic(s)
	fmt.Println(strarr)
}
