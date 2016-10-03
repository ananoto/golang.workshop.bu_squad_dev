package controller

import (
	"github.com/gin-gonic/gin"
	"workshop/controller/service"
	"workshop/model"
	"net/http"
	"fmt"
)

func CreateBu(c *gin.Context) {
	var bu model.Bu
	c.BindJSON(&bu)
	service.SaveBu(bu, c)
}

func GetBuByName(c *gin.Context)  {
	name := c.Param("name")
	service.FindBuByName(name, c)

}

func GetAllBu(c *gin.Context){
	service.FindAllBu(c)

}

func UpdateBuActive(c *gin.Context)  {
	var bu model.Bu
	err := c.BindJSON(&bu)


	if err != nil {
		c.JSON(http.StatusOK,gin.H{"status":"Error","message":err.Error()})
	}else {
		result,err := service.UpdateBuActive(bu)
		if err != nil {
			c.JSON(http.StatusInternalServerError,gin.H{"status":"Error","message":err.Error()})
		}else {
			c.JSON(http.StatusOK,gin.H{"status":"Success","message":result})
		}
	}
}

func SquadToBu(c *gin.Context)  {
	var squadBu model.SquadBu
	err := c.BindJSON(&squadBu)
	if err != nil {
		c.JSON(http.StatusOK,gin.H{"status":"Error","message":err.Error()})
	}else {
		result,err:=service.InsertSquadToBu(squadBu)
		if err != nil {
			c.JSON(http.StatusInternalServerError,gin.H{"status":"Error","message":err.Error()})
		}else {
			if result.Squads != nil {
				c.JSON(http.StatusOK, gin.H{"status":"Success", "message":result})
			}else{
				c.JSON(503, gin.H{"status":"Error", "message":"can't be dublicated"})
			}
		}
	}

}

func BuDeactivate(c *gin.Context)  {
	var buDeactivate model.BuDeactive
	err := c.BindJSON(&buDeactivate)
	if err != nil{
		c.JSON(http.StatusOK,gin.H{"status":"Error","message":err.Error()})
	}else{
		result,err:=service.BuDeacivate(buDeactivate)
		fmt.Println(">>>>>>",result)
		fmt.Println("Devs>>>>>>",result.Squads)
		if err != nil {
			c.JSON(http.StatusInternalServerError,gin.H{"status":"Error","message":err.Error()})
		}else {
			if result.Squads != nil {
				c.JSON(http.StatusOK, gin.H{"status":"Success", "message":result})
			}else{
				c.JSON(503, gin.H{"status":"Error", "message":"can't be dublicated"})
			}
		}
	}
}

