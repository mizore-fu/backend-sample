package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Task struct {
	Id string `json:"id"`
	Name string `json:"name"`
}

var tasks []*Task = []*Task{}

func main() {
	tasks = []*Task{
		{
			Id: "1",
			Name: "Eat",
		},
		{
			Id: "2",
			Name: "Sleep",
		},
		{
			Id: "3",
			Name: "Study",
		},
	}

	e := echo.New()
	e.GET("/tasks", GetTasks)
	e.Logger.Fatal(e.Start(":8080"))
}

func GetTasks(c echo.Context) error {
	return c.JSON(http.StatusOK, tasks)
}
