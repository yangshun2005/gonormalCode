package normalcodeblock

import (
	"fmt"
	"testing"
)

func TestTcplink(t *testing.T) {
	tcpData := "hi"
	ipaddr := "233.5.5.5"
	str := TcpLink(tcpData, ipaddr)
	fmt.Println(str)
}
