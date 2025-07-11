package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: todo [add|list|done]")
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		var dueDate string
		task := os.Args[2]
		if len(os.Args) == 4 {
			dueDate = os.Args[3]
		}
		tasks, _ := loadTasks()
		if dueDate == "" {
			tasks = append(tasks, Task{Description: task})
		} else {
			tasks = append(tasks, Task{Description: task, DueDate: dueDate})
		}
		saveTasks(tasks)
		fmt.Println("Added: ", task)

	case "list":
		tasks, _ := loadTasks()
		fmt.Println("Description", "|", "Status", "|", "Due date")
		for _, task := range tasks {
			fmt.Println(task.Description, "|", task.Done, "|", task.DueDate)
		}

	case "done":
		task := os.Args[2]
		tasks, _ := loadTasks()
		for i := len(tasks) - 1; i >= 0; i-- {
			if tasks[i].Description == task && tasks[i].Done == false {
				tasks[i].Done = true
				saveTasks(tasks)
				fmt.Println("Done: ", tasks[i].Description)
				return
			}
		}
		fmt.Println("Task not found: ", task)
	}

}
