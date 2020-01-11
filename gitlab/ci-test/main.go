package main

import (
	"net/http"
	"os"
)

func main()  {
	http.HandleFunc("/hostname", func(w http.ResponseWriter, r *http.Request) {
		hostname, err := os.Hostname()
		if err != nil {
			w.Write([]byte("hello world"))
		}
		w.Write([]byte(hostname))
	})
	http.ListenAndServe(":8080", nil)
}
