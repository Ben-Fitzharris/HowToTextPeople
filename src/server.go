package main

import "net"

func createListener() (l net.Listener) {
	l, err := net.Listen("tcp4", "127.0.0.1:6111")
	if err != nil {
		println("Create Listener Error: ", err)
	}

	l.Accept()

	return l
}

func parseRequest(b []byte) {

}
