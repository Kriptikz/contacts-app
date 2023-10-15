package api

import (
	"bytes"
	"contacts-app/data/database"
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
		HandleGetIndex(w, r)
	} else if r.URL.Path == "/ping" {
		HandlePing(w, r)
	} else if r.URL.Path == "/contacts" {
		HandleGetContacts(w, r)
	} else {
		NotFoundHandler(w, r)
	}
}

func HandlePing(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "pong")
}

func HandleGetContacts(w http.ResponseWriter, r *http.Request) {
	data, err := database.DB.GetAllContacts()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error: Failed to get all contacts")
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "pong")
	for _, contact := range data {
		fmt.Fprintf(w, "ID: %d, Name: %s, Email: %s, Phone: %s\n", contact.ID, contact.Name, contact.Email, contact.Phone)
	}
}

func HandleGetIndex(w http.ResponseWriter, r *http.Request) {
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
