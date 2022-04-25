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

func (this *DriverController) Get(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))

	drivers, err := this.Service.GetDrivers(page)

	if err == nil {
		c.JSON(http.StatusOK, drivers)
	} else {
		c.JSON(http.StatusBadRequest, nil)
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
