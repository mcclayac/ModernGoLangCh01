package main

import "fmt"

type person struct {
	Name string
	Age int
	Address string
}


func main() {

	jason := person{
		Name:"Jason S.",
		Age: 38,
		Address: "Germany",
	}

	fmt.Println(jason)
	fmt.Println(jason.Name)
}