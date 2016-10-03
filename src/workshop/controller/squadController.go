package controller

import (
	"github.com/gin-gonic/gin"
	"workshop/controller/service"
	"workshop/model"
	"net/http"
	"fmt"
)

func GetSquadByName(c *gin.Context)  {
	name := c.Param("name")
	result := service.FindSquadByName(name)
	fmt.Printf(">>>",result)
	c.JSON(200, gin.H{
		"bu": result ,
	})

}

func CreateSquad(c *gin.Context) {
	var squad model.Squad
	c.BindJSON(&squad)
	service.InsertSquad(squad, c)
}

func DevToSquad(c *gin.Context)  {
	var devSquad model.DevSquad
	err := c.BindJSON(&devSquad)
	if err != nil {
		c.JSON(http.StatusOK,gin.H{"status":"Error","message":err.Error()})
	}else {
		result,err:=service.InsertDevToSquad(devSquad)

		if err != nil {
			c.JSON(http.StatusInternalServerError,gin.H{"status":"Error","message":err.Error()})
		}else {
			if result.Devs != nil {
				c.JSON(http.StatusOK, gin.H{"status":"Success", "message":result})
			}else{
				c.JSON(503, gin.H{"status":"Error", "message":"can't be dublicated"})
			}
		}
	}

}

func SquadDeactivate(c *gin.Context)  {
	var squadDeactive model.SquadDeactive
	err := c.BindJSON(&squadDeactive)
	if err != nil{
		c.JSON(http.StatusOK,gin.H{"status":"Error","message":err.Error()})
	}else{
		result,err:=service.SquadDeacivate(squadDeactive)
		fmt.Println(">>>>>>",result)
		fmt.Println("Devs>>>>>>",result.Devs)
		if err != nil {
			c.JSON(http.StatusInternalServerError,gin.H{"status":"Error","message":err.Error()})
		}else {
			if result.Devs != nil {
				c.JSON(http.StatusOK, gin.H{"status":"Success", "message":result})
			}else{
				c.JSON(503, gin.H{"status":"Error", "message":"can't be dublicated"})
			}
		}
	}
}

