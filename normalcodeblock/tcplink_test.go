package normalcodeblock

import (
	"fmt"
	"testing"
)

func TestTcplink(t *testing.T) {
	tcpData := "hi"
	ipaddr := "233.5.5.5"
	str1 := TcpLink(tcpData, ipaddr)
	fmt.Println(str1)
}
