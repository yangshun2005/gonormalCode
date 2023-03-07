package osexepkg

import (
	"log"
	"os"
	"os/exec"
)

func OsexecPkgTwo(fn string) {

	c := exec.Command(fn)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	err := c.Run()
	if err != nil {
		log.Fatalf("shell error,%v", err.Error())
	}
}
