package router

import (
	"github.com/aqilknz/koda-b7-backend/internal/controller"
	"github.com/gin-gonic/gin"
)

func SetupAuthRoutes(r *gin.Engine, ac *controller.AuthController) {
	r.POST("/register", ac.Register)
	r.POST("/login", ac.Login)
}
