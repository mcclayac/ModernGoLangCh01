package main

import (
	"fmt"
)

func testEIface(v interface{}) {
	// type assertion
	if i, ok := v.(int); ok {
		fmt.Printf("I am an interger : %d\n", i)
	} else if s, ok := v.(string); ok {
		fmt.Printf("I am a string : %s\n", s)
	} else {
		fmt.Printf("my value is : %v\n", v)
	}
}

func main() {
	fmt.Println("Hello, playground")
	v := fmt.Sprintf("%[2]d %[1]d\n", 11, 22)
	//fmt.Println(v)
	testEIface(v)
	// type switch

	v = fmt.Sprintf("%6.2f", 12.0)
	testEIface(v)

	var i int
	var s string
	fmt.Sscanf(" 1234567 ", "%5s%d", &s, &i)

	testEIface(s)
	testEIface(i)

	integer := 23
	fmt.Printf("%T %T\n", integer, &integer)

	answer := 42
	binaryS := fmt.Sprintf("%b\n", answer)
	// Result: 42 42 2a 52 101010
	testEIface(binaryS)

	//smile  := '😀'

	person := struct {
		Name string
		Age  int
	}{"Kim", 22}
	fmt.Printf("%v %+v %#v\n", person, person, person)
	fmt.Printf("%#v\n", person)
	fmt.Printf("%+v\n", person)


}