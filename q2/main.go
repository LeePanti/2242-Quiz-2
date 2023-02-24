// channels in Go
// Lee Panti

package main

import (
	"fmt"
)

func main() {
	// channels connect the independent go routines to main via a "pipe".
	connection := make(chan string)

	go independentChildConnected(connection)

	message := <-connection

	fmt.Println(message)

	// go disconnected()
	// fmt.Println("finished")
}

func independentChildConnected(connection chan string) {
	connection <- "Hello From Outside Main"

	close(connection)
}

// func disconnected() {
// 	fmt.Println("Hello There")
// }
