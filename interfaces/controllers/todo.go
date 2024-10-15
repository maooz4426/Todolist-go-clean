package controllers

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/maooz4426/Todolist/domain/dto"
	"github.com/maooz4426/Todolist/domain/entity"
	"github.com/maooz4426/Todolist/usecases/port"
	"gorm.io/gorm"
	"net/http"
	"strconv"
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

	ctx := c.Request().Context()

	//var res *dto.CreateResponse
	taskRes, err := con.usc.Create(ctx, &taskReq)
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
	ctx := c.Request().Context()

	todos, err := con.usc.FindAll(ctx)
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

	ctx := c.Request().Context()
	task, err := con.usc.FindById(ctx, id)
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

	ctx := c.Request().Context()
	//todo, err := con.usc.FindById(ctx, id)
	//if err != nil {
	//	return c.JSON(http.StatusInternalServerError, err.Error())
	//}

	var req dto.TodoJson
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	var todo entity.Todo

	todoID, err := strconv.ParseUint(id, 10, 64)
	todo.ID = uint(todoID)
	if err != nil {
		return err
	}
	todo.Task = req.Task
	todo.Deadline, err = time.Parse("2006-01-02", req.Deadline)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	todo.Done = req.Done
	todoRes, err := con.usc.Update(ctx, &todo)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	res, err := todoRes.ConvertDTO()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}

func (con TodoController) DeleteController(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("taskId")
	task, err := con.usc.Delete(ctx, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, task)
}
