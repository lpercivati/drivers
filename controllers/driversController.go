package controllers

import (
	"net/http"
	"root/bodies"
	"root/models"

	"github.com/gin-gonic/gin"
)

type DriverService interface {
	Create(data bodies.DriverBody) (models.Driver, error)
}

type DriverController struct {
	Service DriverService
}

func (_ *DriverController) Get(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello! " + c.Param("id"),
	})
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
