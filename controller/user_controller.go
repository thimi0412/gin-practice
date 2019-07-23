package controller

import (
	"errors"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/thimi0412/gin-practice/authentication"
	"github.com/thimi0412/gin-practice/db"
	"github.com/thimi0412/gin-practice/entity"
)

// UserController ユーザコントローラー
type UserController struct{}

// Signup action: POST /signup
func (uc UserController) Signup(c *gin.Context) {
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

// Signin action POST /signin
func (uc UserController) Signin(c *gin.Context) {
	conn := db.DBConnect()
	defer conn.Close()

	email := c.PostForm("email")
	password := c.PostForm("password")

	user := entity.User{}

	if err := existUser(email, password); err != nil {
		log.Println(err)
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
	} else {
		jwt, err := authentication.CreateTokenString(user)
		if err != nil {
			log.Println(err)
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
		} else {
			c.JSON(200, gin.H{
				"jwt": jwt,
			})
		}
	}
}

func existUser(email, password string) error {
	conn := db.DBConnect()
	defer conn.Close()

	user := entity.User{}

	if err := conn.Where("email = ?", email).First(&user).Error; gorm.IsRecordNotFoundError(err) {
		return errors.New("No User")
	}
	if err := authentication.PasswordVerify(user.Password, password); err != nil {
		return errors.New("Password does not match")
	}
	return nil
}

// JWTTest action post /test
func (uc UserController) JWTTest(c *gin.Context) {
	token := c.PostForm("token")
	if user, err := authentication.AuthTokenString(token); err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"message": "success!",
			"user":    user,
		})
	}
}
