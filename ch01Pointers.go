package main

import "fmt"

func main() {
	var pI *int // memory address of value of ==> of value of type int
	var I int = 3
	pI = &I

	increment(pI)
	fmt.Println(I)

	increment(&I)
	fmt.Println(I)

}

func increment(pI *int) { // pI = Pointer to I
	*pI++ // derefference
}
