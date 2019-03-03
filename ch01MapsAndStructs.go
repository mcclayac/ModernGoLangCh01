package main

import "fmt"

func main() {

	//fmt.Printf("", )
	// key ==> value (key value pair)
	//var v map[string]int= make(map[string]int)
	x := make(map[string]int)
	x["first"]= 1
	x["second"]= 2
	x["third"]= 3

	y := map[string]int {
		"first":1,
		"second":2,
		"third":3,
		"fourth":4,
		"fifth":5,
	}

	delete(y,"fifth")
	fmt.Printf("x = %v\n", x)
	fmt.Printf("y = %v\n", y)


	for _, value := range x {
		fmt.Printf("value := %d\n", value)
	}

	_, ok := x["fourth"]
	fmt.Printf("fourth does this key exists ? :%v\n", ok)
	_, ok = x["first"]
	fmt.Printf("firstdoes this key exists ? :%v\n", ok)

	if v,ok :=x["first"]; ok {
		fmt.Printf("value = %v\n", v)
	}

	if v,ok :=x["fourth"]; ok {
		fmt.Printf("value = %v\n", v)
	} else {
		fmt.Printf("value = %v does not exists\n", v)
	}




}