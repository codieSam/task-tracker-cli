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
	if len(os.Args) <= 1 {
		fmt.Println("Please provide at lease one command")
		return
	}
	command := os.Args[1]

	now := time.Now().Format(time.RFC3339)

	tasks := loadTasks()

	if command == "list" {

		for _, task := range tasks {
			if len(os.Args) >= 3 {
				cond := os.Args[2]
				if task.Status == cond {
					fmt.Println(task)
				}
			} else {
				fmt.Println(task)
			}

			// fmt.Println("id :", task.ID)
			// fmt.Println("description :", task.Description)
			// fmt.Println("status :", task.Status)
		}
	} else if command == "add" {
		if len(os.Args) < 4 {
			fmt.Println("Please provide complete details")
			return
		}
		var newId int

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
			return
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
	} else if command == "mark-done" {
		if len(os.Args) >= 3 {
			taskId := os.Args[2]
			id, err := strconv.Atoi(taskId)
			if err != nil {
				fmt.Println("Error", err)
				return
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
				return
			}
			saveTasks(tasks)
		}

	} else if command == "mark-in-progress" {
		if len(os.Args) >= 3 {
			newId := os.Args[2]
			id, err := strconv.Atoi(newId)
			if err != nil {
				fmt.Println("Error", err)
				return
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
				return
			}
			saveTasks(tasks)
		}

	} else if command == "update" {
		if len(os.Args) >= 4 {
			givenId := os.Args[2]
			newDesc := os.Args[3]
			id, err := strconv.Atoi(givenId)
			if err != nil {
				fmt.Println("Error while converting into int", err)
				return
			}
			for i := range tasks {
				if tasks[i].ID == id {
					tasks[i].Description = newDesc
					tasks[i].UpdatedAt = time.Now().Format(time.RFC3339)
				}

			}
			saveTasks(tasks)
		} else {
			fmt.Println("Please provide task ID and new description")
		}
	} else if command == "delete" {

		if len(os.Args) >= 3 {
			idToDelete := os.Args[2]
			id, err := strconv.Atoi(idToDelete)
			if err != nil {
				fmt.Println("There is an error while printing", err)
				return
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

		}
		saveTasks(tasks)
	} else {
		fmt.Println("Command not found,please provide a valid command")
	}

}

// fmt.Println("ID: ", task.ID)
// fmt.Println("Description: ", task.Description)
// fmt.Println("Status: ", task.Status)
// fmt.Println("Created At: ", task.CreatedAt)
// fmt.Println("Updated At: ", task.UpdatedAt)
