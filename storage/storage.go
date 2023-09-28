package storage

import (
    "encoding/json"
    "io/ioutil"
    "task-manager-cli/task" // Import the 'task' package
)

// LoadTasks loads tasks from a JSON file.
func LoadTasks(filename string, tasks *[]task.Task) error {
    data, err := ioutil.ReadFile(filename)
    if err != nil {
        return err
    }
    if err := json.Unmarshal(data, tasks); err != nil {
        return err
    }
    return nil
}

// SaveTasks saves tasks to a JSON file.
func SaveTasks(filename string, tasks []task.Task) error {
    data, err := json.Marshal(tasks)
    if err != nil {
        return err
    }
    if err := ioutil.WriteFile(filename, data, 0644); err != nil {
        return err
    }
    return nil
}
