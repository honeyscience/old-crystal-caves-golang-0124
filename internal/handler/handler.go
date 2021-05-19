package handler

import (
	"fmt"
	"log"
	"net/http"
)

func socketHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("Handled!")
	fmt.Fprintf(w, "Welcome to golang")
}

func Initialize() {
	log.Println("Starting a websocket server to handle connections. The socket is at http://localhost:8080")
	http.HandleFunc("/", socketHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
