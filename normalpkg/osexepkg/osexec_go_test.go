package osexepkg

import (
	"github.com/yangshun2005/gonormalCode/normalpkg/downloadfile"
	"testing"
)

func TestOsexecPkg(t *testing.T) {
	url := "http://dl.chinaase.com/softwarepackage/k8s/script/t1.sh"
	fn := "./test1.sh"
	downloadfile.DownloadFiled(url, fn)

	OsexecPkgTwo(fn)

}
