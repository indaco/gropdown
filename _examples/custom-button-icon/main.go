package main

import (
	"log"
	"net/http"
	"os"

	"github.com/a-h/templ"
	"github.com/indaco/gropdown"
)

func HandleHome(w http.ResponseWriter, r *http.Request) {
	buttonIcon := `<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor">
  <path stroke-linecap="round" stroke-linejoin="round" d="M19.5 13.5 12 21m0 0-7.5-7.5M12 21V3" />
</svg>`
	button := gropdown.DropdownButton{Label: "Menu", Icon: buttonIcon}
	items := []gropdown.DropdownItem{
		{Label: "Settings", Href: "/settings"},
		{Label: "GitHub", Href: "https://github.com", External: true},
		{Divider: true},
		{Label: "Button", Attrs: templ.Attributes{"onclick": "alert('Hello gropdown');"}},
	}
	dropdown := gropdown.NewDropdownBuilder().WithButton(button).WithItems(items)

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
