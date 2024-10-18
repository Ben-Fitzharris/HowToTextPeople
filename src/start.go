package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	//messages := make(chan string)

	fmt.Println("start")
	l := createListener()
	defer l.Close()

	/*
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
	*/
	//http handle server

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello Internet Person \n")
	}

	http.HandleFunc("/hello", helloHandler)
	log.Fatal(http.ListenAndServe("127.0.0.1:6112", nil))

	fmt.Println("end")
}
