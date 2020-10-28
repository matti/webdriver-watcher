package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/matti/webdriver-watcher/internal/checker"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ok, maybe, status := checker.Check("http://localhost:9515")
		fmt.Fprintf(w, status)

		if !maybe {
			fmt.Println(time.Now().String(), status)
		}
		if !ok {
			w.WriteHeader(http.StatusGatewayTimeout)
		}
	})

	fmt.Println("listen 0.0.0.0:9516")
	log.Fatal(http.ListenAndServe("0.0.0.0:9516", nil))
}
