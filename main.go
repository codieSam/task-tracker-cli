package main

import (
	"encoding/json"
	"fmt"
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

	command := os.Args[1]

	if command == "list" {

	}

	now := time.Now().Format(time.RFC3339)

	var tasks []Task

	data, err := os.ReadFile("tasks.json")
	if err == nil {
		err = json.Unmarshal(data, &tasks)
		if err != nil {
			fmt.Println("Error reading tasks", err)
			return
		}
		command := os.Args[1]

		if command == "list" {
			for _, task := range tasks {
				fmt.Println("id :", task.ID)
				fmt.Println("description :", task.Description)
				fmt.Println("status :", task.Status)
			}
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
	if command == "add" {
		if len(os.Args) < 3 {
			fmt.Println("Please provide description as well")
			return
		}
		desc := os.Args[2]
		task := Task{
			ID:          newId,
			Description: desc,
			Status:      "todo",
			CreatedAt:   now,
			UpdatedAt:   now,
		}
		tasks = append(tasks, task)
	}

	data, err = json.MarshalIndent(tasks, "", " ")

	if err != nil {
		fmt.Println("Error marshalling tasks: ", err)
		return
	}

	err = os.WriteFile("tasks.json", data, 0644)
	if err != nil {
		fmt.Println("Error while printing the line", err)
		return
	}

	// fmt.Println("ID: ", task.ID)
	// fmt.Println("Description: ", task.Description)
	// fmt.Println("Status: ", task.Status)
	// fmt.Println("Created At: ", task.CreatedAt)
	// fmt.Println("Updated At: ", task.UpdatedAt)

}
