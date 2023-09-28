package storage

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"task-manager-cli/task"
)

// TaskFilePath is the path to the JSON file where tasks are stored.
const TaskFilePath = "tasks.json"

// LoadTasks loads tasks from a JSON file.
func LoadTasks() ([]task.Task, error) {
	var tasks []task.Task

	// Check if the file exists
	_, err := os.Stat(TaskFilePath)
	if os.IsNotExist(err) {
		// File doesn't exist, return an empty task list
		return tasks, nil
	} else if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadFile(TaskFilePath)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

// SaveTasks saves tasks to a JSON file.
func SaveTasks(tasks []task.Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	dir := filepath.Dir(TaskFilePath)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		// Create the directory if it doesn't exist
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}

	if err := ioutil.WriteFile(TaskFilePath, data, 0644); err != nil {
		return err
	}

	fmt.Println("Tasks saved successfully.")
	return nil
}
