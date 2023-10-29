package main

import (
	"fmt"
	"products-api/api"
	"products-api/db"

	"github.com/gin-gonic/gin"
)

func main() {
	err := db.ConnectDB()
	if err != nil {
		fmt.Printf("Error Connecting: %v", err)
		return
	}

	router := gin.Default()
	router.GET("/products", api.GetProds)
	router.Run("localhost:8080")
}
