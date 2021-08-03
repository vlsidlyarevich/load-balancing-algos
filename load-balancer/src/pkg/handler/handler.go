package handler

import (
	"io"
	"log"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	_, err := io.WriteString(w, "Hello, world!\n")
	if err != nil {
		log.Println("Error during handling response: ", err)
	}
}
