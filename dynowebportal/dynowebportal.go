package dynowebportal

import (
	"fmt"
	"net/http"
)

// RunWebPortal starts running the Dino Webportal on address addr
func RunWebPortal(addr string) error {

	http.HandleFunc("/", roothandler)
	err := http.ListenAndServe(addr, nil) // this is a blocking call
	return err
}

func roothandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Dino web portal - These Nutz %s", r.RemoteAddr)
}
