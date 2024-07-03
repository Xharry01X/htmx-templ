package handler

import (
	"github.com/harshit-1245/htmx-templ/view/user"
	"github.com/labstack/echo/v4"
)


type UserHandler struct{
	Username string
	Password string
}

func (h UserHandler) HandlerUserShow(c echo.Context) error{
	return user.Show().Render(c.Request().Context(), c.Response())
}