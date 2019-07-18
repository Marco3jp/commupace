package main

import (
	"google.golang.org/appengine"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/Marco3jp/commupace/back/model"
	"github.com/gin-gonic/gin"
	"os"
	"fmt"
)

func main() {
	appengine.Main()
	mysqlConnectionParams := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=%s",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_DOMAIN"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_NAME"),		"Asia%2FTokyo")
	db, err := gorm.Open("mysql", mysqlConnectionParams)
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
