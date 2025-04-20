package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	add := flag.String("add", "", "Task to add")
	list := flag.Bool("list", false, "List all tasks")
	done := flag.Int("done", 0, "Mark task as done (pass task number)")

	flag.Parse()

	todoList, err := LoadTasks("tasks.json")
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		os.Exit(1)
	}

	switch {
	case *add != "":
		todoList = append(todoList, Task{Text: *add})
		err = SaveTasks("tasks.json", todoList)
		if err != nil {
			fmt.Println("Error saving task:", err)
		} else {
			fmt.Println("Task added!")
		}
	case *list:
		PrintTasks(todoList)
	case *done > 0:
		index := *done - 1
		if index < 0 || index >= len(todoList) {
			fmt.Println("Invalid task number")
			os.Exit(1)
		}
		todoList[index].Done = true
		err = SaveTasks("tasks.json", todoList)
		if err != nil {
			fmt.Println("Error updating task:", err)
		} else {
			fmt.Println("Task marked as done!")
		}
	default:
		fmt.Println("Usage:")
		flag.PrintDefaults()
	}
}
