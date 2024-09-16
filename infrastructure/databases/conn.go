package databases

import (
	"github.com/maooz4426/Todolist/domain/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	dsn := "user:password@tcp(127.0.0.1:3306)/db?charset=utf8mb4&parseTime=True"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&entity.Todo{})

	if err != nil {
		return nil, err
	}

	return db, nil
}
