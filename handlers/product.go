package handlers

import (
	"net/http"
	"strconv"
	productdto "testgo/dto/product"
	resultdto "testgo/dto/result"
	"testgo/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var path_file = "http://localhost:5000/uploads/"

func (h *HandlersInit) ListProductsPaginated(ctx *gin.Context) {

	pageStr := ctx.DefaultQuery("page", "1")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 1
	}

	limitStr := ctx.DefaultQuery("limit", "5")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 5
	}

	search := ctx.Query("search")

	products, totalCount, err := h.Services.ListProductsPaginated(page, limit, search)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
		return
	}

	response := gin.H{
		"page":       page,
		"limit":      limit,
		"search":     search,
		"totalItems": totalCount,
		"products":   products,
	}

	SuccesResponse(ctx, http.StatusOK, "Succes List Product", response)
}

func (h *HandlersInit) GetProduct(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		id = 0
	}

	product, err := h.Services.GetProductByID(id)
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	product.Image = path_file + product.Image

	SuccesResponse(ctx, http.StatusOK, "Succes Get Product By ID", product)
}

func (h *HandlersInit) CreateProduct(ctx *gin.Context) {
	ctx.Header("Content-Type", "multipart/form-data")

	userLogin := ctx.MustGet("userLogin")
	userAdmin := userLogin.(jwt.MapClaims)["is_admin"].(bool)

	if userAdmin {
		dataFile := ctx.MustGet("datafile").(string)
		Price, _ := strconv.Atoi(ctx.Request.FormValue("price"))

		request := productdto.CreateProductRequest{
			Name:        ctx.Request.FormValue("name"),
			Price:       Price,
			Description: ctx.Request.FormValue("description"),
			Image:       dataFile,
		}

		data, err := h.Services.CreateProduct(request)
		if err != nil {
			ErrorResponse(ctx, http.StatusBadRequest, err.Error())
			return
		}

		ctx.JSON(http.StatusOK, resultdto.SuccessResult{Status: http.StatusOK, Message: "Product berhasil ditambah ", Data: convertResponseProduct(data)})
	} else {
		ErrorResponse(ctx, http.StatusUnauthorized, "Sorry, you're not Admin")
		return
	}
}

func (h *HandlersInit) UpdateProduct(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		id = 0
	}

	request := productdto.Update1ProductRequest{
		Name:        ctx.PostForm("name"),
		Description: ctx.PostForm("description"),
		Image:       ctx.PostForm("image"),
	}

	userLogin, _ := ctx.Get("userLogin")
	isAdmin := userLogin.(jwt.MapClaims)["is_admin"].(bool)

	data, err := h.Services.UpdateProduct(id, isAdmin, request)
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, resultdto.SuccessResult{Status: http.StatusOK, Message: "Product data updated successfully", Data: convertResponseProduct(data)})
}

func (h *HandlersInit) DeleteProduct(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	userLogin := ctx.MustGet("userLogin")
	userAdmin := userLogin.(jwt.MapClaims)["is_admin"].(bool)

	deletedProduct, err := h.Services.DeleteProduct(id, userAdmin)
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, resultdto.SuccessResult{Status: http.StatusOK, Message: "Product berhasil dihapus", Data: convertResponseProduct(deletedProduct)})
}

func convertResponseProduct(product models.Product) productdto.ProductResponse {
	return productdto.ProductResponse{
		ID:          product.ID,
		Name:        product.Name,
		Price:       product.Price,
		Description: product.Description,
		Image:       product.Image,
	}
}

func convertResponseProducts(products []models.Product) []productdto.ProductResponse {
	var responseProducts []productdto.ProductResponse

	for _, product := range products {
		responseProducts = append(responseProducts, convertResponseProduct(product))
	}
	return responseProducts
}
