package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	log.Println("Hello world!")

	var handler = func(w http.ResponseWriter, r *http.Request) {
		_, err := io.WriteString(w, "Hello, world!\n")
		if err != nil {
			log.Println("Error during handling response: ", err)
		}
	}

	http.HandleFunc("/hello", handler)

	log.Fatal(http.ListenAndServe(":8081", nil))
}
