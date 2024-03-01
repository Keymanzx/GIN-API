package routes

import (
	ctl "api-gin/src/controller"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	api := router.Group("api")
	v1 := api.Group("v1")

	// Users
	user := v1.Group("user")
	user.GET("/all", ctl.GetAllUser)
	user.GET("/:id", ctl.GetByUserID)
	user.POST("/create", ctl.CreateUser)
	user.PUT("/update", ctl.UpdateUser)
	user.DELETE("/:id", ctl.DeleteByUserID)

	return router

}
