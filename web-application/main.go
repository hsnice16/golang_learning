package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", sayHelloName) // set router
	http.HandleFunc("/login", login)

	err := http.ListenAndServe(":9090", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

	/// ------- Custom Router ------------
	// mux := &custom_router.MyMux{}
	// http.ListenAndServe(":9090", mux)
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("---------------- Request %v: PROCESSING ----------------\n", r.Method) // get request method

	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		// logic part of log in
		fmt.Println("username:", r.Form["username"])

		if len(r.Form["username"][0]) == 0 {
			fmt.Println("Empty username")
		}

		fmt.Println("Escaped username:", template.HTMLEscapeString(r.Form.Get("username")))
		// fmt.Fprint(w, r.Form.Get("username"))
		template.HTMLEscape(w, []byte(r.Form.Get("username")))

		fmt.Println("password:", r.Form["password"])
		fmt.Println("fruit:", r.Form["fruit"])

		slice := []string{"apple", "pear", "banana"}
		isValid := false
		for _, v := range slice {
			if v == r.Form.Get("fruit") {
				isValid = true
				fmt.Println("Valid select value")
			}
		}
		if !isValid {
			fmt.Println("Invalid select value")
		}

		fmt.Println("gender:", r.Form["gender"])

		slice = []string{"1", "2"}
		isValid = false
		for _, v := range slice {
			if v == r.Form.Get("gender") {
				isValid = true
				fmt.Println("Valid gender value")
			}
		}
		if !isValid {
			fmt.Println("Invalid gender value")
		}

		fmt.Println("interest:", r.Form["interest"])

		slice = []string{"football", "basketball", "tennis"}
		a := Slice_diff(r.Form["interest"], slice)
		fmt.Println("a:", a)

		if a == nil {
			fmt.Println("Valid interest value")
		} else {
			fmt.Println("Invalid interest value")
		}
	}

	fmt.Printf("---------------- Request %v: DONE ----------------\n", r.Method)
}

func In_slice[T string](val T, slice []T) bool {
	for _, v := range slice {
		if v == val {
			return true
		}
	}

	return false
}

func Slice_diff[T string](slice1, slice2 []T) (diffslice []T) {
	for _, v := range slice1 {
		if !In_slice(v, slice2) {
			diffslice = append(diffslice, v)
		}
	}

	return
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
