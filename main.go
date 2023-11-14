package main

import (
	"fmt"
	"testgo/pkg/database"
	"testgo/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	database.DatabaseInit()
	database.RunMigration()

	api := r.Group("/api")
	{
		routes.RouteInit(api)
	}

	fmt.Println("run localhost:8080")
	r.Run("localhost:8080")
}
