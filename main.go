package main

import (
	"encoding/json"
	"os"
	"time"
)

type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

func main() {
	// if len(os.Args) < 2 {
	// 	fmt.Println("Please provide a command")
	// 	return
	// }

	// if len(os.Args) >= 3 {
	// 	arg := os.Args[2]

	// 	if os.Args[1] == "add" {
	// 		fmt.Println("Adding task: ", arg)
	// 	}

	// }

	desc := os.Args[2]
	now := time.Now().Format(time.RFC3339)

	var tasks []Task

	data, err := os.ReadFile("tasks.json")
	if err == nil {
		err = json.Unmarshal(data, &tasks)
		if err != nil {
			println("Error reading tasks", err)
			return
		}

	} else if !os.IsNotExist(err) {
		println("Error opening task file", err)
	}

	var newId int

	if len(tasks) == 0 {
		newId = 1
	} else {
		lastTask := tasks[len(tasks)-1]
		newId = lastTask.ID + 1
	}

	task := Task{
		ID:          newId,
		Description: desc,
		Status:      "todo",
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	tasks = append(tasks, task)

	data, err = json.MarshalIndent(tasks, "", " ")

	if err != nil {
		println("Error marshalling tasks: ", err)
		return
	}

	err = os.WriteFile("tasks.json", data, 0644)
	if err != nil {
		println("Error writing tasks to file: ", err)
		return
	}

	// fmt.Println("ID: ", task.ID)
	// fmt.Println("Description: ", task.Description)
	// fmt.Println("Status: ", task.Status)
	// fmt.Println("Created At: ", task.CreatedAt)
	// fmt.Println("Updated At: ", task.UpdatedAt)

}
