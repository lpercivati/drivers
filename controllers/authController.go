package controllers

import (
	"net/http"
	"root/auth"
	"root/bodies"
	controllers "root/controllers/interfaces"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	DriverService controllers.DriverService
	AuthService   controllers.AuthService
}

func (this *AuthController) GenerateToken(context *gin.Context) {
	var request bodies.TokenRequest

	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	driver, err := this.DriverService.GetDriverByEmail(request.Email)

	credentialError := this.AuthService.CheckPassword(request.Password, driver.PasswordHash)
	if credentialError != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		context.Abort()
		return
	}

	tokenString, err := auth.GenerateJWT(driver.Email, driver.Fullname)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	context.JSON(http.StatusOK, gin.H{"token": tokenString})
}
