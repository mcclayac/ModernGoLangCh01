package main

import (
	"fmt"
)

var x uint8 = 2

func main() {
	fmt.Println(CompareNumbers(1,2))
	//var x uint8 = 2
	fmt.Printf("%d\n\n", x)
	switch x := 2; x {       // can iniatlaize and assign
	case 3:
		fmt.Printf("I am %d\n", x)
	case 2:
		fmt.Printf("I am %d\n", x)
		fallthrough     // we do not want to break
	case 4:
		fmt.Printf("I am %d\n", x)
	}

	if i:=2 ; i < 4 {
		fmt.Printf("I=%d is lessthan 4\n", i)
	}
	// traditional for loop
	for i:=0 ; i < 10 ; i++ {
		fmt.Printf("Loop : %d\n", i)
	}

	// while loop
	i := 0
	for i < 10 {
		fmt.Printf("Loop value : %d\n", i)
		i++
	}

}


func CompareNumbers(i1 , i2 int) (bool, int) {
/*
	if i1 > i2 {
		fmt.Println("first number is greater than the second number")
		return false, i1 - i2
	} else if i2 > i1 {
		fmt.Println("Second number is greater than the first number")
		return false, i2-i1
	}
fmt.Println("The Numbers are equal")
*/

	switch {
	case i1 > i2 :
		fmt.Println("first number is greater than the second number")
		return false, i1 - i2
	case 12 > i1:
		fmt.Println("Second number is greater than the first number")
		return false, i2-i1
	default:
		fmt.Println("The Numbers are equal")
	}

	return true, 0
}






