package controllers

import (
	"net/http"
	"root/auth"
	controllers "root/controllers/interfaces"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	DriverService controllers.DriverService
	AuthService   controllers.AuthService
}

type TokenRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (this *AuthController) GenerateToken(context *gin.Context) {
	var request TokenRequest

	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	record, err := this.DriverService.GetDriverByEmail(request.Email)
	// check if email exists and password is correct

	credentialError := this.AuthService.CheckPassword(request.Password, record.PasswordHash)
	if credentialError != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		context.Abort()
		return
	}

	tokenString, err := auth.GenerateJWT(record.Email, record.Fullname)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	context.JSON(http.StatusOK, gin.H{"token": tokenString})
}
