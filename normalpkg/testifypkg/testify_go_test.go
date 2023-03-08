package testifypkg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEqual(t *testing.T) {
	var a = 100
	var b = 200
	assert.Equal(t, a, b, "")
}
