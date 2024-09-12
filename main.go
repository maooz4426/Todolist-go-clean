package main

import (
	"github.com/maooz4426/Todolist/application/usecases"
	"github.com/maooz4426/Todolist/infrastructure"
	"github.com/maooz4426/Todolist/infrastructure/database"
	"github.com/maooz4426/Todolist/infrastructure/repository"
	"github.com/maooz4426/Todolist/interfaces/controllers"
	"log"
)

func main() {
	db, err := database.NewDB()
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewTodoRepository(db)

	usc := usecases.NewTodoUseCase(repo)

	hnd := controllers.NewHandler(usc)

	infrastructure.NewRouter(hnd)

}

//func main() {
//	e := echo.New()
//
//	e.GET("/", hello)
//	e.Logger.Fatal(e.Start(":8080"))
//
//}
//func hello(c echo.Context) error {
//	return c.String(http.StatusOK, "Hello, World!")
//}
