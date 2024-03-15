package main

import (
	"log"
	"net/http"
	"os"

	"github.com/a-h/templ"
	"github.com/indaco/gropdown"
)

func HandleHome(w http.ResponseWriter, r *http.Request) {
	button := gropdown.DropdownButton{Label: "Menu"}
	items := []gropdown.DropdownItem{
		{Label: "Settings", Href: "/settings"},
		{Label: "GitHub", Href: "https://github.com", External: true},
		{Divider: true},
		{Label: "Button", Attrs: templ.Attributes{"onclick": "alert('Hello gropdown');"}},
	}
	dropdown := gropdown.NewDropdownBuilder().
		SetPosition(gropdown.Right).
		WithButton(button).
		WithItems(items)

	templ.Handler(HomePage(dropdown)).ServeHTTP(w, r)
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
