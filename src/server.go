package main

import "net"

func createListener() (l net.Listener) {
	l, err := net.Listen("tcp4", "locallly")
	if err != nil {
		println("Create Listener Error: ", err)
	}

	l.Accept()

	return l
}
