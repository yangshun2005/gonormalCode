package osexepkg

import (
	"log"
	"os"
	"os/exec"
)

func OsexecPkgRm(filename string) {
	// abpath := os.Getwd()
	// file, err := exec.LookPath(os.Args[0])
	// if err != nil {
	// 	fmt.Println("err")
	// }
	// fmt.Println(filepath.Abs(file))

	c := exec.Command("rm -f " + "/root/svn_git/mycode/projectcode/gonormalCode/normalpkg/osexepkg/" + filename)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	err := c.Run()
	if err != nil {
		log.Fatalf("shell error,%v", err.Error())
	}
}
