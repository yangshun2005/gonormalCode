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
	str := Client(conn, tcpData)
	return str
}

func Client(conn net.Conn, sms string) string {
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
