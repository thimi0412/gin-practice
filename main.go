package main

import (
	"github.com/gin-gonic/gin"
	"github.com/thimi0412/gin-practice/controller"
)

func main() {
	r := gin.Default()
	userController := controller.UserController{}
	r.POST("/test", userController.JWTTest)
	r.POST("/signup", userController.Signup)
	r.POST("/signin", userController.Signin)
	r.Run()
}
