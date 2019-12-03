package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/aria-music/aria-alexa-skip/models/alexa"
)

func startServer() error {
	handler := http.NewServeMux()
	server := &http.Server{
		Addr:    ":" + config.Port,
		Handler: handler,
	}

	handler.HandleFunc("/", dispatch)

	log.Println("Starting server...")
	return server.ListenAndServe()
}

func dispatch(w http.ResponseWriter, r *http.Request) {
	var root alexa.RootRequest
	err := json.NewDecoder(r.Body).Decode(&root)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("Failed to parse RootRequest: %v\n", err)
		return
	}

	if !config.ValidClientID(root.ApplicationID) {
		w.WriteHeader(http.StatusForbidden)
		log.Printf("Invalid ClientID")
		return
	}

	var request alexa.Request
	err = json.Unmarshal(root.Request, &request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("Failed to parse Request: %v\n", err)
		return
	}

	// log.Printf("Handling request type: %s\n", request.Type)
	switch request.Type {
	case "IntentRequest":
		handleIntent(w, &root)
	case "LaunchRequest":
		handleLaunch(w, &root)
	default:
		log.Printf("No handler matched for %s\n", request.Type)
	}
}
