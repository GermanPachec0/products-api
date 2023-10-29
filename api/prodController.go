package api

import (
	"log"
	"net/http"

	"products-api/db" // Replace with your actual project path

	"github.com/gin-gonic/gin"
)

func GetProds(c *gin.Context) {
	prds, err := db.GetProducts()
	if err != nil {
		log.Fatal(err)
	}

	c.IndentedJSON(http.StatusOK, prds)
}
