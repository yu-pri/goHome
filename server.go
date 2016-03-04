package main

import (
	"html/template"
	"fmt"
	"io/ioutil"
  "net/http"
)

type Page struct {
    Title string
    Body  []byte
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func loadPage(name string) (*Page, error) {
    filename := title + ""
    body, err := ioutil.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    return &Page{Title: title, Body: body}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
    path := r.URL.Path[len("/v/"):]
    p, _ := loadPage(path)
    t, _ := template.ParseFiles("index")
    t.Execute(w, p)
}


func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
    t, _ := template.ParseFiles(tmpl + ".html")
    t.Execute(w, p)
}

func main() {
    http.HandleFunc("/", handler)
		http.HandleFunc("/view/", viewHandler)
    http.ListenAndServe(":8080", nil)
}
