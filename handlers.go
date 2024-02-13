package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"

	io "github.com/ismail-alokin/go-todo/utils"
)

func viewTasks() {
	fmt.Println("\nView tasks\n-----------")
	if len(TodoList) == 0 {
		fmt.Println("No task found")
		return
	}
	for i, _todo := range TodoList {
		todo, _ := json.MarshalIndent(_todo, "", "    ")
		fmt.Printf("Task %v: %s", i+1, todo)
		fmt.Println()
	}
	fmt.Println()
}

func addTasks() {
	fmt.Println("\nAdd tasks\n-----------")
	title := io.GetAString("Input TODO title")
	description := io.GetAString("Input Description")
	completedStr := io.GetAString("Input completion status (Y/N)")

	completed := false
	if strings.ToLower(completedStr) == "y" {
		completed = true
	} else if strings.ToLower(completedStr) == "n" {
		completed = false
	} else {
		log.Println("Invalid Input")
		return
	}

	id := TodoList[len(TodoList)-1].ID + 1

	newTodo := Todo{
		ID:          id,
		Title:       title,
		Description: description,
		Completed:   completed,
	}
	fmt.Println("Added new task")
	TodoList = append(TodoList, newTodo)
}

func deleteTasks() {
	fmt.Println("\nDelete tasks\n-----------")
	fmt.Println("These are the saved tasks. Enter the ID of the task to delete.")

	for _, todo := range TodoList {
		fmt.Printf("\tID: %v, Title: %v, Completed: %v\n", todo.ID, todo.Title, todo.Completed)
		fmt.Println()
	}

	arrayIdStr := io.GetAString("Enter ID")
	fmt.Println(arrayIdStr)

	arrayId, err := strconv.Atoi(arrayIdStr)
	if err != nil {
		return
	}

	removeElement(TodoList, arrayId-1)

}

func removeElement(slice []Todo, s int) []Todo {
	return append(slice[:s], slice[s+1:]...)
}
