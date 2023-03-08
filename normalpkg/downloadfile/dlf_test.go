package downloadfile

import (
	"fmt"
	"testing"
)

func TestDLF(t *testing.T) {
	url := "http://dl.chinaase.com/softwarepackage/k8s/script/t1.sh"

	filename := "test1.sh"
	DownloadFiled(url, filename)

	fmt.Println("hi")
}
