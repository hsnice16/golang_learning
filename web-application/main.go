package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", sayHelloName) // set router

	err := http.ListenAndServe(":9090", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}


	/// ------- Custom Router ------------
	// mux := &custom_router.MyMux{}
	// http.ListenAndServe(":9090", mux)
}

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	fmt.Println("--------------- Request: PROCESSING ---------------")

	r.ParseForm()       // parse arguments
	fmt.Println(r.Form) // print form information in server side
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])

	for k, v := range r.Form {
		fmt.Println("key: ", k)
		fmt.Println("val: ", strings.Join(v, ", "))
	}

	fmt.Fprintf(w, "Hello hsnice16!") // send data to client side
	fmt.Println("--------------- Request: DONE ---------------")
}
