package main

import (
	"fmt"
		"net/http"
		)

		func handler(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello, Cloud Run, hello")
			}

			func main() {
				http.HandleFunc("/", handler)
					http.ListenAndServe(":8080", nil)
					}
