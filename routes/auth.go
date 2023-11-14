package routes

import (
	"testgo/handlers"
	"testgo/pkg/middleware"
	"testgo/pkg/database"
	"testgo/repository"
	"testgo/services"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.RouterGroup) {
	authRepo := repository.NewAuthRepository(database.DB)
	userService := services.InitiateServicessAuthInterface(authRepo)

	h := handlers.InitHandlers(userService)

	r.POST("/register", h.Register)
	r.POST("/login", h.Login)
	r.GET("/profile", middleware.Auth(h.CheckAuth))
}
