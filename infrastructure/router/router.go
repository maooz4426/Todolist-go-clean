package router

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/maooz4426/Todolist/interfaces/controllers"
)

func NewRouter(cnt *controllers.TodoController) {
	e := echo.New()

	fmt.Println("starting server")

	e.POST("/tasks/create", cnt.CreateController)

	e.Logger.Fatal(e.Start(":8080"))
}
