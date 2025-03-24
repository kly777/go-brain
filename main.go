// main.go
package main

import (
	"fmt"
	"go-brain/initializer"
	"log"
	"time"
)

func main() {
	router, cleanup, err := initializer.Init()
	if err != nil {
		log.Fatalf("Failed to initialize app: %v", err)
	}
	defer cleanup()

	fmt.Println(time.Now().Format("2006/01/02 15:04:05"), "Server is running on port 8080")
	fmt.Println("\nGREAT!!!!!")

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
