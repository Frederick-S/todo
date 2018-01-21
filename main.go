package main

import (
	"fmt"
	"log"
	"os"
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

	arguments := os.Args[1:]
	command := arguments[0]

	switch command {
	case "add":
		if len(arguments) < 2 {
			log.Fatal("Missing todo title")
		}

		title := arguments[1]

		todo.add(title)

		break
	default:
		log.Fatal(fmt.Sprintf("Unknown command %s", command))

		break
	}
}
