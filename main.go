package main

import (
	"fmt"
	"os"
	"strconv"
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
		for i, task := range tasks {
			status := " "
			if task.Done {
				status = "✓"
			}
			fmt.Printf("%d. [%s] %s\n", i+1, status, task.Description)
		}

	case "done":
		taskId, _ := strconv.Atoi(os.Args[2])
		tasks, _ := loadTasks()
		for i := 0; i <= len(tasks)-1; i++ {
			if i+1 == taskId && tasks[i].Done == false {
				tasks[i].Done = true
				saveTasks(tasks)
				fmt.Println("Done: ", tasks[i].Description)
				return
			}
		}
		fmt.Println("Task not found: ", taskId)
	}

}
