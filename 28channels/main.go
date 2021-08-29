package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in golang-  LearnCodeOnline.in")

	myCh := make(chan int, 2)
	wg := &sync.WaitGroup{}

	// fmt.Println(<-myCh)
	// myCh <- 5
	wg.Add(2)
	// R ONLY
	go func(ch <-chan int, wg *sync.WaitGroup) {

		val, isChanelOpen := <-myCh

		fmt.Println(isChanelOpen)
		fmt.Println(val)

		//fmt.Println(<-myCh)

		wg.Done()
	}(myCh, wg)
	// send ONLY
	go func(ch chan<- int, wg *sync.WaitGroup) {
		myCh <- 0
		close(myCh)
		// myCh <- 6
		wg.Done()
	}(myCh, wg)

	wg.Wait()
}
