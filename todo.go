package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/ryanuber/columnize"
)

type Todo struct {
	storagePath string
	items       []TodoItem
}

func newTodo(storagePath string) *Todo {
	todo := &Todo{storagePath: storagePath}

	if _, err := os.Stat(todo.getStorageFilePath()); os.IsNotExist(err) {
		_, err := os.Create(todo.getStorageFilePath())

		if err != nil {
			log.Fatal(err)
		}
	}

	todo.parse()

	return todo
}

func (todo *Todo) getStorageFilePath() string {
	return path.Join(todo.storagePath, "todo.json")
}

func (todo *Todo) parse() {
	todo.items = []TodoItem{}

	data, err := ioutil.ReadFile(todo.getStorageFilePath())

	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(data, &todo.items)
}

func (todo *Todo) add(title string) {
	todoItem := TodoItem{Done: false, Title: title}

	todo.items = append(todo.items, todoItem)

	todo.writeToFile()
}

func (todo *Todo) done(id int) {
	if id >= 1 && id <= len(todo.items) {
		todo.items[id-1].Done = true

		todo.writeToFile()
	}
}

func (todo *Todo) undone(id int) {
	if id >= 1 && id <= len(todo.items) {
		todo.items[id-1].Done = false

		todo.writeToFile()
	}
}

func (todo *Todo) delete(id int) {
	if id >= 1 && id <= len(todo.items) {
		todoItems := []TodoItem{}

		for i, todoItem := range todo.items {
			if i != id-1 {
				todoItems = append(todoItems, todoItem)
			}
		}

		todo.items = todoItems

		todo.writeToFile()
	}
}

func (todo *Todo) clear() {
	todo.items = []TodoItem{}

	todo.writeToFile()
}

func (todo *Todo) list() {
	content := []string{
		"ID|Status|Title",
	}

	for i, todoItem := range todo.items {
		if todoItem.Done {
			content = append(content, fmt.Sprintf("%d|%s|%s", (i+1), "Done", todoItem.Title))
		} else {
			content = append(content, fmt.Sprintf("%d|%s|%s", (i+1), "Undone", todoItem.Title))
		}
	}

	fmt.Printf("%s\n", columnize.SimpleFormat(content))
}

func (todo *Todo) writeToFile() {
	content, err := json.Marshal(todo.items)

	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(todo.getStorageFilePath(), content, 0644)

	if err != nil {
		log.Fatal(err)
	}
}
