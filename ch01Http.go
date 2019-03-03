package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)
func testEIface(v interface{}) {
	// type assertion
	fmt.Printf("%T\n", v)
	if i, ok := v.(int); ok {
		fmt.Printf("I am an interger : %d\n", i)
	} else if s, ok := v.(string); ok {
		fmt.Printf("I am a string : %s\n", s)
	} else {
		fmt.Printf("my value is : %v\n", v)
	}
}

func main() {

	resp, err := http.Get("http://example.com/")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	testEIface(body)
	//println(resp)


	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n These Nutz")
	}

	//http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	//})

	http.HandleFunc("/hello", helloHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
