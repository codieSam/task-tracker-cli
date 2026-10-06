package main

import (
	"encoding/json"
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

func addTask(tasks []Task) []Task {
	if len(os.Args) < 4 {
		fmt.Println("Please provide complete details")
		return tasks
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
		return tasks
	}
	task := Task{
		ID:          newId,
		Description: desc,
		Status:      status,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	tasks = append(tasks, task)
	saveTasks(tasks)
	return tasks
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

func markDone(tasks []Task, taskID string) []Task {

	id, err := strconv.Atoi(taskID)
	if err != nil {
		fmt.Println("Error", err)
		return tasks
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
		fmt.Println("Provided ID doesn't exist")
		return tasks
	}
	saveTasks(tasks)
	return tasks
}

func markInProgress(tasks []Task, taskID string) []Task {
	id, err := strconv.Atoi(taskID)
	if err != nil {
		fmt.Println("Error", err)
		return tasks
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
		fmt.Println("Provided ID doesn't exist")
		return tasks
	}
	saveTasks(tasks)
	return tasks
}

func updateTask(tasks []Task, newID string, description string) []Task {
	id, err := strconv.Atoi(newID)
	if err != nil {
		fmt.Println("Error while converting into int", err)
		return tasks
	}
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Description = description
			tasks[i].UpdatedAt = time.Now().Format(time.RFC3339)
		}

	}
	saveTasks(tasks)
	return tasks
}

func deleteTask(tasks []Task, newID string) []Task {
	id, err := strconv.Atoi(newID)
	if err != nil {
		fmt.Println("There is an error while printing", err)
		return tasks
	}
	found := false
	for i := range tasks {
		if tasks[i].ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			found = true
			fmt.Printf("Task delete successfully")
			break
		}
	}
	if !found {
		fmt.Println("Tasks not found !")
	}
	saveTasks(tasks)
	return tasks
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
		tasks = addTask(tasks)

	} else if command == "mark-done" {
		if len(os.Args) >= 3 {
			tasks = markDone(tasks, os.Args[2])
		}

	} else if command == "mark-in-progress" {
		if len(os.Args) >= 3 {
			tasks = markInProgress(tasks, os.Args[2])
		}

	} else if command == "update" {
		if len(os.Args) >= 4 {
			updateTask(tasks, os.Args[2], os.Args[3])
		} else {
			fmt.Println("Please provide task ID and new description")
		}
	} else if command == "delete" {

		if len(os.Args) >= 3 {
			tasks = deleteTask(tasks, os.Args[2])
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
