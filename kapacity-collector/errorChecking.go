package main

import "fmt"

func errorCheckingK8s(errorIn error) {
	if errorIn != nil {
		fmt.Println("Error: ", errorIn)
		return
	}
}
