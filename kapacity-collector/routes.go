package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	htmlContent := GetHomePage()

	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(htmlContent))
}

func JsonPage(w http.ResponseWriter, r *http.Request) {
	var allContainerInfo []TypeInfo
	var err error

	allContainerInfo, err = gatherKubernetesInfo()

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "    ") // Set indentation: here using four spaces

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("Error: " + err.Error())
		return
	}

	err = encoder.Encode(allContainerInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("Error: " + err.Error())
		return
	}
}

func PodsJsonPage(w http.ResponseWriter, r *http.Request) {
	var allContainerInfo []Pod
	var err error

	allContainerInfo, err = gatherKubernetesPods()

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "    ") // Set indentation: here using four spaces

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("Error: " + err.Error())
		return
	}

	err = encoder.Encode(allContainerInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("Error: " + err.Error())
		return
	}
}
