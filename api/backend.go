package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func handleCreate(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /create request\n")
	io.WriteString(w, "Resource created!\n")
}

func handleUpdate(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /update request\n")
	io.WriteString(w, "Resource updated!\n")
}

func handleDelete(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /delete request\n")
	io.WriteString(w, "Resource deleted\n")
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/create", handleCreate)
	mux.HandleFunc("/update", handleUpdate)
	mux.HandleFunc("/delete", handleDelete)

	err := http.ListenAndServe(":9999", mux)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
