package main

import (
	"fmt"
	"net"
	"strings"
)

func StartServer(addr string) {
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println("pinit: Error: failed to start socket server\nMore: " + err.Error())
		return
	}

	for {
		conn, _ := ln.Accept()
		go ConnectionHandler(conn)
	}
}

func ConnectionHandler(conn net.Conn) {
	conn.Write([]byte("INFO: pinit services controller"))
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
					conn.Write([]byte("START SUCCESS"))
				} else {
					conn.Write([]byte("START FAIL"))
				}
			}
		case "PING":
			conn.Write([]byte("PONG"))
		}
	}
}
