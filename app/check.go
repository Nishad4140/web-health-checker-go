package app

import (
	"fmt"
	"net"
	"time"
)

func check(dest string, port string) string {
	addr := dest + ":" + port
	timeout := time.Duration(5 * time.Second)

	conn, err := net.DialTimeout("tcp", addr, timeout)
	var status string
	if err != nil {
		status = fmt.Sprintf("[DOWN] %v is unreachable,\n Error: %v", dest, err)
	} else {
		status = fmt.Sprintf("[UP] %v is reachable,\n From: %v\n To: %v", dest, conn.LocalAddr(), conn.RemoteAddr())
	}

	return status
}
