package main

import (
	"fmt"
	"test/internal/env" // Adjust the import path to your module name
)

func main() {
	// Load environment variables
	env.LoadEnv()

	// Get the environment variable
	myVariable := env.GetEnvVariable("MY_VARIABLE")

	// Print the environment variable to the console
	fmt.Println("MY_VARIABLE:", myVariable)
}
