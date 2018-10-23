package timeservice

import (
	"fmt"
	"net"
)

func TcpClientRun(host, port string) (int64, error) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		return 0, err
	}
	buf := make([]byte, 4)
	_, err = conn.Read(buf)
	if err != nil {
		return 0, err
	}
	return decodeTime(buf).Unix(), nil

}
