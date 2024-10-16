package main

import (
	"github.com/maooz4426/Todolist/infrastructure/mysql"
	"github.com/maooz4426/Todolist/infrastructure/mysql/repository"
	"github.com/maooz4426/Todolist/infrastructure/router"
	"github.com/maooz4426/Todolist/interfaces/controllers"
	"github.com/maooz4426/Todolist/usecases/interactor"
	"log"
)

func main() {
	db, err := mysql.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewTodoRepository(db)

	txm := repository.NewTransactionManager(db)

	usc := interactor.NewTodoUseCase(repo, txm)

	hnd := controllers.NewController(usc)

	router.NewRouter(hnd)

}
