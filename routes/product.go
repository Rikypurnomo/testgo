package routes

import (
	"testgo/handlers"
	"testgo/pkg/database"
	"testgo/pkg/middleware"
	"testgo/repository"
	"testgo/services"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(g *gin.RouterGroup) {
	
	productRepo := repository.NewProductRepository(database.DB)
	productService := services.InitiateServicessProductInterface(productRepo)

	h := handlers.InitHandlers(productService)

	productGroup := g.Group("/products")
	{
		productGroup.GET("/list", h.ListProductsPaginated)
		productGroup.GET("/getbyid/:id", h.GetProduct)
		productGroup.POST("/create", middleware.Auth(middleware.UploadFile(h.CreateProduct)))
		productGroup.DELETE("/delete/:id", middleware.Auth(h.DeleteProduct))
		productGroup.PATCH("/update/:id", middleware.Auth(middleware.UploadFile(h.UpdateProduct)))
	}
}
