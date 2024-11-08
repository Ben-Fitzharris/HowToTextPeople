package main

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"strings"
)

func createListener() (l net.Listener) {
	l, err := net.Listen("tcp4", "127.0.0.1:6111")
	if err != nil {
		println("Create Listener Error: ", err)
	}

	l.Accept()

	return l
}

func parseRequest(conn net.Conn, b []byte) {

	num, err := conn.Read(b)
	if err != nil {
		fmt.Println("Read Error: ", err)
	}
	fmt.Println(num)

	r, err := http.ReadRequest(bufio.NewReader(strings.NewReader(string(b))))
	if err != nil {
		fmt.Println("Error http request parse: ", err)
	}

	fmt.Println(r.Method)
	fmt.Println(r.RequestURI)

	if r.Method == "GET" {
		//getResource(r.RequestURI)
		//respond
	}

}
