package handlers

import (
	"context"
	"net/http"
	"strings"

	backend "github.com/billzayy/timesheet-management-be"
	"github.com/billzayy/timesheet-management-be/internal/dto"
	"github.com/billzayy/timesheet-management-be/internal/middleware"
	"github.com/billzayy/timesheet-management-be/internal/services"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service    services.AuthService
	permission services.PermissionService
}

func NewAuthHandler(s services.AuthService, p services.PermissionService) *AuthHandler {
	return &AuthHandler{s, p}
}

// Login godoc
//
//	@Summary		Login
//	@Description	Login Timesheet
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			request	body		dto.LoginDTO	true	"Send Login Information"
//	@Success		200		{object}	backend.ResponseData
//	@Router			/auth/login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	ctx := context.Background()
	c.Header("Content-Type", "application/json")

	var input dto.LoginDTO
	var errStr string

	if err := c.ShouldBindJSON(&input); err != nil {
		errStr = err.Error()

		c.JSON(http.StatusBadRequest, backend.ResponseData{
			Result:  nil,
			Success: false,
			Error:   &errStr,
		})
		return
	}

	data, err := h.service.Login(ctx, &input)

	if err != nil {
		errStr = err.Error()

		c.JSON(http.StatusInternalServerError, backend.ResponseData{
			Result:  nil,
			Success: false,
			Error:   &errStr,
		})
		return
	}

	c.JSON(http.StatusOK, backend.ResponseData{
		Result:  data,
		Success: true,
		Error:   nil,
	})
}

// Refresh godoc
//
//	@Summary		Refresh
//	@Description	Refresh Timesheet
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			request	body		dto.RefreshDTO	true	"Send Token"
//	@Success		200		{object}	backend.ResponseData
//	@Router			/auth/refresh-token [post]
func (h *AuthHandler) RefreshToken(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	var input dto.RefreshDTO
	var errStr string

	if err := c.ShouldBindJSON(&input); err != nil {
		errStr = err.Error()

		c.JSON(http.StatusBadRequest, backend.ResponseData{
			Result:  nil,
			Success: false,
			Error:   &errStr,
		})
		return
	}

	parts := strings.Split(input.RefreshToken, " ")
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
		return
	}

	userID, err := middleware.VerifyRefreshToken(parts[1])
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	accessToken, exp, err := middleware.AccessToken(userID)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	refreshToken, err := middleware.RefreshToken(userID)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, backend.ResponseData{
		Result: dto.RespLoginDTO{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
			ExpiredTime:  int(exp),
			TokenType:    "Bearer",
		},
		Success: true,
		Error:   &errStr,
	})
}

// Get User Config godoc
//
//	@Summary		Get User Config
//	@Description	Get User Configuration for App
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Success		200		{object}	backend.ResponseData
//	@Router			/user-config [get]
func (h *AuthHandler) UserConfig(c *gin.Context) {
	ctx := context.Background()
	c.Header("Content-Type", "application/json")

	var errStr string

	id := backend.GetTokenId(c)

	result, err := h.permission.GetAuthConfig(ctx, id)

	if err != nil {
		errStr = err.Error()

		c.JSON(http.StatusInternalServerError, backend.ResponseData{
			Result:  nil,
			Success: false,
			Error:   &errStr,
		})
		return
	}

	c.JSON(http.StatusOK, backend.ResponseData{
		Result:  &result,
		Success: true,
		Error:   nil,
	})
}
