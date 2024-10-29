package main

import (
	"log"
	"myproject/initialize"
)

func main() {

	r := initialize.InitializeApp()

	if err := r.Run(":8082"); err != nil {
		log.Fatalf("Error starting the server: %v", err)
	}
}
