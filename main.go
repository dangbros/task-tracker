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
		fmt.Println("Error while converting task id into num: ", err)
		os.Exit(1)
	}
	return numID
}

func main() {
	fmt.Println("Welcome! This is Task Manager CLI")
	if len(os.Args) < 2 {
		fmt.Println("Usuage: go run main.go <command> [arguments]")
		return
	}

	command := os.Args[1]
	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Usage: go run main.go add <task>")
			return
		}
		task := strings.Join(os.Args[2:], " ")
		fmt.Println("Adding task:", task)
		addTask(task)

	case "list":
		fmt.Println("Listing all tasks...")
		option := ""
		if len(os.Args) > 2 {
			option = os.Args[2]
		}

		listTasks(option)

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Usage: go run main.go delete <task_id>")
			return
		}
		taskID := getTaskID()
		fmt.Println("Deleting task with ID:", taskID)
		deleteTask(taskID)
	case "update":
		if len(os.Args) < 3 {
			fmt.Println("Usage: go run main.go update <task_id> <status-changed>")
			return
		}
		newStatus := os.Args[3]
		TaskID := getTaskID()
		updateTaskStatus(TaskID, newStatus)

	default:
		fmt.Println("Unknown command:", command)
	}
}
