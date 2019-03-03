package main

func main() {
	buffch := make(chan int, 5)   // 5 = buffer size

	buffch <- 3
	buffch <- 2
	buffch <- 1
	buffch <- 0
	println(<-buffch) /// FIFO
	println(<-buffch) /// FIFO
	println(<-buffch) /// FIFO
	println(<-buffch) /// FIFO

}

