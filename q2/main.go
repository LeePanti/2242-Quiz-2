package main

import (
	"fmt"
)

func main() {
	connection := make(chan string)

	go independentChildConnected(connection)

	message := <-connection

	fmt.Println(message)

	//go disconnected()
}

func independentChildConnected(connection chan string) {
	connection <- "Hello From Outside Main"

	close(connection)
}

// func disconnected() {
// 	fmt.Println("Hello There")
// }
