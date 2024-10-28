package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/maooz4426/Todolist/domain/entity"
	mock_port "github.com/maooz4426/Todolist/domain/mockport"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestCreateController(t *testing.T) {
	ctrl := gomock.NewController(t)
	mp := mock_port.NewMockITodoUseCase(ctrl)

	deadline, err := time.Parse("2006-01-02", "2024-10-11")
	require.NoError(t, err)

	expectedTodo := &entity.Todo{
		Task:     "test",
		Done:     false,
		Deadline: deadline,
	}

	returnTodo := &entity.Todo{
		Model:    gorm.Model{ID: 1},
		Task:     "test",
		Done:     false,
		Deadline: deadline,
	}

	mp.EXPECT().Create(gomock.Any(), expectedTodo).Return(returnTodo, nil)

	ctr := NewController(mp)
	reqBody := `{
        "task": "test",
        "deadline": "2024-10-11",
        "done": false
    }`
	req := httptest.NewRequest(http.MethodPost, "/task/create", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	ctx := echo.New().NewContext(req, rec)

	err = ctr.CreateController(ctx)
	require.NoError(t, err)
}

func TestGetAllController(t *testing.T) {
	ctrl := gomock.NewController(t)
	mp := mock_port.NewMockITodoUseCase(ctrl)

	deadline, err := time.Parse("2006-01-02", "2024-10-11")
	require.NoError(t, err)

	returnTodos := []*entity.Todo{
		{
			Model:    gorm.Model{ID: 1},
			Task:     "test1",
			Done:     false,
			Deadline: deadline,
		},
		{
			Model:    gorm.Model{ID: 2},
			Task:     "test2",
			Done:     true,
			Deadline: deadline,
		},
	}

	mp.EXPECT().FindAll(gomock.Any()).Return(returnTodos, nil)

	ctr := NewController(mp)
	req := httptest.NewRequest(http.MethodGet, "/tasks", nil)
	rec := httptest.NewRecorder()
	ctx := echo.New().NewContext(req, rec)

	err = ctr.GetAllController(ctx)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, rec.Code)
}

func TestGetDetailController(t *testing.T) {
	ctrl := gomock.NewController(t)
	mp := mock_port.NewMockITodoUseCase(ctrl)

	deadline, err := time.Parse("2006-01-02", "2024-10-11")
	require.NoError(t, err)

	returnTodo := &entity.Todo{
		Model:    gorm.Model{ID: 1},
		Task:     "test",
		Done:     false,
		Deadline: deadline,
	}

	mp.EXPECT().FindById(gomock.Any(), "1").Return(returnTodo, nil)

	ctr := NewController(mp)
	req := httptest.NewRequest(http.MethodGet, "/task/1", nil)
	rec := httptest.NewRecorder()

	e := echo.New()
	ctx := e.NewContext(req, rec)
	ctx.SetParamNames("taskId")
	ctx.SetParamValues("1")

	err = ctr.GetDetailController(ctx)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, rec.Code)
}

func TestUpdateController(t *testing.T) {
	ctrl := gomock.NewController(t)
	mp := mock_port.NewMockITodoUseCase(ctrl)

	deadline, err := time.Parse("2006-01-02", "2024-10-11")
	require.NoError(t, err)

	expectedTodo := &entity.Todo{
		Model:    gorm.Model{ID: 1},
		Task:     "updated test",
		Done:     true,
		Deadline: deadline,
	}

	returnTodo := &entity.Todo{
		Model:    gorm.Model{ID: 1},
		Task:     "updated test",
		Done:     true,
		Deadline: deadline,
	}

	mp.EXPECT().Update(gomock.Any(), expectedTodo).Return(returnTodo, nil)

	ctr := NewController(mp)
	reqBody := `{
        "task": "updated test",
        "deadline": "2024-10-11",
        "done": true
    }`

	req := httptest.NewRequest(http.MethodPut, "/task/1", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	e := echo.New()
	ctx := e.NewContext(req, rec)
	ctx.SetParamNames("taskId")
	ctx.SetParamValues("1")

	err = ctr.UpdateController(ctx)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, rec.Code)
}

func TestDeleteController(t *testing.T) {
	ctrl := gomock.NewController(t)
	mp := mock_port.NewMockITodoUseCase(ctrl)

	deadline, err := time.Parse("2006-01-02", "2024-10-11")
	require.NoError(t, err)

	returnTodo := &entity.Todo{
		Model:    gorm.Model{ID: 1},
		Task:     "test",
		Done:     false,
		Deadline: deadline,
	}

	mp.EXPECT().Delete(gomock.Any(), "1").Return(returnTodo, nil)

	ctr := NewController(mp)
	req := httptest.NewRequest(http.MethodDelete, "/task/1", nil)
	rec := httptest.NewRecorder()

	e := echo.New()
	ctx := e.NewContext(req, rec)
	ctx.SetParamNames("taskId")
	ctx.SetParamValues("1")

	err = ctr.DeleteController(ctx)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, rec.Code)
}
