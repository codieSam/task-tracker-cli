package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"
)

type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

func loadTasks() []Task {

	var tasks []Task

	data, err := os.ReadFile("tasks.json")
	if err == nil {
		err = json.Unmarshal(data, &tasks)
		if err != nil {
			fmt.Println("Error reading tasks", err)
			return tasks
		}

	} else if !os.IsNotExist(err) {
		println("Error opening task file", err)
	}

	return tasks
}

func saveTasks(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", " ")

	if err != nil {
		fmt.Println("Error while marshelling", err)
		return err
	}

	err = os.WriteFile("tasks.json", data, 0644)

	if err != nil {
		fmt.Println("Error while writing file", err)
		return err
	}
	return nil
}

func addTask(tasks []Task) ([]Task, error) {
	if len(os.Args) < 4 {
		fmt.Println("Please provide complete details")
		return tasks, nil
	}
	var newId int
	now := time.Now().Format(time.RFC3339)

	if len(tasks) == 0 {
		newId = 1
	} else {
		lastTask := tasks[len(tasks)-1]
		newId = lastTask.ID + 1
	}
	desc := os.Args[2]
	status := os.Args[3]
	if status != "todo" && status != "in-progress" && status != "done" {
		fmt.Println("Please provide a valid status")
		return tasks, nil
	}
	task := Task{
		ID:          newId,
		Description: desc,
		Status:      status,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	tasks = append(tasks, task)
	err := saveTasks(tasks)
	if err != nil {
		fmt.Println("An error occured", err)
		return tasks, err
	}
	return tasks, nil
}

func listTask(tasks []Task) {
	if len(os.Args) < 3 {
		fmt.Println(tasks)

	} else {
		condition := os.Args[2]

		if condition != "done" && condition != "todo" && condition != "in-progress" {
			fmt.Println("Please provide the valid condition")
			return
		}
		for _, task := range tasks {
			if len(os.Args) >= 3 {
				cond := os.Args[2]
				if task.Status == cond {
					fmt.Println(task)
				}
			}

		}
	}

}

func markDone(tasks []Task, taskID string) ([]Task, error) {

	id, err := strconv.Atoi(taskID)
	if err != nil {
		return tasks, err
	}
	found := false
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Status = "done"
			found = true
			break
		}

	}
	if !found {
		return tasks, errors.New("ID doesn't exist.")
	}
	err = saveTasks(tasks)
	if err != nil {
		return tasks, err
	}
	return tasks, nil
}

func markInProgress(tasks []Task, taskID string) ([]Task, error) {
	id, err := strconv.Atoi(taskID)
	if err != nil {
		return tasks, err
	}

	found := false

	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Status = "in-progress"
			found = true
			break
		}

	}
	if !found {
		return tasks, errors.New("ID doesn't exixt.")
	}
	err = saveTasks(tasks)
	if err != nil {
		return tasks, err
	}
	return tasks, nil
}

func updateTask(tasks []Task, newID string, description string) ([]Task, error) {
	id, err := strconv.Atoi(newID)

	if err != nil {
		return tasks, err
	}
	var found bool
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Description = description
			tasks[i].UpdatedAt = time.Now().Format(time.RFC3339)
			found = true
		}

	}
	if !found {
		return tasks, errors.New("ID doesn't exist.")
	}
	err = saveTasks(tasks)
	if err != nil {
		return tasks, err
	}
	return tasks, nil
}

func deleteTask(tasks []Task, newID string) ([]Task, error) {
	id, err := strconv.Atoi(newID)
	if err != nil {
		return tasks, err
	}
	found := false
	for i := range tasks {
		if tasks[i].ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			found = true
			fmt.Println("Task delete successfully")
			break
		}
	}
	if !found {
		return tasks, errors.New("ID doesn't exist")
	}
	err = saveTasks(tasks)
	if err != nil {
		return tasks, err
	}
	return tasks, nil
}

func main() {

	if len(os.Args) <= 1 {
		fmt.Println("Please provide at lease one command")
		return
	}

	command := os.Args[1]
	tasks := loadTasks()

	if command == "list" {
		listTask(tasks)

	} else if command == "add" {
		var err error
		tasks, err = addTask(tasks)
		if err != nil {
			fmt.Println("Error while adding task", err)
			return
		}

	} else if command == "mark-done" {
		var err error
		if len(os.Args) >= 3 {
			tasks, err = markDone(tasks, os.Args[2])
			if err != nil {
				fmt.Println("Error while marking done the task.", err)
				return
			}

		}

	} else if command == "mark-in-progress" {
		var err error
		if len(os.Args) >= 3 {
			tasks, err = markInProgress(tasks, os.Args[2])
			if err != nil {
				fmt.Println("Error while marking taska as in-progress.", err)
				return
			}
		}

	} else if command == "update" {
		var err error
		if len(os.Args) >= 4 {
			tasks, err = updateTask(tasks, os.Args[2], os.Args[3])
			if err != nil {
				fmt.Println("Error while updating the task", err)
				return
			}
		} else {
			fmt.Println("Please provide task ID and new description")
		}
	} else if command == "delete" {
		var err error
		if len(os.Args) >= 3 {
			tasks, err = deleteTask(tasks, os.Args[2])
			if err != nil {
				fmt.Println("Error while deleting the task", err)
				return
			}

		}

	} else {
		fmt.Println("Command not found,please provide a valid command")
	}

}

// fmt.Println("ID: ", task.ID)
// fmt.Println("Description: ", task.Description)
// fmt.Println("Status: ", task.Status)
// fmt.Println("Created At: ", task.CreatedAt)
// fmt.Println("Updated At: ", task.UpdatedAt)
