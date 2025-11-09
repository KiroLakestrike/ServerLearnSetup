package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func RenderTemplateTest(w http.ResponseWriter, tmpl string) {
	parsedTemplate, err := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
	if err != nil {
		log.Printf("Error parsing template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = parsedTemplate.Execute(w, nil)
	if err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// Render with Cache
// tc takes the key string with the value of template.Template
var tc = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	// Check to see if this template is already in the cache
	_, inMap := tc[t]
	if !inMap {
		// need to create the template
		err = createTemplateCache(t)
		if err != nil {
			log.Println(err)
		}
		log.Println("Creating new template")
	} else {
		// we have the template
		log.Println("Serving cached template")
	}

	tmpl = tc[t]
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Printf("Error executing template: %v", err)
	}
}

func createTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("templates/%s", t),
		"./templates/base.layout.tmpl",
		//any other templates we need
	}

	//parse the template
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}

	// add template to cache
	tc[t] = tmpl
	return nil
}
