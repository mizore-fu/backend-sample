package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Task struct {
	ID string `json:"id"`
	Name string `json:"name"`
}

func NewTask() *Task {
	return &Task{}
}

var tasks []*Task = []*Task{}

func main() {
	tasks = []*Task{
		{
			ID: "1",
			Name: "Eat",
		},
		{
			ID: "2",
			Name: "Sleep",
		},
		{
			ID: "3",
			Name: "Study",
		},
	}

	e := echo.New()
	e.GET("/tasks", GetTasks)
	e.POST("/tasks", AddTask)
	e.DELETE("/tasks/:id", DeleteTask)
	e.Logger.Fatal(e.Start(":8080"))
}

func GetTasks(c echo.Context) error {
	return c.JSON(http.StatusOK, tasks)
}

//http://localhost:8080/tasks
//body: {"id":"5", "name": "aaa"}
func AddTask(c echo.Context) error {
	newTask := NewTask()
	if err := c.Bind(newTask); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	tasks = append(tasks, newTask)
	return c.JSON(http.StatusCreated, newTask)
}

//http://localhost:8080/tasks/1
func DeleteTask(c echo.Context) error {
	id := c.Param("id")

	position := -1
	for i, task := range tasks {
		if task.ID == id {
			position = i
			break
		}
	}
	if position == -1 {
		return c.NoContent(http.StatusNotFound)
	}

	tasks[position] = tasks[len(tasks)-1]
	newTasks := tasks[:len(tasks)-1]
	tasks = newTasks
	return c.NoContent(http.StatusNoContent)
}
