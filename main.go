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
		user, err := parseJsonRequest(r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		switch r.Method {
		case http.MethodPost:
			result, err := db.ExecContext(context.Background(), "INSERT INTO user(id, name, mail, create_user) VALUES(?, ?, ?, ?)",
				uuid.New().String(), user.Name, user.Mail, "admin")
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

func parseJsonRequest(r *http.Request) (*user, error) {
	if r.Header.Get("Content-Type") != "application/json" {
		return nil, fmt.Errorf("invalid header Content-Type: %s", r.Header.Get("Content-Type"))
	}

	length, err := strconv.Atoi(r.Header.Get("Content-Length"))
	if err != nil {
		return nil, fmt.Errorf("invalid header Content-Length: %s, Error: %s", r.Header.Get("Content-Length"), err.Error())
	}

	body := make([]byte, length)
	length, err = r.Body.Read(body)
	if err != nil && err != io.EOF {
		return nil, fmt.Errorf("invalid Body Error: %s", err.Error())
	}

	var u *user
	err = json.Unmarshal(body[:length], &u)
	if err != nil {
		return nil, fmt.Errorf("invalid Body Error: %s", err.Error())
	}

	return u, nil
}

type user struct {
	Name string `json:"user"`
	Mail string `json:"mail"`
}
