package interactor

import (
	"context"
	"github.com/maooz4426/Todolist/domain/entity"
	mockrepository "github.com/maooz4426/Todolist/domain/mockreposiotry"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"gorm.io/gorm"
	"testing"
	"time"
)

func TestCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()

	mrTodo := mockrepository.NewMockITodoRepository(ctrl)
	mrTx := mockrepository.NewMockITransactionManager(ctrl)

	usc := NewTodoUseCase(mrTodo, mrTx)

	deadline, err := time.Parse("2006-01-02", "2024-10-11")
	require.NoError(t, err)
	task := entity.Todo{
		Task:     "test",
		Deadline: deadline,
	}

	//mrTx.EXPECT().RunInTx(gomock.Any(), mrTodo.EXPECT().Insert(gomock.Any(), gomock.Any()).Return(&entity.Todo{Task: "test", Done: false, Deadline: deadline}, nil))
	//
	//taskRes, err := usc.Create(ctx, &task)

	expectedTodo := &entity.Todo{Task: "test", Done: false, Deadline: deadline}
	mrTx.EXPECT().RunInTx(gomock.Any(), gomock.Any()).DoAndReturn(
		func(_ context.Context, f func(context.Context) error) error {
			return f(ctx)
		},
	)

	mrTodo.EXPECT().Insert(gomock.Any(), gomock.Any()).Return(expectedTodo, nil)
	mrTodo.EXPECT().FindById(gomock.Any(), gomock.Any()).Return(expectedTodo, nil)

	taskRes, err := usc.Create(ctx, &task)

	require.NoError(t, err)
	require.NotNil(t, taskRes)
	require.Equal(t, "test", taskRes.Task)
	require.False(t, false, taskRes.Done)
	require.Equal(t, deadline, taskRes.Deadline)
}

func TestUpdate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()

	mrTodo := mockrepository.NewMockITodoRepository(ctrl)
	mrTx := mockrepository.NewMockITransactionManager(ctrl)
	usc := NewTodoUseCase(mrTodo, mrTx)

	mrTx.EXPECT().RunInTx(gomock.Any(), gomock.Any()).DoAndReturn(
		func(_ context.Context, f func(context.Context) error) error {
			return f(ctx)
		},
	)

	deadline, err := time.Parse("2006-01-02", "2024-10-11")
	require.NoError(t, err)
	var task *entity.Todo = &entity.Todo{Task: "test", Done: true, Deadline: deadline}

	mrTodo.EXPECT().FindById(gomock.Any(), gomock.Any()).Return(task, nil)
	mrTodo.EXPECT().Update(gomock.Any(), gomock.Any()).Return(task, nil)
	taskRes, err := usc.Update(ctx, task)

	require.NoError(t, err)
	require.NotNil(t, taskRes)
	require.Equal(t, "test", taskRes.Task)
	require.Equal(t, true, taskRes.Done)
	require.Equal(t, deadline, taskRes.Deadline)

}

func TestFindAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()

	mrTodo := mockrepository.NewMockITodoRepository(ctrl)
	mrTx := mockrepository.NewMockITransactionManager(ctrl)
	usc := NewTodoUseCase(mrTodo, mrTx)

	expectedTodos := []*entity.Todo{
		{Task: "task1", Done: false},
		{Task: "task2", Done: true},
	}

	mrTodo.EXPECT().FindAll(gomock.Any()).Return(expectedTodos, nil)

	todos, err := usc.FindAll(ctx)

	require.NoError(t, err)
	require.NotNil(t, todos)
	require.Len(t, todos, 2)
	require.Equal(t, expectedTodos, todos)
}

func TestFindById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()

	mrTodo := mockrepository.NewMockITodoRepository(ctrl)
	mrTx := mockrepository.NewMockITransactionManager(ctrl)
	usc := NewTodoUseCase(mrTodo, mrTx)

	expectedTodo := &entity.Todo{
		Task: "test task",
		Done: false,
	}

	mrTodo.EXPECT().FindById(gomock.Any(), "1").Return(expectedTodo, nil)

	todo, err := usc.FindById(ctx, "1")

	require.NoError(t, err)
	require.NotNil(t, todo)
	require.Equal(t, expectedTodo, todo)
}

func TestDelete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()

	mrTodo := mockrepository.NewMockITodoRepository(ctrl)
	mrTx := mockrepository.NewMockITransactionManager(ctrl)
	usc := NewTodoUseCase(mrTodo, mrTx)

	mrTx.EXPECT().RunInTx(gomock.Any(), gomock.Any()).DoAndReturn(
		func(_ context.Context, f func(context.Context) error) error {
			return f(ctx)
		},
	)

	deadline, err := time.Parse("2006-01-02", "2024-10-11")
	require.NoError(t, err)
	var task *entity.Todo = &entity.Todo{Model: gorm.Model{ID: uint(1)}, Task: "test", Done: true, Deadline: deadline}

	mrTodo.EXPECT().FindById(gomock.Any(), "1").Return(task, nil)
	mrTodo.EXPECT().Delete(gomock.Any(), "1").Return(nil)

	deletedTask, err := usc.Delete(ctx, "1")

	require.NoError(t, err)
	require.NotNil(t, deletedTask)
	require.Equal(t, uint(1), deletedTask.ID)
	require.Equal(t, "test", deletedTask.Task)
	require.Equal(t, true, deletedTask.Done)
	require.Equal(t, deadline, deletedTask.Deadline)
}
