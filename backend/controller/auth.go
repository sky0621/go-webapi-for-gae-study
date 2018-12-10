package controller

import (
	"go-webapi-for-gae-study/backend/controller/form"
	"go-webapi-for-gae-study/backend/model"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

// Auth ...
type Auth interface {
	Handle(g *echo.Group)
}

type auth struct {
	db *gorm.DB
}

// NewAuth ...
func NewAuth(db *gorm.DB) Auth {
	return &auth{db: db}
}

// Handle ... "認証処理系パスのルーティング
func (a *auth) Handle(g *echo.Group) {
	g.POST("/login", a.login)
	g.POST("/logout", a.logout)
}

func (a *auth) login(c echo.Context) error {
	u := &form.Login{}

	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, errorJSON(http.StatusBadRequest, err.Error()))
	}

	dao := model.NewAuthDao(a.db)
	jwtToken, err := dao.Login(u.ParseToDto())
	if err != nil {
		return c.JSON(http.StatusBadRequest, errorJSON(http.StatusBadRequest, err.Error()))
	}

	return c.JSON(http.StatusOK, &form.ID{ID: jwtToken})
}

func (a *auth) logout(c echo.Context) error {
	u := &form.Logout{}

	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, errorJSON(http.StatusBadRequest, err.Error()))
	}

	dao := model.NewAuthDao(a.db)
	err := dao.Logout(u.ParseToDto())
	if err != nil {
		return c.JSON(http.StatusBadRequest, errorJSON(http.StatusBadRequest, err.Error()))
	}

	return c.NoContent(http.StatusNoContent)
}
