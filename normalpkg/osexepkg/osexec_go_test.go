package osexepkg

import (
	"github.com/yangshun2005/gonormalCode/normalpkg/downloadfile"
	"testing"
)

func TestOsexecPkg(t *testing.T) {
	url := "http://dl.chinaase.com/softwarepackage/k8s/script/t1.sh"

	DownloadFiled(url)

	fn := "./test1.sh"
	OsexecPkgTwo(fn)

}
