package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	io "github.com/ismail-alokin/go-todo/utils"
)

func viewTasks() {
	if isTodoListEmpty() {
		return
	}
	fmt.Println("\nView tasks\n-----------")
	listTasks()
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

	newTodo := Todo{
		Title:       title,
		Description: description,
		Completed:   completed,
	}
	fmt.Println("Added new task")
	TodoList = append(TodoList, newTodo)
}

func deleteTasks() {
	fmt.Println("hello")
	if isTodoListEmpty() {
		return
	}

	fmt.Println("\nDelete tasks\n-----------")
	fmt.Println("These are the saved tasks. Enter the task number to delete.")
	listTasks()

	arrayId, _ := io.ReceiveArrayElementId(TodoList)

	TodoList = removeElement(TodoList, arrayId)
	fmt.Printf("Task %v deleted successfully", arrayId+1)
}

func completeTask() {
	if isTodoListEmpty() {
		return
	}

	fmt.Println("\nMark Task as completed\n-----------")
	fmt.Println("These are the saved tasks. Enter the task number to complete.")
	listTasks()

	arrayId, _ := io.ReceiveArrayElementId(TodoList)
	markTodoAsCompleted(arrayId)
}

func listTasks() {
	for i, _todo := range TodoList {
		todo, _ := json.MarshalIndent(_todo, "", "    ")
		fmt.Printf("Task %v: %s", i+1, todo)
		fmt.Println()
	}
	fmt.Println()
}

func isTodoListEmpty() bool {
	if len(TodoList) == 0 {
		fmt.Println("No task found")
		return true
	}
	return false
}

func removeElement(slice []Todo, s int) []Todo {
	return append(slice[:s], slice[s+1:]...)
}

func markTodoAsCompleted(arrayId int) {
	if arrayId > len(TodoList) {
		fmt.Println("Invalid input")
		return
	}
	if TodoList[arrayId].Completed {
		fmt.Println("Task is already completed")
		return
	}

	TodoList[arrayId].Completed = true
	fmt.Printf("Task %v completed successfully", arrayId+1)
}
