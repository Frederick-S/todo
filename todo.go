package main

import (
	"log"
	"os"
	"path"
)

type Todo struct {
	storagePath string
	file        *os.File
}

func newTodo(storagePath string) *Todo {
	todo := &Todo{storagePath: storagePath}

	if _, err := os.Stat(todo.getStorageFilePath()); os.IsNotExist(err) {
		file, err := os.Create(todo.getStorageFilePath())

		if err != nil {
			log.Fatal(err)
		}

		todo.file = file
	} else {
		file, err := os.Open(todo.getStorageFilePath())

		if err != nil {
			log.Fatal(err)
		}

		todo.file = file
	}

	return todo
}

func (todo *Todo) getStorageFilePath() string {
	return path.Join(todo.storagePath, "todo.db")
}

func (todo *Todo) add(task string) {

}

func (todo *Todo) done(id int32) {

}

func (todo *Todo) undone(id int32) {

}

func (todo *Todo) delete(id int32) {

}

func (todo *Todo) deleteAll() {

}

func (todo *Todo) list() {

}
