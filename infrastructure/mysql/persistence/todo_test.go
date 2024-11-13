package persistence

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/maooz4426/Todolist/domain/entity"
	"github.com/maooz4426/Todolist/infrastructure/mysql"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
	"regexp"
	"testing"
	"time"
)

func TestInsert(t *testing.T) {

	sqlDB, mock, err := mysql.NewDbMock()
	require.NoError(t, err)

	ctx := context.Background()

	taskName := "test"
	done := false
	deadline, err := time.Parse("2006-01-02", "2024-10-11")
	require.NoError(t, err)
	var task *entity.Todo = &entity.Todo{Model: gorm.Model{ID: uint(1)}, Task: taskName, Done: done, Deadline: deadline}

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(
		"INSERT INTO `todos` (`created_at`,`updated_at`,`deleted_at`,`task`,`done`,`deadline`) VALUES (?,?,?,?,?,?)")).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), nil, taskName, done, deadline).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	m := NewTodoRepository(sqlDB)
	taskRes, err := m.Insert(ctx, task)
	require.NoError(t, err)
	require.Equal(t, uint(1), taskRes.ID)
	require.Equal(t, "test", task.Task)
	require.Equal(t, deadline, taskRes.Deadline)
}

func TestFindAll(t *testing.T) {
	sqlDB, mock, err := mysql.NewDbMock()
	require.NoError(t, err)

	ctx := context.Background()
	m := NewTodoRepository(sqlDB)

	taskName1 := "test"
	done1 := false
	deadline1, err := time.Parse("2006-01-02", "2024-10-11")
	require.NoError(t, err)
	taskName2 := "test2"
	done2 := false
	deadline2, err := time.Parse("2006-01-02", "2024-10-12")
	require.NoError(t, err)

	rows := sqlmock.NewRows([]string{"id", "task", "done", "deadline", "created_at", "updated_at", "deleted_at"}).
		AddRow(1, taskName1, done1, deadline1, time.Now(), time.Now(), nil).
		AddRow(2, taskName2, done2, deadline2, time.Now(), time.Now(), nil)

	mock.ExpectQuery(
		regexp.QuoteMeta("SELECT * FROM `todos` WHERE `todos`.`deleted_at` IS NULL")).
		WillReturnRows(rows)

	foundTasks, err := m.FindAll(ctx)
	require.NoError(t, err)
	require.Len(t, foundTasks, 2)
	require.Equal(t, taskName1, foundTasks[0].Task)
	require.Equal(t, taskName2, foundTasks[1].Task)
}

func TestFindById(t *testing.T) {

	sqlDB, mock, err := mysql.NewDbMock()
	require.NoError(t, err)

	taskName := "test"
	done := false
	deadline, _ := time.Parse("2006-01-02", "2024-10-11")
	rows := sqlmock.NewRows([]string{"id", "task", "done", "deadline", "created_at", "updated_at", "deleted_at"}).
		AddRow(1, taskName, done, deadline, time.Now(), time.Now(), nil)

	mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT * FROM `todos` WHERE id = ? AND `todos`.`deleted_at` IS NULL")).
		WithArgs("1").
		WillReturnRows(rows)

	m := NewTodoRepository(sqlDB)
	ctx := context.Background()

	foundTask, err := m.FindById(ctx, "1")
	require.NoError(t, err)
	require.Equal(t, uint(1), foundTask.ID)
	require.Equal(t, taskName, foundTask.Task)
	require.Equal(t, done, foundTask.Done)
	require.Equal(t, deadline, foundTask.Deadline)
}

func TestUpdate(t *testing.T) {
	sqlDB, mock, err := mysql.NewDbMock()
	require.NoError(t, err)

	ctx := context.Background()
	mr := NewTodoRepository(sqlDB)

	deadline, _ := time.Parse("2006-01-02", "2024-10-11")
	task := &entity.Todo{Model: gorm.Model{ID: 1}, Task: "test", Done: false, Deadline: deadline}

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE `todos` SET").
		WithArgs(
			sqlmock.AnyArg(),
			sqlmock.AnyArg(),
			nil,
			"updated",
			true,
			deadline,
			1,
		).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	rows := sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "task", "done", "deadline"}).
		AddRow(1, time.Now(), time.Now(), nil, "updated", true, deadline)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `todos` WHERE id = ? AND `todos`.`deleted_at` IS NULL")).
		WithArgs("1").
		WillReturnRows(rows)

	task.Task = "updated"
	task.Done = true
	updatedTask, err := mr.Update(ctx, task)
	require.NoError(t, err)

	foundTask, err := mr.FindById(ctx, "1")
	require.NoError(t, err)
	require.Equal(t, updatedTask.Task, foundTask.Task)
	require.Equal(t, updatedTask.Done, foundTask.Done)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestDelete(t *testing.T) {
	sqlDB, mock, err := mysql.NewDbMock()
	require.NoError(t, err)

	ctx := context.Background()
	mr := NewTodoRepository(sqlDB)

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE `todos` SET").
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `todos` WHERE id = ? AND `todos`.`deleted_at` IS NULL")).
		WithArgs("1").
		WillReturnError(gorm.ErrRecordNotFound)

	err = mr.Delete(ctx, "1")
	require.NoError(t, err)

	_, err = mr.FindById(ctx, "1")
	require.Error(t, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}
