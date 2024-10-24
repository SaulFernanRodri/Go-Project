package main

import (
	"log"
	"myproject/app"
)

func main() {

	r := app.InitializeApp()

	if err := r.Run(":8082"); err != nil {
		log.Fatalf("Error starting the server: %v", err)
	}
}
