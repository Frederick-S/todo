package main

import (
	"encoding/json"
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
	return path.Join(todo.storagePath, "todo.json")
}

func (todo *Todo) parse() {
	todo.items = []TodoItem{}

	//data, err := ioutil.ReadFile(todo.getStorageFilePath())
}

func (todo *Todo) add(title string) {
	todoItem := TodoItem{Done: false, Title: title}

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

	_, err = todo.file.WriteString(string(content))

	if err != nil {
		log.Fatal(err)
	}
}
