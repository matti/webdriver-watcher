package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/matti/webdriver-watcher/internal/checker"
)

func main() {
	listen := "0.0.0.0:9516"
	port := os.Getenv("PORT")
	if port != "" {
		listen = "0.0.0.0:" + port
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ok, status := checker.Check("http://localhost:9515")
		fmt.Fprintf(w, status)
		if !ok {
			w.WriteHeader(http.StatusGatewayTimeout)
		}
	})

	fmt.Println("listening", listen)
	log.Fatal(http.ListenAndServe(listen, nil))
}
