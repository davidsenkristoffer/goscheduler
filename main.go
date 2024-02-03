package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Static("/static", "static")
	e.File("/", "index.html")
	e.File("/contact/:id/edit", "editcontact.html")

	e.Logger.Fatal(e.Start(":1323"))
}
