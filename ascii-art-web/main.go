package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var templates *template.Template

func init() {
	// Load all templates from the templates folder
	templates = template.Must(template.ParseGlob("templates/*.html"))
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		inputText := r.FormValue("inputText")
		banner := r.FormValue("banner")
		data := struct {
			InputText string
			Banner    string
		}{
			InputText: inputText,
			Banner:    banner,
		}
		// Use the output.html template to render the response
		err := templates.ExecuteTemplate(w, "output.html", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else {
		// Use the form.html template to render the form
		err := templates.ExecuteTemplate(w, "form.html", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func main() {
	http.HandleFunc("/", formHandler)
	// Serve static files from the static folder
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	fmt.Println("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server failed to start:", err)
	}
}
