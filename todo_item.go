package main

type TodoItem struct {
	ID    int    `json:"id"`
	Done  bool   `json:"done"`
	Title string `json:"title"`
}
