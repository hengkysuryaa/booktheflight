package controllers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hengkysuryaa/booktheflight/backend/repository"
	"github.com/hengkysuryaa/booktheflight/backend/services"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewHandler() http.Handler {
	err := godotenv.Load()
	if err != nil {
		panic("failed to load env")
	}

	// init dependencies
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	repo := repository.NewRepository(db)
	seatSvc := services.NewSeatService(repo)
	seatHandler := NewSeat(seatSvc)

	// define routes
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		c.String(http.StatusOK, "Welcome Gin Server")
	})

	v1Api := router.Group("/v1")
	v1Api.GET("/seat", seatHandler.Get)

	return router.Handler()
}
