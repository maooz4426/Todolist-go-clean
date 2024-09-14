package main

import (
	"github.com/maooz4426/Todolist/infrastructure"
	"github.com/maooz4426/Todolist/interfaces/controllers"
	"github.com/maooz4426/Todolist/interfaces/gateways/repository/mysql"
	"github.com/maooz4426/Todolist/usecases/interactor"
	"log"
)

func main() {
	db, err := mysql.NewDB()
	if err != nil {
		log.Fatal(err)
	}

	repo := mysql.NewTodoRepository(db)

	usc := interactor.NewTodoUseCase(repo)

	hnd := controllers.NewController(usc)

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
