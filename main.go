package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Task struct {
	Id string `json:"id"`
	Name string `json:"name"`
}

func NewTask() *Task {
	return &Task{}
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
	e.POST("/add", AddTask)
	e.Logger.Fatal(e.Start(":8080"))
}

func GetTasks(c echo.Context) error {
	return c.JSON(http.StatusOK, tasks)
}

//http://localhost:8080/add
//body: {"id":"5", "name": "aaa"}
func AddTask(c echo.Context) error {
	newTask := NewTask()
	if err := c.Bind(newTask); err != nil {
		return err
	}
	tasks = append(tasks, newTask)
	return c.JSON(http.StatusOK, newTask)
}
