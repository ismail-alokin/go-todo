package main

import (
	"fmt"

	io "github.com/ismail-alokin/go-todo/utils"
)

func main() {
	title := "Welcome to Todo list"
	io.PrintTitle(title)

	for {
		message := "What do you want to do?\n\t1. View Tasks\n\t2. Add Tasks\n\t3. Delete Tasks\n\t4. Mark Task as Completed\n\t5. Exit\n"
		input := io.GetAString(message)

		if input == "5" {
			break
		}
		if !io.ValidateInput(input) {
			continue
		}

		switch input {

		case "1":
			viewTasks()
		case "2":
			addTasks()
		case "3":
			deleteTasks()

		}

		fmt.Print("\n\n")
	}
}
