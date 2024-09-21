package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/maooz4426/Todolist/domain/dto"
	"github.com/maooz4426/Todolist/domain/entity"
	"github.com/maooz4426/Todolist/usecases/port"
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

	return c.JSON(http.StatusOK, res)
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
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	res, err := task.ConvertDTO()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}
