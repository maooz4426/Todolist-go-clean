package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/maooz4426/Todolist/domain/entity"
	mock_port "github.com/maooz4426/Todolist/domain/mockport"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestCreateController(t *testing.T) {
	ctrl := gomock.NewController(t)

	mp := mock_port.NewMockITodoUseCase(ctrl)

	deadline, err := time.Parse("2006-01-02", "2024-10-11")
	require.NoError(t, err)
	mocktask := entity.Todo{
		Task:     "test",
		Deadline: deadline,
		Done:     false,
	}
	mp.EXPECT().Create(gomock.Any(), mocktask).Return(mocktask, gomock.Any())

	ctr := NewController(mp)
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	ctx := echo.New().NewContext(req, rec)
	err = ctr.CreateController(ctx)

	require.NoError(t, err)
}
