package backend

import (
	"net/http"
	"text/template"
)

func Webhandler(write http.ResponseWriter, r *http.Request) {
	temp := template.Must(template.ParseFiles("index.html"))
	temp.Execute(write, nil)
}
