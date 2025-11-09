package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	// Create a template cache (load and parse all templates)
	tc, err := createTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	// Retrieve the requested template from the cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal(err) // Stop program if template not found
	}

	// Create a buffer to temporarily hold the rendered output
	buf := new(bytes.Buffer)

	// Execute the template, passing nil as data context
	err = t.Execute(buf, nil)
	if err != nil {
		log.Println(err)
	}

	// Write the rendered template content to the HTTP response
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func createTemplateCache() (map[string]*template.Template, error) {
	// Create a map to store compiled templates by name
	myCache := map[string]*template.Template{}

	// Find all files with the pattern "*.page.tmpl" inside the ./templates directory
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	// Loop through every file found
	for _, page := range pages {
		// Extract only the file name (without path)
		name := filepath.Base(page)

		// Create a new template with that name and parse the main page template
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		// Find all layout templates, typically used for wrapping pages
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		// If layout templates exist, parse and associate them with the page template
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}

		// Save the compiled template set to the cache map
		myCache[name] = ts
	}

	// Return the cache map and nil error
	return myCache, nil
}
