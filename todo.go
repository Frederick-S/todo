package main

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/ryanuber/columnize"
)

type Todo struct {
	storagePath string
	file        *os.File
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

	file, err := os.OpenFile(todo.getStorageFilePath(), os.O_WRONLY, os.ModeAppend)

	if err != nil {
		log.Fatal(err)
	}

	todo.file = file

	todo.parse()

	return todo
}

func (todo *Todo) getStorageFilePath() string {
	return path.Join(todo.storagePath, "todo.db")
}

func (todo *Todo) parse() {
	todo.items = []TodoItem{}
}

func (todo *Todo) add(title string) {
	todoItem := TodoItem{done: false, title: title}

	todo.items = append(todo.items, todoItem)

	todo.writeToFile()
}

func (todo *Todo) done(id int32) {

}

func (todo *Todo) undone(id int32) {

}

func (todo *Todo) delete(id int32) {

}

func (todo *Todo) clear() {

}

func (todo *Todo) list() {

}

func (todo *Todo) writeToFile() {
	content := []string{
		"ID|Status|Title",
	}

	for i, todoItem := range todo.items {
		if todoItem.done {
			content = append(content, fmt.Sprintf("%d|%s|%s", (i+1), "Done", todoItem.title))
		} else {
			content = append(content, fmt.Sprintf("%d|%s|%s", (i+1), "Undone", todoItem.title))
		}
	}

	_, err := todo.file.WriteString(fmt.Sprintf("%s\n", columnize.SimpleFormat(content)))

	if err != nil {
		log.Fatal(err)
	}
}
