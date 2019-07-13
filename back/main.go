package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"./model"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := gorm.Open("sqlite3", "test/test.db")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	db.AutoMigrate(
		&model.ManagerAccount{},
		&model.CommunityAccount{},
		&model.Location{},
		&model.Space{},
		&model.Community{},
		&model.CommunityUser{},
		&model.Thread{},
		&model.Post{},
	)

	router := gin.Default()

	initObjects(db)
	initRouting(router)

	router.Run(":5622")
}
