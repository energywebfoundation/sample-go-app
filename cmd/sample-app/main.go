package main

import (
	"io"
	"log"
	"net/http"

	"github.com/energywebfoundation/sample-go-app/internal/latestblock"
)

func main() {

	shouldReturnHealth := true

	handlers := latestblock.WrapperStruct{RpcUrl: "https://volta-rpc.energyweb.org"}
	// Hello world

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	}

	healthHandler := func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case "GET":
			if shouldReturnHealth {
				w.WriteHeader(http.StatusOK)
				io.WriteString(w, "Ok")

				return
			}

			w.WriteHeader(http.StatusInternalServerError)
			io.WriteString(w, "not ok")
		case "POST":
			shouldReturnHealth = !shouldReturnHealth

			log.Printf("current health status is %s", shouldReturnHealth)

			w.WriteHeader(http.StatusOK)
			io.WriteString(w, "ok")
		}

	}

	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/block", handlers.GetBlockHandler)
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/health", healthHandler)

	log.Println("Listing for requests at http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
