package controller

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"go-webapi-for-gae-study/backend/service"
	"go-webapi-for-gae-study/backend/controller/form"
	"net/http"
)

// User ...
type User interface {
	Handle(g *echo.Group)
}

type user struct {
	db *gorm.DB
}

// NewUser ...
func NewUser(db *gorm.DB) User {
	return &user{db: db}
}

// HandleUser ...
func (u *user) Handle(g *echo.Group) {
	g.POST("/users", u.createUser)
	g.GET("/users/:id", u.getUser)
	g.PUT("/users/:id", u.updateUser)
	g.DELETE("/users/:id", u.deleteUser)
}

func (u *user) createUser(c echo.Context) error {
	uf := &form.User{}
	if err := c.Bind(uf); err != nil {
		return c.JSON(http.StatusBadRequest, errorJSON(http.StatusBadRequest, err.Error()))
	}

	id, err := service.NewUser(u.db).CreateUser(uf.ParseToDto())
	if err != nil {
		return c.JSON(http.StatusBadRequest, errorJSON(http.StatusBadRequest, err.Error()))
	}

	return c.JSON(http.StatusCreated, &form.ID{ID: id})
}

func (u *user) getUser(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, errorJSON(http.StatusBadRequest, "no parameter"))
	}

	mdl, err := service.NewUser(u.db).GetUser(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errorJSON(http.StatusInternalServerError, err.Error()))
	}

	if mdl.ID == "" {
		return c.NoContent(http.StatusNotFound)
	} else {
		return c.JSON(http.StatusOK, mdl)
	}
}

func (u *user) updateUser(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, errorJSON(http.StatusBadRequest, "no parameter"))
	}

	uf := &form.User{ID: id}
	if err := c.Bind(uf); err != nil {
		return c.JSON(http.StatusBadRequest, errorJSON(http.StatusBadRequest, err.Error()))
	}

	mdl, err := service.NewUser(u.db).UpdateUser(uf.ParseToDto())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errorJSON(http.StatusInternalServerError, err.Error()))
	}

	if mdl.ID == "" {
		return c.NoContent(http.StatusNotFound)
	} else {
		return c.JSON(http.StatusOK, mdl)
	}
}

func (u *user) deleteUser(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, errorJSON(http.StatusBadRequest, "no parameter"))
	}

	err := service.NewUser(u.db).DeleteUser(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errorJSON(http.StatusInternalServerError, err.Error()))
	}

	return c.NoContent(http.StatusNoContent)
}
