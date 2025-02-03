package cmd

import (
	"encoding/json"
	"os"
	"sort"
)

type Task struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Done     bool   `json:"done"`
	Priority int    `json:"priority"`
}

const todosFile = "todos.json"

func loadTasks() ([]Task, error) {
	var tasks []Task
	data, err := os.ReadFile(todosFile)

	if err != nil {
		if os.IsNotExist(err) {
			return tasks, nil
		}
		return nil, err
	}
	err = json.Unmarshal(data, &tasks)
	return tasks, err
}

func saveTasks(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", " ")

	if err != nil {
		return err
	}
	return os.WriteFile(todosFile, data, 0644)
}

func genNextId(tasks []Task) int {
	if len(tasks) == 0 {
		return 1
	}

	sort.Slice(tasks, func(i, j int) bool { return tasks[i].ID < tasks[j].ID })
	return tasks[len(tasks)-1].ID + 1
}
