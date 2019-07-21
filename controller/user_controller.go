package controller

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/thimi0412/gin-practice/authentication"
	"github.com/thimi0412/gin-practice/db"
	"github.com/thimi0412/gin-practice/entity"
)

// UserController ユーザコントローラー
type UserController struct{}

// Create action: POST /user
func (uc UserController) Create(c *gin.Context) {
	conn := db.DBConnect()
	defer conn.Close()

	email := c.PostForm("email")
	password := c.PostForm("password")

	hashPw, err := authentication.PasswordHash(password)
	if err != nil {
		c.JSON(400, err)
	}

	user := entity.User{}
	user.Email = email
	user.Password = hashPw

	if err := conn.Create(&user).Error; err != nil {
		log.Println(err)
		c.JSON(400, err)
	} else {
		c.JSON(201, user)
	}
}

// GetSampleUser action: GET /user
func (uc UserController) GetSampleUser(c *gin.Context) {
	user := entity.User{}
	user.ID = 1
	user.Email = "kosuke@hogehoge.com"
	user.Password = "hogehuga"

	c.JSON(201, user)
}
