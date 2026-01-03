package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/minhduoc-tran/get-started-go-backend/internal/handlers"
)

func main() {

	r := gin.Default() // create a Gin router with default middleware


	r.GET("/hello", handlers.HelloHandler)

	log.Println("Server run http://localhost:8080")

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}

}
