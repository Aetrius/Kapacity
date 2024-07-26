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
	var allContainerInfo []PodInfo
	var err error

	allContainerInfo, err = gatherKubernetesInfo()

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "    ") // Set indentation: here using four spaces

	err = encoder.Encode(err)
	errorCheck(err, w) // write error from gather check

	err = encoder.Encode(allContainerInfo)
	errorCheck(err, w)
}

func errorCheck(errorIn error, w http.ResponseWriter) {
	if errorIn != nil {
		http.Error(w, errorIn.Error(), http.StatusInternalServerError)
		fmt.Println(errorIn.Error())

		return
	}
}
