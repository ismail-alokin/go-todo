package main

type Todo struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

var TodoList = []Todo{{
	Title:       "1st Task",
	Description: "1st Description",
	Completed:   false,
}}
