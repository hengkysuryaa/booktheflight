package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func NewHandler() http.Handler {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		c.String(http.StatusOK, "Welcome Gin Server")
	})

	// handlers
	seatHandler := NewSeat()

	v1Api := router.Group("/v1")
	v1Api.GET("/seat", seatHandler.Get)

	return router.Handler()
}
