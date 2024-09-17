package controllers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/maooz4426/Todolist/domain/entity"
	"github.com/maooz4426/Todolist/usecases/port"
	"net/http"
	"time"
)

type TodoController struct {
	svc port.TodoUseCaser
}

type CreateRequest struct {
	Task     string `json:"task"`
	Deadline string `json:"deadline"`
}

func NewController(svc port.TodoUseCaser) *TodoController {
	return &TodoController{svc}
}

func (r TodoController) CreateController(c echo.Context) error {

	var req CreateRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	var task entity.Todo
	fmt.Println(&req)
	task.Task = req.Task
	var err error
	task.Deadline, err = time.Parse("2006-01-02", req.Deadline)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	task.Done = false

	err = r.svc.Create(&task)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, task)
}
