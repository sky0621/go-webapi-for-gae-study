package main

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"go-webapi-for-gae-study/backend/controller"
	"go-webapi-for-gae-study/backend/model"
	"google.golang.org/appengine"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main()  {
	// この接続情報は、あくまでローカル（docker-composeで立ち上げるMySQLへの接続限定）用
	// GAE環境にデプロイする時は、接続情報はもとより、形式も下記のように変える必要がある。
	// "【ユーザ】:【パスワード】@unix(/cloudsql/【インスタンス接続文字列（※おそらく「【プロジェクト名】:【リージョン】:【インスタンスID】」）】)/【DB名】"
	db, err := gorm.Open("mysql", "testuser:testpass@tcp(127.0.0.1:3306)/testdb?charset=utf8&parseTime=True&loc=Local")
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
