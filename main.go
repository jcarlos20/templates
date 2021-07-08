package main

import (
	"html/template"
	"log"
	"net/http"
)

var testTemplate *template.Template

type ViewData struct {
	Name string
}

func main() {
	var err error
	testTemplate, err = template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func handler(writer http.ResponseWriter, request *http.Request)  {
	writer.Header().Set("Content-Type", "text/html")
	vd := ViewData{"John Smith"}
	err := testTemplate.Execute(writer, vd)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}