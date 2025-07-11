package main

// create the structure for tasks
type Task struct {
	Description string `json:"description"`
	Done        bool   `json:"done"`
	DueDate     string `json:"due_date"`
}
