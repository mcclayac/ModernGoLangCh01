package main

import "fmt"

type testiface interface {
	SayHello()
	Say(string)
	increment()
	GetInternalValue() int
}

type testConcreteImpl struct{
	i int
}

func NewTestConcreteImpl() testiface  {  // this is our Constructor

	testiface := new(testConcreteImpl)
	testiface.i = 30
	return testiface
}

func NewTestConcreteImplWithValue(v int) testiface  { // this is our Constructor
/*
	testiface := new(testConcreteImpl)
	testiface.i = v
	return testiface
	*/

	return &testConcreteImpl{ i:v}
}



func (tst *testConcreteImpl) SayHello() {
	fmt.Println("Hello")
}

func (tst *testConcreteImpl) Say(s string) {
	fmt.Printf("%s = Say()", s)
}

func (tst *testConcreteImpl) increment() {
	tst.i++
}

func (tst *testConcreteImpl) GetInternalValue() int {
	return tst.i
}


type testEmbedding struct {  // we want this struct to have all the features of *testConcreteImpl
	*testConcreteImpl //embedding  -  The Inner Type
}

func testEIface(v interface{}) {

	// type assertion
	if i, ok := v.(int); ok {
		fmt.Printf("I am an interger : %d\n", i)
	} else if s, ok := v.(string); ok {
		fmt.Printf("I am a string : %s\n", s)
	} else {
		fmt.Printf("my value is : %v\n", v)
	}

	// type switch
	switch val := v.(type) {
	case int:
		fmt.Printf("I am an interger : %d\n", val)
	case string:
		fmt.Printf("I am a string : %s\n", val)
	default:
		fmt.Printf("my value is : %v\n", val)
	}
	fmt.Printf("\n")
	//fmt.Println(v)
}




func main() {

	fmt.Println("Hello, These nutz")

	var tiface testiface
	var tiface2 testiface
	//tiface = &testConcreteImpl{}
	//tiface = new(testConcreteImpl)
	tiface = NewTestConcreteImpl()
	tiface2 = NewTestConcreteImplWithValue(100)


	tiface.Say("Hi Mom")
	tiface.SayHello()
	tiface.increment()
	tiface.increment()
	tiface.increment()
	fmt.Println(tiface.GetInternalValue())

	tiface2.Say("Hi These nutz")
	tiface2.SayHello()
	tiface2.increment()
	tiface2.increment()
	tiface2.increment()
	fmt.Println(tiface2.GetInternalValue())

	te := testEmbedding{
		&testConcreteImpl{ i:50}}
	te.SayHello()
	te.Say("Your Mama")
	te.increment()
	te.increment()
	te.increment()
	te.increment()
	fmt.Println(te.GetInternalValue())

	// empty Interface example
	testEIface(3)
	testEIface("Hi Mom")
	testEIface(tiface)
	testEIface(tiface2)

}
