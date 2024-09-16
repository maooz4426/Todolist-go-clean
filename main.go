package main

import (
	"github.com/maooz4426/Todolist/infrastructure/databases"
	"github.com/maooz4426/Todolist/infrastructure/router"
	"github.com/maooz4426/Todolist/interfaces/controllers"
	"github.com/maooz4426/Todolist/interfaces/gateways/repository/datasource"
	"github.com/maooz4426/Todolist/usecases/interactor"
	"log"
)

func main() {
	db, err := databases.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	repo := datasource.NewTodoRepository(db)

	usc := interactor.NewTodoUseCase(repo)

	hnd := controllers.NewController(usc)

	router.NewRouter(hnd)

}
