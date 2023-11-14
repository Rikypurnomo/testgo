# Api-go
### user hanya dapat mengakses 
  ```bash
  /api/register
/api/login
/api/profile
/api/products/list
/api/products/getbyid
  ```

### untuk login sebagain admin register is_admin:true
##akses admin
 ```bash
/api/register
/api/login
/api/profile
/api/products/list
/api/products/getbyid
 /api/products/create
/api/products/delete
/api/products/update
/api/users/list
/api/user
  ```

# Prepare

### 1. Download

- Golang: [Click Here](https://go.dev/dl/)
- XAMPP: [Click Here](https://www.apachefriends.org/download.html)
- Postman: [Click Here](https://www.postman.com/downloads/?utm_source=postman-home)

### 2. Install

- Follow the instruction from [Golang Official Website](https://go.dev/doc/install)

- How to Check on Terminal or Command Propt:

  ```bash
  go version
  ```

### 3. Initializing project

```bash
go mod init _project_name_
```

### 4. Install Gin

```bash
go get -u github.com/gin-gonic/gin
```

# cara pakai api

- pergi ke postman lalu login dengan route yang sesuai

`http://localhost:5050/api/v1/login`

- Kemudian Create Product terlebih dahulu

- pergi ke folder routes 

contoh

`http://localhost:5050/api/v1/product`


> File : `routes/product.go`

```go
package routes

import (
	"server/handlers"
	"server/pkg/middleware"
	"server/pkg/mysql"
	"server/repositories"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(g *gin.RouterGroup) {
	productRepository := repositories.RepositoryProduct(mysql.DB)
	h := handlers.HandlerProduct(productRepository)

	g.GET("/products", middleware.Auth(h.FindProductsPaginated))
	g.GET("/product/:id", middleware.Auth(h.GetProduct))
	g.POST("/product", middleware.Auth(middleware.UploadFile(h.CreateProduct)))
	g.DELETE("/product/:id", middleware.Auth(h.DeleteProduct))
	g.PATCH("/product/:id", middleware.Auth(middleware.UploadFile(h.UpdateProduct)))
}
```

# find product with fitur paginate 

### contoh

  ```bash
  http://localhost:5050/api/products/list?page=1&limit=5
  ```
