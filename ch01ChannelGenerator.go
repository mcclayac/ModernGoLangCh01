package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("------------------------")
	fmt.Println("Channel Generator")
	fmt.Println("------------------------")
	//fmt.Println(<-waitAndSend(3, 5))

	ic := make(chan int)
	select {
	case v1 := <-waitAndSend(3, 5):
		fmt.Println(v1)
	case v2 := <-waitAndSend(5, 10):
		fmt.Println(v2)
	case ic <- 25:
		fmt.Println("ic recieved a value")
	/*default:
		fmt.Println("All Channes are slow")*/
	}

	//fmt.Println(<-ic)

}

// go channel generator
func waitAndSend(v, i int) chan int { // will wait for i seconds before sending value v on the return channel

	retCh := make(chan int)
	go func() {
		time.Sleep(time.Duration(i) * time.Second)
		retCh <- v
	}()
	return retCh
}
