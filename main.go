package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getTaskID() int {
	taskID := os.Args[2]
	numID, err := strconv.Atoi(taskID)
	if err != nil {
		fmt.Println(red+"❌Error: Task ID must be a number."+reset, err)
		os.Exit(1)
	}
	return numID
}

func main() {
	fmt.Println(green + "Welcome! This is Task Manager CLI" + reset)
	if len(os.Args) < 2 {
		fmt.Println(yellow + "Usuage: go run main.go <command> [arguments]" + reset)
		return
	}

	command := os.Args[1]
	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println(yellow + "Usage: go run main.go add <task>" + reset)
			return
		}
		task := strings.Join(os.Args[2:], " ")
		fmt.Println(yellow+"Adding task:"+reset, task)
		addTask(task)

	case "list":
		fmt.Println(yellow + "Listing all tasks..." + reset)
		option := ""
		if len(os.Args) > 2 {
			option = os.Args[2]
		}

		listTasks(option)

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println(red + "❌Error: Missing Task ID." + reset)
			fmt.Println(yellow + "Usage: go run main.go delete <task_id>" + reset)
			return
		}
		taskID := getTaskID()
		fmt.Println(yellow+"Deleting task with ID:"+reset, taskID)
		deleteTask(taskID)
	case "update":
		if len(os.Args) < 3 {
			fmt.Println(yellow + "Usage: go run main.go update <task_id> <status-changed>" + reset)
			return
		}
		newStatus := os.Args[3]
		TaskID := getTaskID()
		updateTaskStatus(TaskID, newStatus)

	case "edit":
		if len(os.Args) < 4 {
			fmt.Println(yellow + "Usage: go run main.go edit <task_id> <new title>" + reset)
			return
		}

		taskID := getTaskID()
		newtitle := strings.Join(os.Args[3:], " ")
		editTaskTitle(taskID, newtitle)

	default:
		fmt.Println(red+"❌Error: Unknown command:"+reset, command)
	}
}
