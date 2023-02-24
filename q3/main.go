package main

import (
	"fmt"
	"time"
)

func main() {
	go deliveryBoy()

	time.Sleep(time.Second * 2)
	fmt.Println("closing the Shop")
}

func deliveryBoy() {
	fmt.Println("started my errand")
	time.Sleep(time.Second * 1)
	fmt.Println("finished my errand")
}
