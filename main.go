package main

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"strconv"
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
	case "done":
		if len(arguments) < 2 {
			log.Fatal("Missing todo id")
		}

		id, err := strconv.Atoi(arguments[1])

		if err != nil {
			log.Fatal(err)
		}

		todo.done(id)

		break
	case "undone":
		if len(arguments) < 2 {
			log.Fatal("Missing todo id")
		}

		id, err := strconv.Atoi(arguments[1])

		if err != nil {
			log.Fatal(err)
		}

		todo.undone(id)

		break
	case "delete":
		if len(arguments) < 2 {
			log.Fatal("Missing todo id")
		}

		id, err := strconv.Atoi(arguments[1])

		if err != nil {
			log.Fatal(err)
		}

		todo.delete(id)

		break
	case "clear":
		todo.clear()

		break
	case "list":
		todo.list()

		break
	default:
		log.Fatal(fmt.Sprintf("Unknown command %s", command))

		break
	}
}
