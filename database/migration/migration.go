package migration

import (
	"fmt"

	"github.com/Maniexie/crud-fiber/database"
	"github.com/Maniexie/crud-fiber/models/entity"
)

func RunMigrate() {

	err := database.DB.AutoMigrate(&entity.User{})

	if err != nil {
		panic(err)
	}
	fmt.Println("Migrate Berhasil")

}
