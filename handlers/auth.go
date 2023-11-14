package handlers

import (
	"net/http"
	"testgo/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func (h *HandlersInit) Register(ctx *gin.Context) {

	var req models.User

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	err := h.Services.ServicesSignUp(ctx, &req)

	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	SuccesResponse(ctx, http.StatusOK, "yeay Succes", nil)
}

func (h *HandlersInit) Login(ctx *gin.Context) {
	var req models.User
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	user, err := h.Services.ServicesLogin(ctx, &req)
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	SuccesResponse(ctx, http.StatusOK, "Login Succes", user)

}

func (h *HandlersInit) CheckAuth(ctx *gin.Context) {
	userLogin := ctx.MustGet("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(float64)

	user, _ := h.Services.CheckAuth(int(userId))
	SuccesResponse(ctx, http.StatusOK, "Succes Get Profile", user)
}
