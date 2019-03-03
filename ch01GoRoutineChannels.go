package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	quitsignal := make(chan bool)

	go SayHelloFromGoroutine(quitsignal)

	go func() {
		fmt.Println("Hello, from Anomomous Function goroutine 1")
		fmt.Println("Hello, from Anomomous Function goroutine 2")
		fmt.Println("Hello, from Anomomous Function goroutine 3")
		fmt.Println("Hello, from Anomomous Function goroutine 4")
		quitsignal <- true // send a True to the boolean channel
	}()

	fmt.Println("Hello World")

	v := <-quitsignal // wait on channel for incoming value
	fmt.Println(v)

	v = <-quitsignal // wait on channel for incoming value
	fmt.Println(v)

	//time.Sleep(time.Second * 3)

	ic := make(chan int)
	go periodicSend(ic)

	for i := range ic { //  blocking call
		fmt.Printf("from ic : %d\n", i)
	}

		// check to see if the channel is open
	_, ok := <- ic
	fmt.Println(ok)

}

func SayHelloFromGoroutine(qs chan bool) {
	fmt.Println("Hello, from goroutine 1")
	fmt.Println("Hello, from goroutine 2")
	fmt.Println("Hello, from goroutine 3")
	fmt.Println("Hello, from goroutine 4")
	qs <- false // send a True to the boolean channel
}

func periodicSend(ic chan int) {
	rand.Seed(time.Now().UTC().UnixNano())    // random number seed
	for i := 0; i < 3; i++ {
		//ic <- 3
		ic <- rand.Intn(100)
		time.Sleep(1 * time.Second)
	}
	close(ic) // closes a channel  the range will stop
}
