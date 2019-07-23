package main

import (
	"github.com/gin-gonic/gin"
	"github.com/thimi0412/gin-practice/controller"
)

func main() {
	r := gin.Default()
	ctrl := controller.UserController{}
	r.POST("/test", ctrl.GetTest)
	r.POST("/user", ctrl.Create)
	r.POST("/login", ctrl.Login)
	r.Run()
}
