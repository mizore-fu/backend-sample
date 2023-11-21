package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Task struct {
	id string
	name string
}

func NewTask() *Task {
	return &Task{}
}

var tasks []*Task = []*Task{}

func main() {
	tasks = []*Task{
		{
			id: "1",
			name: "Eat",
		},
		{
			id: "2",
			name: "Sleep",
		},
		{
			id: "3",
			name: "Study",
		},
	}
	for _, task := range tasks {
		fmt.Println(task.name)
	}

	e := echo.New()
	e.GET("/", hello)
	e.Logger.Fatal(e.Start(":8080"))
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "hello")
}
