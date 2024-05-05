package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/indaco/gropdown"
)

func HandleHome(w http.ResponseWriter, r *http.Request) {
	configOne := gropdown.NewConfigBuilder().WithPlacement(gropdown.Top).Build()
	configTwo := gropdown.NewConfigBuilder().WithCloseOnOutsideClick(false).Build()
	configMap := gropdown.NewConfigMap()
	configMap.Add("menu-1", configOne)
	configMap.Add("menu-2", configTwo)

	ctx := context.WithValue(r.Context(), gropdown.ConfigContextKey, configMap)
	err := HomePage().Render(ctx, w)
	if err != nil {
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
