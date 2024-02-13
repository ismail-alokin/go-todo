package main

type Todo struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

var TodoList = []Todo{{
	ID:          1,
	Title:       "1st Task",
	Description: "1st Description",
	Completed:   false,
}}
