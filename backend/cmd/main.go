package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"go-webapi-for-gae-study/backend/controller"
	"google.golang.org/appengine"
	"net/http"
)

func main()  {
	// https://echo.labstack.com/guide
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Secure())

	http.Handle("/", e)
	g := e.Group("/api/v1")

	controller.HandleUser(g)

	appengine.Main()
}
