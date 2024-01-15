package main

import (
	"errors"
	"io"
	"log"
	"net/http"
	"os"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s /\n", r.Method)
	io.WriteString(w, "ok")
}

func main() {
	http.HandleFunc("/", RootHandler)

	err := http.ListenAndServe(":8090", nil)

	if errors.Is(err, http.ErrServerClosed) {
		log.Println("server closed")
	} else if err != nil {
		log.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
