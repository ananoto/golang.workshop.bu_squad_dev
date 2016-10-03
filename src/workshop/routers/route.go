package routers

import (
	"github.com/gin-gonic/gin"
	"workshop/controller"
)

func Route() {
	r := gin.Default()
	//POST
	r.POST("/create_bu", controller.CreateBu)
	r.POST("/create_dev", controller.CreateDev)
	r.POST("/create_squad", controller.CreateSquad)
	//GET
	r.GET("/get/bu/:name", controller.GetBuByName)
	r.GET("/get/squad/:name", controller.GetSquadByName)
	r.GET("/get/bus", controller.GetAllBu)
	//PUT
	r.PUT("/active_bu", controller.UpdateBuActive)
	r.PUT("/active_dev", controller.UpdateDevActive)
	r.PUT("/squad_to_bu", controller.SquadToBu)
	r.PUT("/dev_to_squad", controller.DevToSquad)
	r.PUT("/squad_deactivate", controller.SquadDeactivate)
	r.PUT("/bu_deactivate", controller.BuDeactivate)
	r.Run(":8080") // listen an

	/*	Action
		1.create BU
		2.create Squad
		3.create dev
		4.insert squad to bu
		5.insert dev to squad
		6.Deactivate BU then deactivate squad under BU then deactivate dev under that squad
		7.Deactivate Squad then deactivate dev
		8.Display squad in any BU
		9.Display dev in any squad*/
}
