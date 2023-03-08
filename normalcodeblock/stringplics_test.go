package normalcodeblock

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestStrSpl(t *testing.T) {
	s := "We are humans. We are social animals."
	strarr := CodeSplic(s)
	fmt.Println(strarr)
	fmt.Println(reflect.TypeOf(strarr))
}

func TestStrEqual(t *testing.T) {
	s := "We are humans."
	strarr1 := CodeSplic(s)
	var strarr2 []string
	strarr2 = append(strarr2, "We")
	strarr2 = append(strarr2, "are")
	strarr2 = append(strarr2, "humans")
	assert.Equal(t, strarr1, strarr2, "")
}
