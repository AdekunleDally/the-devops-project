package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "Hello World")
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Not a post request")
	}
}
func main() {
	http.HandleFunc("/api/v1/health", handler)
	fmt.Println("Server is listening on port 8082...")
	err := http.ListenAndServe(":8082", nil)
	if err != nil {
		log.Fatalf("Server Error: %v", err)
	}
}
