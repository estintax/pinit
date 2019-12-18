package main

import (
	"net"
	"strings"
)

func StartServer(addr string) {
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		Error("failed to start socket server", err)
		return
	}

	for {
		conn, _ := ln.Accept()
		go ConnectionHandler(conn)
	}
}

func ConnectionHandler(conn net.Conn) {
	conn.Write([]byte("INFO: pinit services controller\n"))
	for {
		var data string
		bytes := make([]byte, 1024)
		length, err := conn.Read(bytes)
		if err != nil {
			conn.Close()
			return
		}

		data = string(bytes[:length])
		data = strings.Replace(data, "\n", "", -1)
		data = strings.Replace(data, "\r", "", -1)
		params := strings.Split(data, " ")
		switch params[0] {
		case "START":
			if len(params) > 1 {
				proc := StartService(params[1], false)
				if proc != nil {
					conn.Write([]byte("START SUCCESS\n"))
				} else {
					conn.Write([]byte("START FAIL\n"))
				}
			}
		case "PING":
			conn.Write([]byte("PONG\n"))
		}
	}
}
