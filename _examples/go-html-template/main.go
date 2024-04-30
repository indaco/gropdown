package main

import (
	"context"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/indaco/gropdown"
)

// PageData represents the data to be rendered in the HTML template
type PageData struct {
	Dropdown DropdownComponent
}

type DropdownComponent struct {
	CSS  template.HTML
	HTML template.HTML
	JS   template.HTML
}

// HandleHome is the handler function for the home page "/"
func HandleHome(w http.ResponseWriter, r *http.Request) {
	config := gropdown.NewConfigBuilder().Build()
	configMap := gropdown.NewConfigMap()
	configMap.Add("demo", config)
	ctx := context.WithValue(r.Context(), gropdown.ConfigContextKey, configMap)

	dropdownHTMLGenerator := gropdown.NewHTMLGenerator()

	// Render the needed CSS for dropdown component as template.HTML
	dropdownCSS, err := dropdownHTMLGenerator.GropdownCSSToGoHTML()
	if err != nil {
		log.Printf("failed to render template: %v", err)
	}

	// Render the dropdown component into a template.HTML
	dropdownHtml, err := dropdownHTMLGenerator.Render(ctx, GropdownDemo())
	if err != nil {
		log.Printf("failed to render template: %v", err)
	}

	// Render the needed JS for dropdown component as template.HTML
	dropdownJS, err := dropdownHTMLGenerator.GropdownJSToGoHTML(configMap)
	if err != nil {
		log.Printf("failed to render template: %v", err)
	}

	data := PageData{
		Dropdown: DropdownComponent{
			CSS:  dropdownCSS,
			HTML: dropdownHtml,
			JS:   dropdownJS,
		},
	}

	// Parse the HTML template
	tmpl := template.Must(template.New("index").Parse(`
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>gropdown - template/html</title>
			<!-- inject dropdown styles -->
			{{ .Dropdown.CSS }}
			<!-- styles -->
			<style type="text/css">
				.parent-container {
					display: flex;
					justify-content: center;
					align-items: center;
					height: 100vh;
				}
			</style>
		</head>
		<body class="parent-container">
			<div class="centered">
				<!-- display the dropdown -->
				{{ .Dropdown.HTML }}
			</div>
			<!-- inject dropdown javascript -->
			{{ .Dropdown.JS }}
		</body>
		</html>
	`))

	// Execute the template with the provided data and write the output to the response writer
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "failed to render template", http.StatusInternalServerError)
		log.Printf("failed to render template: %v", err)
		return
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", HandleHome)

	port := ":3300"
	log.Printf("Listening on %s", port)
	if err := http.ListenAndServe(port, mux); err != nil {
		log.Printf("failed to start server: %v", err)
		os.Exit(1)
	}
}
