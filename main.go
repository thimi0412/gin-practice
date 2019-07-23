package main

import (
	"github.com/gin-gonic/gin"
	"github.com/thimi0412/gin-practice/controller"
)

func main() {
	r := gin.Default()
	ctrl := controller.UserController{}
	r.POST("/test", ctrl.JWTTest)
	r.POST("/signup", ctrl.Signup)
	r.POST("/signin", ctrl.Signin)
	r.Run()
}
