package main

import (
	"html/template"
	"log"
	"net/http"
)

var testTemplate *template.Template

type ViewData struct {
	Name 	string
	Widgets []Widget
}

type Widget struct {
	Name string
	Price int
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
	vd := ViewData{
		"John Smith",
		[]Widget{
			{"Blue Widget", 12},
			{"Red Widget", 12},
			{"Green Widget", 12},
		},
	}
	err := testTemplate.Execute(writer, vd)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}