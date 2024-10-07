package controllers

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/maooz4426/Todolist/domain/dto"
	"github.com/maooz4426/Todolist/domain/entity"
	"github.com/maooz4426/Todolist/usecases/port"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type TodoController struct {
	usc port.TodoUseCaser
}

type Controller interface {
	CreateController(c echo.Context) error
	GetAllController(c echo.Context) error
	GetDetailController(c echo.Context) error
	UpdateController(c echo.Context) error
	DeleteController(c echo.Context) error
}

func NewController(svc port.TodoUseCaser) *TodoController {
	return &TodoController{svc}
}

func (con TodoController) CreateController(c echo.Context) error {

	var req dto.TodoJson
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	var taskReq entity.Todo
	//fmt.Println(&req)
	taskReq.Task = req.Task
	var err error

	taskReq.Deadline, err = time.Parse("2006-01-02", req.Deadline)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	taskReq.Done = false

	//var res *dto.CreateResponse
	taskRes, err := con.usc.Create(&taskReq)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	res, err := taskRes.ConvertDTO()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, res)
}

func (con TodoController) GetAllController(c echo.Context) error {
	todos, err := con.usc.FindAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	var res []dto.TodoJson

	for _, todo := range todos {
		todoDto, err := todo.ConvertDTO()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		res = append(res, *todoDto)
	}

	return c.JSON(http.StatusOK, res)
}

func (con TodoController) GetDetailController(c echo.Context) error {
	id := c.Param("taskId")

	task, err := con.usc.FindById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, err.Error())
		}
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	res, err := task.ConvertDTO()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}

func (con TodoController) UpdateController(c echo.Context) error {

	id := c.Param("taskId")

	todo, err := con.usc.FindById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	var req dto.TodoJson
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	todo.Task = req.Task
	todo.Deadline, err = time.Parse("2006-01-02", req.Deadline)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	todo.Done = req.Done

	todo, err = con.usc.Update(todo)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	res, err := todo.ConvertDTO()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}

func (con TodoController) DeleteController(c echo.Context) error {
	id := c.Param("taskId")
	task, err := con.usc.Delete(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, task)
}
