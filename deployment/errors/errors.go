package errors

import (
	"net/http"
	"text/template"

	"github.com/hsnice16/golang_learning/deployment/logs"
)

type MyMux string // custom type for demo

func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		sayhelloName(w, r)
		return
	}
	NotFound404(w, r)
	return
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {

}

func NotFound404(w http.ResponseWriter, r *http.Request) {
	logs.Logger.Critical(" page not found ")     // error logging
	t, _ := template.ParseFiles("tmpl/404.html") // parse the template file
	ErrorInfo := " File not found "              // Get the current user information
	t.Execute(w, ErrorInfo)                      // execute the template merger operation
}

func SystemError(w http.ResponseWriter, r *http.Request) {
	logs.Logger.Critical(" System Error")              // system error triggered Critical, then logging will not only send a message
	t, _ := template.ParseFiles("tmpl/error.html")     // parse the template file
	ErrorInfo := " system is temporarily unavailable " // Get the current user information
	t.Execute(w, ErrorInfo)                            // execute the template merger operation
}
