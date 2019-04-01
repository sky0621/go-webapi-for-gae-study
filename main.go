package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

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

	// --------------------------------------------------------------
	// ユーザー登録を行うWebAPIの定義
	// --------------------------------------------------------------
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("[GWFGS] /users")

		if r.Header.Get("Content-Type") != "application/json" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		length, err := strconv.Atoi(r.Header.Get("Content-Length"))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		body := make([]byte, length)
		length, err = r.Body.Read(body)
		if err != nil && err != io.EOF {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		var jsonBody map[string]interface{}
		err = json.Unmarshal(body[:length], &jsonBody)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		switch r.Method {
		case http.MethodPost:
			fmt.Println("[GWFGS] POST")

			result, err := db.ExecContext(context.Background(), "INSERT INTO user(id, name, mail, create_user) VALUES(?, ?, ?, ?)",
				uuid.New().String(), jsonBody["name"], jsonBody["mail"], "admin")
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

	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	fmt.Printf("[GWFGS] PORT: %s\n", port)
	if isLocal {
		port = ":8080"
	}
	if err := http.ListenAndServe(port, nil); err != nil {
		panic(err)
	}
}

func errWrapN(n int, err error) {
	panic(err)
}

func errWrap(err error) {
	panic(err)
}
