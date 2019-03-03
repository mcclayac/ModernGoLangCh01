package main

import "fmt"

func main() {
	fmt.Println("Hello, playground")
	fmt.Printf("Hello %s \n", "World")
	s := struct {
		i int
		f float32
	}{i:32, f:3.33}
	fmt.Printf("%+v\n", s)
	fmt.Printf("i is %d, while f id %.2f\n", s.i, s.f)
	fmt.Println("Hello again")


}