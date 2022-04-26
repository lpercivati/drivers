package controllers

import (
	"net/http"
	"root/bodies"
	controllers "root/controllers/interfaces"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DriverController struct {
	Service     controllers.DriverService
	AuthService controllers.AuthService
}

func (this *DriverController) Get(context *gin.Context) {
	page, err := strconv.Atoi(context.Query("page"))

	if err != nil {
		context.JSON(http.StatusBadRequest, nil)
		return
	}

	count, err := strconv.Atoi(context.Query("count"))

	if err != nil {
		context.JSON(http.StatusBadRequest, nil)
		return
	}

	drivers, err := this.Service.GetDrivers(page, count)

	if err == nil {
		context.JSON(http.StatusOK, drivers)
	} else {
		context.JSON(http.StatusBadRequest, nil)
	}
}

func (this *DriverController) Add(c *gin.Context) {
	var body bodies.DriverBody
	c.Bind(&body)

	passwordHash, _ := this.AuthService.HashPassword(body.Password)
	body.Password = passwordHash

	driver, err := this.Service.Create(body)

	if err == nil {
		c.JSON(http.StatusOK, driver)
	} else {
		c.JSON(http.StatusBadRequest, nil)
	}
}

func (this *DriverController) GetAvailableDrivers(context *gin.Context) {
	drivers, err := this.Service.GetAvailableDrivers()

	if err == nil {
		context.JSON(http.StatusOK, drivers)
	} else {
		context.JSON(http.StatusBadRequest, nil)
	}
}
