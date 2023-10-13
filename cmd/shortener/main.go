package main

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

var shortsURLs = make(map[string]string)

// webhookPost HTTP request handler
func webhookPost(w http.ResponseWriter, r *http.Request) {

	// Receive the request body
	bodyReq, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	if string(bodyReq) != "" {

		// Checks the presence of the received data in the map
		_, ok := shortsURLs[string(bodyReq)]
		if !ok {
			shortsURLs[string(bodyReq)] = "/" + generatorRandomShortString(8)
		}

		valueShortURL := shortsURLs[string(bodyReq)]

		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusCreated)

		if _, err := w.Write([]byte("http://" + r.Host +
			valueShortURL)); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusBadRequest)
}

func webhookGet(w http.ResponseWriter, r *http.Request) {

	var keyTrue string

	if r.RequestURI != "/" {

		for key, value := range shortsURLs {
			if value == r.RequestURI {
				keyTrue = key
			}
		}

		if keyTrue != "" {
			w.Header().Set("Location", keyTrue)
			w.WriteHeader(http.StatusTemporaryRedirect)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusBadRequest)
}

// generatorRandomShortString
func generatorRandomShortString(length int) string {
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz" +
		"0123456789")
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	return b.String()
}

// run initializing server dependencies before startup
func run(r *chi.Mux) error {
	return http.ListenAndServe(":8080", r)
}

func main() {
	r := chi.NewRouter()

	r.Route("/", func(r chi.Router) {
		r.Post("/", webhookPost)
		r.Get("/{short_url}", webhookGet)
	})

	if err := run(r); err != nil {
		fmt.Printf("The server did not start because %v\n", err)
	}
}
