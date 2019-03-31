package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/google/uuid"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// --------------------------------------------------------------
	// GCP - CloudSQL への接続情報は環境変数から取得
	// --------------------------------------------------------------
	var (
		connectionName = os.Getenv("CLOUDSQL_CONNECTION_NAME")
		user           = os.Getenv("CLOUDSQL_USER")
		password       = os.Getenv("CLOUDSQL_PASSWORD")
		database       = os.Getenv("CLOUDSQL_DATABASE")
	)
	dataSource := fmt.Sprintf("%s:%s@unix(/cloudsql/%s)/%s?parseTime=True", user, password, connectionName, database)

	// ローカル環境（docker-composeでMySQLを起動）での動作確認用
	isLocal := os.Getenv("IS_LOCAL") != ""
	if isLocal {
		dataSource = "testuser:testpass@tcp(localhost:3306)/testdb?charset=utf8&parseTime=True&loc=Local"
	}

	// --------------------------------------------------------------
	// GCP - CloudSQL へ接続
	// --------------------------------------------------------------
	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	defer func() {
		if db != nil {
			err := db.Close()
			if err != nil {
				panic(err)
			}
		}
	}()

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("[GWFGS] /users")
		switch r.Method {
		case http.MethodPost:
			fmt.Println("[GWFGS] POST")
			if err := r.ParseForm(); err != nil {
				if _, err := fmt.Fprintln(w, err.Error()); err != nil {
					panic(err)
				}
			}
			form := r.PostForm
			fmt.Printf("[GWFGS] PostForm: %#v\n", form)
			name := form.Get("name")
			fmt.Printf("[GWFGS] PostForm[name]: %s\n", name)
			mail := form.Get("mail")
			fmt.Printf("[GWFGS] PostForm[mail]: %s\n", mail)

			result, err := db.ExecContext(context.Background(), "INSERT INTO user(id, name, mail, create_user) VALUES(?, ?, ?, ?)",
				uuid.New().String(), name, mail, "admin")
			if err != nil {
				if _, err := fmt.Fprintln(w, err.Error()); err != nil {
					panic(err)
				}
			}
			if result == nil {
				panic(errors.New("result is nil"))
			}
		}
	})

	addr := ":80"
	if isLocal {
		addr = ":8080"
	}
	if err := http.ListenAndServe(addr, nil); err != nil {
		panic(err)
	}
}
