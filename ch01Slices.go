package main

import "fmt"

func main() {

	var a [5]int = [5]int{1,2,3,4,5}
	mySlice := []int{6,7,8,9,10}  // can initialize []int{}
	mySlice = append(mySlice, 11,12,13,14)
	mySlice = append(mySlice, 15)
	mySlice2 := make([]byte,5)  // make([]byte,5,5)
	mySlice2 = append(mySlice2, 2,4,6,8)
	mySlice2[0] = 20
	mySlice2[1] = 22
	mySlice2[2] = 24
	mySlice2[3] = 26

	s := make([]int, 5)
	s[0],s[1],s[2],s[3],s[4] = 40,44,48,52,56   // quick assignments
	s1 := s[2:5]    ///  a part of a slice

	s2 := make([]int,5)
	copy(s2,s[1:3])




	fmt.Printf("a = %v\n", a)
	fmt.Printf("a = Capacity of a : %d\n", cap(a))
	fmt.Printf("a = Length of a : %d\n\n", len(a))
	fmt.Printf("s1 = %v\n", s1)
	fmt.Printf("s1 = Capacity of : %d\n", cap(s1))
	fmt.Printf("s1 = Length of : %d\n\n", len(s1))
	fmt.Printf("s2 = %v\n", s2)
	fmt.Printf("s2 = Capacity of : %d\n", cap(s2))
	fmt.Printf("s2 = Length of : %d\n\n", len(s2))
	fmt.Printf("myslice = %v\n", mySlice)
	fmt.Printf("myslice = Capacity of : %d\n", cap(mySlice))
	fmt.Printf("myslice = Length of : %d\n\n", len(mySlice))
	fmt.Printf("myslice2 = %v\n", mySlice2)
	fmt.Printf("myslice2 = Capacity of : %d\n", cap(mySlice2))
	fmt.Printf("myslice2 = Length of : %d\n\n", len(mySlice2))


	for _, value := range a {
		fmt.Printf("a = %d\n", value)
	}

	for key, value := range mySlice {
		fmt.Printf("mySlice =  key :%d    value: %d\n", key, value)
	}
	for key, value := range mySlice2 {
		fmt.Printf("MySlice2 =  key :%d    value: %d\n", key, value)
	}
	for key, value := range s {
		fmt.Printf("s = key :%d    value: %d\n", key, value)
	}
	for key, value := range s1 {
		fmt.Printf("s1 = key :%d    value: %d\n", key, value)
	}

}

/*
func gettwo( s []int, fi, si int) []int {
	s2 := make([]int, si-fi)
	copy(s2,s[fi,si])
	return s2
}
*/



