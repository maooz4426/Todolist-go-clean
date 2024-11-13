package main

import (
	"github.com/maooz4426/Todolist/infrastructure/mysql"
	"github.com/maooz4426/Todolist/infrastructure/mysql/persistence"
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

	repo := persistence.NewTodoRepository(db)

	txm := persistence.NewTransactionManager(db)

	usc := interactor.NewTodoUseCase(repo, txm)

	hnd := controllers.NewController(usc)

	router.NewRouter(hnd)

}
