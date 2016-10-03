package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"workshop.test/service"
	"workshop.test/model"
	"fmt"
)

func CreateSquad(c *gin.Context)  {
	var squad model.Squad
	err := c.BindJSON(&squad)

	if err != nil {
		c.JSON(http.StatusOK,gin.H{"status":"Error","message":err.Error()})
	}else {
		result,err := service.InsertSquad(squad)
		if err != nil {
			c.JSON(http.StatusInternalServerError,gin.H{"status":"Error","message":err.Error()})
		}else {
			c.JSON(http.StatusOK,gin.H{"status":"Success","message":result})
		}
	}
}

func FindSquad(c *gin.Context)  {
	var squad model.Squad
	err := c.BindJSON(&squad)
	if err != nil {
		c.JSON(http.StatusOK,gin.H{"status":"Error","message":err.Error()})
	}else {
		result,err := service.FindSquadByName(squad.Name)
		if err != nil {
			c.JSON(http.StatusInternalServerError,gin.H{"status":"Error","message":err.Error()})
		}else {
			c.JSON(http.StatusOK,gin.H{"status":"Success","message":result})
		}
	}
}

func InsertDevToSquad(c *gin.Context)  {
	var devSquad model.DevSquad
	err := c.BindJSON(&devSquad)
	fmt.Println("dev>>>>>",devSquad.Dev_name,"\nsquad>>>>>",devSquad.Squad_name)
	if err != nil {
		c.JSON(http.StatusOK,gin.H{"status":"Error","message":err.Error()})
	}else {
		result,err := service.InsertDevToSquad(devSquad.Dev_name,devSquad.Squad_name)
		if err != nil {
			c.JSON(http.StatusInternalServerError,gin.H{"status":"Error","message":err.Error()})
		}else {
			c.JSON(http.StatusOK,gin.H{"status":"Success","message":result})
		}
	}
}


func DeactiveSquad(c *gin.Context)  {
	var squad model.Squad
	err := c.BindJSON(&squad)
	fmt.Println(">>>>>Squad name",squad.Name)
	if err != nil {
		c.JSON(http.StatusOK,gin.H{"status":"Error","message":err.Error()})
	}else {
		result,err := service.DeactiveSquad(squad.Name)
		if err != nil {
			c.JSON(http.StatusInternalServerError,gin.H{"status":"Error","message":err.Error()})
		}else {
			c.JSON(http.StatusOK,gin.H{"status":"Success","message":result})
		}
	}
}
