package main

import (
	"html/template"
	"os"
)

type Friend struct {
	Fname string
}

type Person struct {
	UserName string
	Emails   []string
	Friends  []*Friend
}

func main() {
	f1 := Friend{Fname: "hsnice16"}
	f2 := Friend{Fname: "xushiwei"}
	t := template.New("fieldname example")
	t, _ = t.Parse(`hello {{.UserName}}!
	{{range .Emails}}
		an email {{.}}
	{{end}}
	{{with .Friends}}
	{{range .}}
		my friend name is {{.Fname}}
	{{end}}
	{{end}}`)

	p := Person{UserName: "Himanshu", Emails: []string{"hs@bee.me", "astax@example.com"}, Friends: []*Friend{&f1, &f2}}
	t.Execute(os.Stdout, p)
}
