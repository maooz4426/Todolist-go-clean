package infrastructure

import (
	"github.com/labstack/echo/v4"
	"github.com/maooz4426/Todolist/interfaces/controllers"
)

func NewRouter(hnd *controllers.TodoHandler) {
	e := echo.New()

	e.GET("/create", hnd.CreateHandler)
}
