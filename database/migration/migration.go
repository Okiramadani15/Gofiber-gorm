package migration

import (
	"fmt"
	"go_fiber_gorm/database"
	"go_fiber_gorm/model/entity"
	"log"
)

func RunMigration() {
	err := database.DB.AutoMigrate(&entity.User{}, &entity.Book{}, &entity.Category{}, &entity.Photo{})
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Database Migrated")
}
