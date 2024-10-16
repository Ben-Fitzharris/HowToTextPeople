package main

import "fmt"

func main() {
	fmt.Println("start")
	l := createListener()
	l.Close()
	fmt.Println("end")
}
