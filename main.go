package main

import (
	"fmt"
	"log"
	"os/user"
)

func getHomeFolder() string {
	usr, err := user.Current()

	if err != nil {
		log.Fatal(err)
	}

	return usr.HomeDir
}

func main() {
	todo := newTodo(getHomeFolder())

	fmt.Println(todo.getStorageFilePath())
}
