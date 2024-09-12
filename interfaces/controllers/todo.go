package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/maooz4426/Todolist/domain"
	"github.com/maooz4426/Todolist/domain/usecases"
	"net/http"
	"time"
)

type TodoHandler struct {
	svc usecases.TodoUseCaser
}

type CreateRequest struct {
	Task     string    `json:"task"`
	Deadline time.Time `json:"deadline"`
}

func NewHandler(svc usecases.TodoUseCaser) *TodoHandler {
	return &TodoHandler{svc}
}

func (r TodoHandler) CreateHandler(c echo.Context) error {

	var req CreateRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	var task domain.Todo
	task.Task = req.Task
	task.Deadline = req.Deadline
	task.Done = false

	err := r.svc.Create(c, &task)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, task)
}
