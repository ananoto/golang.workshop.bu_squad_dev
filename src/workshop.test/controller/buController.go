package controller

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"workshop.test/model"
	"net/http"
	"workshop.test/service"
)

func CreateBu(c *gin.Context)  {
	var bu model.Bu
	err := c.BindJSON(&bu)

	if err != nil {
		c.JSON(http.StatusOK,gin.H{"status":"Error","message":err.Error()})
	}else {
		result,err := service.InsertBu(bu)
		if err != nil {
			c.JSON(http.StatusInternalServerError,gin.H{"status":"Error","message":err.Error()})
		}else {
			c.JSON(http.StatusOK,gin.H{"status":"Success","message":result})
		}
	}
}

func FindBu(c *gin.Context)  {
	var bu model.Bu
	err := c.BindJSON(&bu)
	if err != nil {
		c.JSON(http.StatusOK,gin.H{"status":"Error","message":err.Error()})
	}else {
		result,err := service.FindBuByName(bu.Name)
		if err != nil {
			c.JSON(http.StatusInternalServerError,gin.H{"status":"Error","message":err.Error()})
		}else {
			c.JSON(http.StatusOK,gin.H{"status":"Success","message":result})
		}
	}
}

func InsertSquadToBu(c *gin.Context)  {
	var squadBu model.SquadBu
	err := c.BindJSON(&squadBu)
	fmt.Println("squad>>>>>",squadBu.Squad_name,"\nbu>>>>>",squadBu.Bu_name)
	if err != nil {
		c.JSON(http.StatusOK,gin.H{"status":"Error","message":err.Error()})
	}else {
		result,err := service.InsertSquadToBu(squadBu.Squad_name,squadBu.Bu_name)
		if err != nil {
			c.JSON(http.StatusInternalServerError,gin.H{"status":"Error","message":err.Error()})
		}else {
			c.JSON(http.StatusOK,gin.H{"status":"Success","message":result})
		}
	}
}


func DeactiveBu(c *gin.Context)  {
	var bu model.Bu
	err := c.BindJSON(&bu)
	fmt.Println(">>>>>Bu name",bu.Name)
	if err != nil {
		c.JSON(http.StatusOK,gin.H{"status":"Error","message":err.Error()})
	}else {
		result,err := service.DeactiveBu(bu.Name)
		if err != nil {
			c.JSON(http.StatusInternalServerError,gin.H{"status":"Error","message":err.Error()})
		}else {
			c.JSON(http.StatusOK,gin.H{"status":"Success","message":result})
		}
	}
}