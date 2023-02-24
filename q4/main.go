package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(2)

	go deliveryBoy1(&wg)
	go deliveryBoy2(&wg)

	wg.Wait()
}

func deliveryBoy1(wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Went to do an errand")
	time.Sleep(time.Second)
	fmt.Println("finished with my errand")
}

func deliveryBoy2(wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Went to do a delivery")
	time.Sleep(time.Second * 2)
	fmt.Println("finished with my delivery")
}
