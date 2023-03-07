package osexepkg

import "testing"

func TestOsexecPkg(t *testing.T) {
	url := "http://dl.chinaase.com/softwarepackage/k8s/script/t1.sh"

	downloadfile.DownloadFiled(url)

	fn := ".e/test1.sh"
	OsexecPkgTwo(fn)

}
