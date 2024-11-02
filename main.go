package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("impression counter system started.")

	config, err := loadConfig()
	if err != nil {
		fmt.Printf("config load failed. %e\n", err)
		os.Exit(1)
	}
	
	fmt.Println("impression counter system ended.")
}
