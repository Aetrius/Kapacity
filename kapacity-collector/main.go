package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", WebHandler)
	http.HandleFunc("/json", JsonHandler)
	fmt.Println("Server starting on port 8080...")
	http.ListenAndServe(":8080", nil)
}

func WebHandler(w http.ResponseWriter, r *http.Request) {
	HomePage(w, r)
}

func JsonHandler(w http.ResponseWriter, r *http.Request) {
	JsonPage(w, r)
}
