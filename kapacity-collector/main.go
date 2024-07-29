package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", WebHandler)
	http.HandleFunc("/json", JsonHandler)
	http.HandleFunc("/pods", PodsHandler)
	fmt.Println("Server starting on port 8080...")
	http.ListenAndServe(":8080", nil)
}

func WebHandler(w http.ResponseWriter, r *http.Request) {
	HomePage(w, r)
}

func JsonHandler(w http.ResponseWriter, r *http.Request) {
	JsonPage(w, r)
}

func PodsHandler(w http.ResponseWriter, r *http.Request) {
	PodsJsonPage(w, r)
}
