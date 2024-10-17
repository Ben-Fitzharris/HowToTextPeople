package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	//messages := make(chan string)

	fmt.Println("start")
	l := createListener()
	defer l.Close()

	for {

		connection, err := l.Accept()
		if err != nil {
			log.Fatal("Listener Error: ", err)
		}

		go func(conn net.Conn) {
			incoming := make([]byte, 1000)
			num, err := conn.Read(incoming)
			if err != nil {
				fmt.Println("Read Error: ", err)
			}
			fmt.Println("num: ", num)
			strIncoming := string(incoming)
			fmt.Println(strIncoming)

			conn.Close()
		}(connection)
	}

	fmt.Println("end")
}
