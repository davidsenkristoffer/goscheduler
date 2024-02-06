package main

import (
	"fmt"
	types "goschedule/types"
	"html/template"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	t := &types.Template{
		Templates: template.Must(template.ParseGlob("public/views/*.html")),
	}
	e.Renderer = t
	e.Static("/static", "static")

	e.GET("/", index)
	e.GET("/schedule", newSchedule)
	e.POST("/schedule/user", newUser)

	e.Logger.Fatal(e.Start(":1323"))
}

func index(c echo.Context) error {
	return c.Render(200, "index", "hello, world!")
}

func newSchedule(c echo.Context) error {
	return c.Render(200, "schedule", types.DummyData())
}

func newUser(c echo.Context) error {
	name := c.FormValue("name")
	user := *types.NewUser(name)
	fmt.Printf("User: %s", user.Name)
	return c.Render(201, "user", user)
}
