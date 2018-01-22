package main

type TodoItem struct {
	Done  bool   `json:"done"`
	Title string `json:"title"`
}
