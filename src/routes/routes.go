package routes

import (
	ctl "api-gin/src/controller"
	"api-gin/src/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	api := router.Group("api")
	v1 := api.Group("v1")

	// Auth
	auth := v1.Group("auth")
	auth.POST("/log-in", ctl.UserLogin)

	// Users
	user := v1.Group("user")
	user.GET("/all", middleware.JWTMiddleware(), ctl.GetAllUser)
	user.GET("/:id", middleware.JWTMiddleware(), ctl.GetByUserID)
	user.POST("/create", middleware.JWTMiddleware(), ctl.CreateUser)
	user.PUT("/update", middleware.JWTMiddleware(), ctl.UpdateUser)
	user.DELETE("/:id", middleware.JWTMiddleware(), ctl.DeleteByUserID)

	return router

}
