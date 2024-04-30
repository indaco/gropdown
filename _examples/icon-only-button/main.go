package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/indaco/gropdown"
)

func HandleHome(w http.ResponseWriter, r *http.Request) {
	config := gropdown.NewConfigBuilder().WithPosition(gropdown.Right).Build()
	configMap := gropdown.NewConfigMap()
	configMap.Add("demo", config)

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
