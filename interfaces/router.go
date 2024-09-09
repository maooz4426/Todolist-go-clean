package interfaces

import (
	"github.com/labstack/echo/v4"
	"github.com/maooz4426/Todolist/interfaces/handlers"
)

func NewRouter(hnd *handlers.TodoHandler) {
	e := echo.New()

	e.GET("/create", hnd.CreateHandler)
}
