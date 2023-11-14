package handlers

import (
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func (h *HandlersInit) ListUsers(ctx *gin.Context) {
	UserLogin := ctx.MustGet("userLogin")
	userAdmin := UserLogin.(jwt.MapClaims)["is_admin"].(bool)

	if userAdmin {
		users, err := h.Services.ListUsers()
		if err != nil {
			ErrorResponse(ctx, http.StatusBadRequest, err.Error())
			return
		}

		SuccesResponse(ctx, http.StatusOK, "Success List User", users)
	} else {
		ErrorResponse(ctx, http.StatusUnauthorized, "You're not Admin. Admin only.")
	}
}

func (h *HandlersInit) GetUser(ctx *gin.Context) {
	UserLogin := ctx.MustGet("userLogin")
	userAdmin := UserLogin.(jwt.MapClaims)["is_admin"].(bool)

	if userAdmin {
		id, _ := strconv.Atoi(ctx.Param("id"))
		
		user, err := h.Services.GetUser(id)
		if err != nil {
			ErrorResponse(ctx, http.StatusBadRequest, err.Error())
			return
		}

		SuccesResponse(ctx, http.StatusOK, "Success Get User", user)
	} else {
		ErrorResponse(ctx, http.StatusUnauthorized, "You're not Admin. Admin only.")
	}
}
