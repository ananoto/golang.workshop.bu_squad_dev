package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"workshop.test/service"
	"workshop.test/model"
)

func CreateDev(c *gin.Context)  {
	var dev model.Dev
	err := c.BindJSON(&dev)

	if err != nil {
		c.JSON(http.StatusOK,gin.H{"status":"Error","message":err.Error()})
	}else {
		result,err := service.InsertDev(dev)
		if err != nil {
			c.JSON(http.StatusInternalServerError,gin.H{"status":"Error","message":err.Error()})
		}else {
			c.JSON(http.StatusOK,gin.H{"status":"Success","message":result})
		}
	}
}

func FindDev(c *gin.Context)  {
	var dev model.Dev
	err := c.BindJSON(&dev)
	if err != nil {
		c.JSON(http.StatusOK,gin.H{"status":"Error","message":err.Error()})
	}else {
		result,err := service.FindDevByName(dev.Dev_name)
		if err != nil {
			c.JSON(http.StatusInternalServerError,gin.H{"status":"Error","message":err.Error()})
		}else {
			c.JSON(http.StatusOK,gin.H{"status":"Success","message":result})
		}
	}
}

