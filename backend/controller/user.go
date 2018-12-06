package controller

import "github.com/labstack/echo"

// HandleUser ... "/user"パスのルーティング
func HandleUser(g *echo.Group) {
	g.POST("/user", createUser)
	g.GET("/user/:id", getUser)
	g.DELETE("/user/:id", deleteUser)
}
