package main

import (
	"io"
	"log"
	"net/http"

	"github.com/energywebfoundation/sample-go-app/internal/latestblock"
)


func main() {

	handlers := latestblock.WrapperStruct{RpcUrl: "https://volta-rpc.energyweb.org"}
	// Hello world, the web server

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	}

	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/block", handlers.GetBlockHandler)
	http.HandleFunc("/", helloHandler)
	log.Println("Listing for requests at http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
