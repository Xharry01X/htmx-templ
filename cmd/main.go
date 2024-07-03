package main

import (
	"github.com/harshit-1245/htmx-templ/handler"
	"github.com/labstack/echo/v4"
)

type DB struct {

}

func main()  {
	app := echo.New()
	//database configure
	// var db DB

	//routes of all api
	userHandler := handler.UserHandler{}

	app.GET("/user",userHandler.HandlerUserShow)


	app.Start(":4000")
}