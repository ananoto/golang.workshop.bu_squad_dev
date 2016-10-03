package route

import (
	"github.com/gin-gonic/gin"
	"workshop.test/controller"
	"workshop.test/dao"
)

func Route() {

	r := gin.Default()
	//Create
	r.POST("/create/dev", controller.CreateDev)
	r.POST("/create/squad", controller.CreateSquad)
	r.POST("/create/bu", controller.CreateBu)

	//Update
	r.POST("/insert/dev_to_squad", controller.InsertDevToSquad)
	r.POST("/insert/squad_to_bu", controller.InsertSquadToBu)
	r.POST("/deactive/squad", controller.DeactiveSquad)
	r.POST("/deactive/bu", controller.DeactiveBu)

	//Find
	r.POST("/find/dev", controller.FindDev)
	r.POST("/find/squad", controller.FindSquad)
	r.POST("/find/bu", controller.FindBu)

	//r.GET("/clear/database", func (c *gin.Context){dao.RessetAllDB()})
	//r.GET("/clear/database", dao.RessetAllDB)
	r.Run(":8080") // listen an
}

/*	Action
	--1.create BU
	--2.create Squad
	--3.create dev
	--4.insert squad to bu
	--5.insert dev to squad
	--6.Deactivate BU then deactivate squad under BU then deactivate dev under that squad
	--7.Deactivate Squad then deactivate dev
	--8.Display squad in any BU
	--9.Display dev in any squad*/
