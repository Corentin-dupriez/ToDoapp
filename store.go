package main

import (
	"encoding/json"
	"os"
)

const dataFile = "tasks.json"

func saveTasks(tasks []Task) error {
	file, err := os.Create(dataFile)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)

	return encoder.Encode(tasks)
}

func loadTasks() ([]Task, error) {
	var tasks []Task

	file, err := os.Open(dataFile)
	if os.IsNotExist(err) {
		//no file yet
		return tasks, nil
	} else if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&tasks)
	return tasks, err

}
