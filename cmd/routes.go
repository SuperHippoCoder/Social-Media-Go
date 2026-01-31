package main

import (
	"github.com/gin-gonic/gin"
)

func setupRoutes(router *gin.Engine, h *Handler /*store sessions.Store*/) {
	//router.Use(sessions.Sessions("pizza-tracker", store))

	router.GET("/login", h.HandleLoginGet)
	router.GET("/register", h.HandleRegistrationGet)
	router.POST("/login", h.HandleLoginPost)
	router.POST("/register", h.HandleRegistrationPost)

	router.GET("/lol", h.HandleLolGet)

	router.Static("/static", "./templates/static")
}
