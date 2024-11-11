package landingpage

import (
	"embed"
	"html/template"
	"log"
	"net/http"
)

//go:embed templates/*
var templates embed.FS

func Handler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving landing page")

	tmpl, err := template.ParseFS(templates, "templates/index.html")
	if err != nil {
		log.Println("Failed to parse template:", err)
		http.Error(w, "Failed to load page", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")

	if err := tmpl.Execute(w, nil); err != nil {
		log.Println("Failed to execute template:", err)
		http.Error(w, "Failed to render page", http.StatusInternalServerError)
	}
}
