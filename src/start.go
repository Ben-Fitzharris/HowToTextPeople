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
			parseRequest(conn, incoming)
			conn.Close()
		}(connection)
	}

	//http handle server
	/*
			helloHandler := func(w http.ResponseWriter, req *http.Request) {
				io.WriteString(w, "Hello Internet Person \n")
			}

			fs := http.FileServer(http.Dir(path.Dir("./src")))

			http.HandleFunc("/hello", helloHandler)
			http.Handle("/", fs)
			log.Fatal(http.ListenAndServe("127.0.0.1:6112", nil))


		fmt.Println("end")
	*/
}
