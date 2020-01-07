package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/hostname", func(w http.ResponseWriter, r *http.Request) {
		if hostname, err := os.Hostname(); err != nil {
			_, _ = w.Write([]byte("无法获取hostname"))
		} else {
			_, _ = w.Write([]byte(hostname))
		}
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
