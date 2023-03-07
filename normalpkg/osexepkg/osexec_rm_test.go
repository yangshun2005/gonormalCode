package osexepkg

import "testing"

func TestOsexeRm(t *testing.T) {
	fn := "test1.sh"
	OsexecPkgRm(fn)
}
