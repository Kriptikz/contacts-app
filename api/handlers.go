package api

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templates = make(map[string]*template.Template)

func init() {
	var err error
	templates["index"], err = template.ParseFiles("api/templates/index.html", "api/templates/contact-item.html")
	if err != nil {
		log.Fatal(err)
	}

	templates["404"], err = template.ParseFiles("api/templates/404.html")
	if err != nil {
		log.Fatal(err)
	}
}

func RoutesHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		GetIndexHandler(w, r)
	} else if r.URL.Path == "/ping" {
		HandlePing(w, r)
	} else {
		NotFoundHandler(w, r)
	}
}

func HandlePing(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "pong")
}

func GetIndexHandler(w http.ResponseWriter, r *http.Request) {
	err := templates["index"].ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		log.Println("Error when executing template", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	var buf bytes.Buffer
	err := templates["404"].ExecuteTemplate(&buf, "404.html", nil)
	if err == nil {
		w.WriteHeader(http.StatusNotFound)
		buf.WriteTo(w)
		return
	}

	if err != nil {
		log.Println("Error when executing template", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
