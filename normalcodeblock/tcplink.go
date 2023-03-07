package normalcodeblock

import (
	"fmt"
	"net"
)

func TcpLink(tcpData string, ipaddr string) (str string) {
	conn, err := net.Dial("tcp", ipaddr)
	if err != nil {
		fmt.Println("connect fail")
		return
	}
	defer conn.Close()
	// tcpData := "hi"
	sms := tcpData
	str3 := Client(conn, sms)
	return str3
}

func Client(conn net.Conn, sms string) (str2 string) {
	fmt.Println("want to send message:" + sms)
	conn.Write([]byte(sms))
	buf := make([]byte, 2)
	n, err := conn.Read((buf))
	if err != nil {
		fmt.Println("server read sms excapition")
	}
	fmt.Println("server respon :" + string(buf[0:n]))

	return string(buf[0:n])
}
