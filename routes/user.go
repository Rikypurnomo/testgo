package routes

import (
	"testgo/handlers"
	"testgo/pkg/database"
	"testgo/pkg/middleware"
	"testgo/repository"
	"testgo/services"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.RouterGroup) {
	userRepo := repository.NewUsersRepository(database.DB)
	userService := services.InitiateServicessUserInterface(userRepo)

	h := handlers.InitHandlers(userService)

	r.GET("/users/list", middleware.Auth(h.ListUsers))
	r.GET("/user/:id", middleware.Auth(h.GetUser))

}
