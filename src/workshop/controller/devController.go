package controller

import (
	"github.com/gin-gonic/gin"
	"workshop/model"
	"workshop/controller/service"
	"net/http"
)

func CreateDev(c *gin.Context) {
	var dev model.Dev
	c.BindJSON(&dev)
	service.InsertDev(dev, c)
}

func UpdateDevActive(c *gin.Context)  {
	var dev model.Dev
	err := c.BindJSON(&dev)


	if err != nil {
		c.JSON(http.StatusOK,gin.H{"status":"Error","message":err.Error()})
	}else {
		result,err := service.UpdateDevActive(dev)
		if err != nil {
			c.JSON(http.StatusInternalServerError,gin.H{"status":"Error","message":err.Error()})
		}else {
			c.JSON(http.StatusOK,gin.H{"status":"Success","message":result})
		}
	}
}