package database

import (
	"fmt"
	"testgo/models"
)

func RunMigration() {
	err := DB.AutoMigrate(
		&models.User{},
		&models.Product{},
	)

	if err != nil {
		fmt.Println(err)
		panic("Migration Failed")
	}

	fmt.Println("Migration Success")
}
