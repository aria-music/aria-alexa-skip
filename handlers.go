package main

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/sarisia/aria-alexa-skip/models/alexa"
	"github.com/sarisia/aria-alexa-skip/models/aria"
)

func handleIntent(w http.ResponseWriter, r *alexa.RootRequest) {
	var intent alexa.IntentRequest
	err := json.Unmarshal(r.Request, &intent)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("Failed to parse IntentRequest: %v\n", err)
		return
	}

	log.Printf("Handling IntentRequest type: %s\n", intent.Intent.Name)

	switch intent.Intent.Name {
	case "Skip":
		performSkip(w)
	default:
		w.WriteHeader(http.StatusBadRequest)
		log.Println("No intent matches.")
	}
}

func handleLaunch(w http.ResponseWriter, r *alexa.RootRequest) {
	log.Printf("Handling LaunchRequest")
	performSkip(w)
}

func performSkip(w http.ResponseWriter) {
	go requestAria(&aria.Request{
		Token: config.AriaToken,
		OP:    "skip",
	})

	packet := alexa.NewSpeechResponse("スキップしました")
	err := json.NewEncoder(w).Encode(&packet)
	if err != nil {
		log.Printf("Failed to write response: %v\n", err)
	}
}

func requestAria(r *aria.Request) {
	// TODO: io.Pipe
	packet, err := json.Marshal(r)
	if err != nil {
		log.Printf("Failed to marshal aria.Request: %v\n", err)
		return
	}

	resp, err := http.Post(
		config.AriaEndpoint,
		"application/json",
		bytes.NewReader(packet),
	)
	if err != nil {
		log.Printf("Failed to request to aria server: %v\n", err)
		return
	}

	// CLOSE!
	io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
}
