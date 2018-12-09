package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"go-webapi-for-gae-study/backend/controller"
	"go-webapi-for-gae-study/backend/model"
	"google.golang.org/appengine"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main()  {
	var (
		connectionName = os.Getenv("CLOUDSQL_CONNECTION_NAME")
		user           = os.Getenv("CLOUDSQL_USER")
		password       = os.Getenv("CLOUDSQL_PASSWORD")
		database       = os.Getenv("CLOUDSQL_DATABASE")
	)
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@unix(/cloudsql/%s)/%s", user, password, connectionName, database))
	if appengine.IsDevAppServer() {
		db, err = gorm.Open("mysql", "testuser:testpass@tcp(127.0.0.1:3306)/testdb?charset=utf8&parseTime=True&loc=Local")
	}
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.LogMode(true)
	if err := db.DB().Ping(); err != nil {
		panic(err)
	}
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(10)

	db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8").AutoMigrate(&model.User{})

	e := echo.New()
	defer e.Close()

	http.Handle("/", e)
	g := e.Group("/api/v1")

	controller.NewUser(db).Handle(g)

	appengine.Main()
}
