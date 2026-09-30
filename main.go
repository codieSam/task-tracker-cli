package main

import (
	"fmt"
	"os"
)

// type task struct {
// 	ID          int    `json:"id"`
// 	Description string `json:"description"`
// 	Status      string `json:"status"`
// 	CreatedAt   string `json:"createdAt"`
// 	UpdatedAt   string `json:"updatedAt"`
// }

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a command")
		return
	}

	if len(os.Args) >= 3 {
		arg := os.Args[2]

		if os.Args[1] == "add" {
			fmt.Println("Adding task: ", arg)
		}

	}

}
