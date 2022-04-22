package controllers

import (
	"net/http"
	"root/bodies"
	"root/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DriverService interface {
	Create(data bodies.DriverBody) (models.Driver, error)
	GetDrivers(page int) ([]models.Driver, error)
}

type DriverController struct {
	Service DriverService
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

	driver, err := this.Service.Create(body)

	if err == nil {
		c.JSON(http.StatusOK, driver)
	} else {
		c.JSON(http.StatusBadRequest, nil)
	}

}
