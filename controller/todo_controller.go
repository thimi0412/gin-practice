package controller

import (
	"time"

	"github.com/thimi0412/gin-practice/db"
	"github.com/thimi0412/gin-practice/entity"
)

// TodoController todoコントローラー
type TodoController struct{}

// CreateTodo action: POST /todo
func (tc TodoController) CreateTodo(user entity.User, context string, limitDate string) (entity.Todo, error) {
	conn := db.DBConnect()
	defer conn.Close()

	timeformat := "2006-01-02 15:04:05"
	loc, _ := time.LoadLocation("Asia/Tokyo")

	todo := entity.Todo{}

	t, err := time.ParseInLocation(timeformat, limitDate, loc)
	if err != nil {
		return todo, err
	}

	todo.UserID = user.ID
	todo.Context = context
	todo.LimitDate = &t

	if err := conn.Create(&todo).Error; err != nil {
		return todo, err
	}
	return todo, nil
}
