package main

import (
	"io"
	"math/rand"
	"net/http"
	"strings"
)

var shortsURLs = make(map[string]string)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", webhook)

	if err := run(mux); err != nil {
		panic(err)
	}

}

// run initializing server dependencies before startup
func run(mux *http.ServeMux) error {
	return http.ListenAndServe(":8080", mux)
}

// webhookPost HTTP request handler
func webhook(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {

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
		return
	}

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
