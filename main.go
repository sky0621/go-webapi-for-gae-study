package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var (
		isLocal = os.Getenv("IS_LOCAL")
	)
	var (
		connectionName = os.Getenv("CLOUDSQL_CONNECTION_NAME")
		user           = os.Getenv("CLOUDSQL_USER")
		password       = os.Getenv("CLOUDSQL_PASSWORD")
		database       = os.Getenv("CLOUDSQL_DATABASE")
	)

	dataSource := fmt.Sprintf("%s:%s@unix(/cloudsql/%s)/%s?parseTime=True", user, password, connectionName, database)
	if isLocal != "" {
		dataSource = "testuser:testpass@tcp(localhost:3306)/testdb?charset=utf8&parseTime=True&loc=Local"
	}

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

	db.SetConnMaxLifetime(60 * time.Second)
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)
	if err := db.Ping(); err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer func() {
		if cancel != nil {
			cancel()
		}
	}()

	result, err := db.ExecContext(ctx, "INSERT INTO user(id, name, mail, create_user) VALUES(?, ?, ?, ?)",
		uuid.New().String(), "Taro", "taro@example.com", "dummy")
	if err != nil {
		panic(err)
	}
	if result == nil {
		panic(errors.New("result is nil"))
	}

	id, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	fmt.Printf("LastInsertId: %d\n", id)

	rows, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Printf("RowsAffected: %d\n", rows)
}
