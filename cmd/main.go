package main

import (
	"github.com/aqilknz/koda-b7-backend/internal/controller"
	"github.com/aqilknz/koda-b7-backend/internal/router"
	"github.com/aqilknz/koda-b7-backend/internal/service"
	"github.com/gin-gonic/gin"
)

func main() {
	// Inisialisasi Service & Controller
	userService := service.NewUserService()
	authController := controller.NewAuthController(userService)

	r := gin.Default()

	// Setup Routes
	router.SetupAuthRoutes(r, authController)

	r.Run("localhost:9000")
}
