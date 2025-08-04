package main

import (
	"fmt"
	"log"
	"net/http"
	"slices"
	"strings"

	"github.com/heptalium/httputil"
)

func main() {
	log.Print("Starting Signal Gateway...")
	readConfig()
	log.Printf("Using signal-cli JSON-RPC API at %s", config.SignalCliEndpoint)
	initRpcClient()
	log.Printf("Using Signal account %s", config.Account)

	if len(config.AllowedRecipients) > 0 {
		log.Printf("Allowed recipients: %s", strings.Join(config.AllowedRecipients, ", "))
	} else {
		log.Print("All recipients are allowed")
	}

	log.Printf("API endpoint: %s", config.Endpoint)

	http.HandleFunc(config.Endpoint, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			log.Printf("Method %s is not allowed!", r.Method)
			w.Header().Set("Allow", http.MethodPost)
			httputil.WriteHttpStatus(w, http.StatusMethodNotAllowed)
			return
		}

		var data struct {
			Recipient string
			Message   string
		}

		if err := httputil.ParseRequest(w, r, &data); err != nil {
			log.Printf("Could not parse request: %v", err)
			return
		}

		if len(config.AllowedRecipients) > 0 {
			if !slices.Contains(config.AllowedRecipients, data.Recipient) {
				log.Printf("Recipient %s is not allowed!", data.Recipient)
				httputil.WriteHttpStatus(w, http.StatusForbidden)
				return
			}
		}

		log.Printf("Sending message to %s: %s", data.Recipient, data.Message)

		if err := sendMessage(data.Recipient, data.Message); err != nil {
			log.Printf("Could not send message: %v", err)
			httputil.WriteHttpStatus(w, http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	})

	if config.FormEndpoint != "" {
		log.Printf("Test form endpoint: %s", config.FormEndpoint)
		http.HandleFunc(config.FormEndpoint, func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(fmt.Sprintf(form, config.Endpoint)))
		})
	}

	log.Printf("Starting service on :%d", config.Port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", config.Port), nil)
	log.Fatal(err)
}
