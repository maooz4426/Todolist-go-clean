package interactor

import (
	"context"
	"github.com/maooz4426/Todolist/domain/entity"
	mockrepository "github.com/maooz4426/Todolist/domain/mockreposiotry"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
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
