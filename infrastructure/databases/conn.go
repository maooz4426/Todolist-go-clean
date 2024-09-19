package databases

import (
	"github.com/maooz4426/Todolist/domain/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

func ConnectDB() (*gorm.DB, error) {
	var db *gorm.DB
	var err error

	dsn := "user:password@tcp(todo_db:3306)/db?charset=utf8mb4&parseTime=True"

	count := 5

	for count > 1 {
		if db, err = gorm.Open(mysql.Open(dsn)); err != nil {
			time.Sleep(2 * time.Second)
			count--
			log.Printf("retry... count:%v\n", count)
			continue
		}
		break
	}

	err = db.AutoMigrate(&entity.Todo{})

	if err != nil {

		return nil, err
	}

	return db, nil
}