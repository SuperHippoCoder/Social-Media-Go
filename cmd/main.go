package main

import (
	"os"
	"social-media-go/internal/models"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	dbModel, err := models.InitDB("./database/users.db")
	if err != nil {
		os.Exit(1)
	}

	h := NewHandler(dbModel)

	loadTemplates(router)

	setupRoutes(router, h)

	router.Run(":8080")
}
