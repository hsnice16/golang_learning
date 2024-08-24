package main

import (
	"html/template"
	"os"
)

type Person struct {
	UserName string
	email    string // Field in not exported.
}

func main() {
	t := template.New("fieldname exmaple")
	// {{.}} to output the object in formatted string
	t, _ = t.Parse("{{.}}- hello {{.UserName}}! {{.email}}") // email not exported, so will have empty("") string
	p := Person{UserName: "Himanshu", email: "h@gmai.com"}
	t.Execute(os.Stdout, p)
}
