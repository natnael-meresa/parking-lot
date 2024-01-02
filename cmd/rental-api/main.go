package main

import (
	"log"
	"os"

	"parking-lot/internal/app"
)

func main() {
	// Create a new instance of the application.
	if err := app.Run(); err != nil {
		log.Println(err)
		// Exit with status 1 if an error occurs.	
		os.Exit(1)
	}
}
