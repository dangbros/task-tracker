package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

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
		listTasks()

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Usage: go run main.go delete <task_id>")
			return
		}
		taskID := os.Args[2]
		fmt.Println("Deleting task with ID:", taskID)
	case "update":
		if len(os.Args) < 3 {
			fmt.Println("Usage: go run main.go update <task_id> <status-changed>")
			return
		}
		taskID := os.Args[2]
		newStatus := os.Args[3]
		numID, err := strconv.Atoi(taskID)
		if err != nil {
			fmt.Println("Error while converting task id into num: ", err)
			os.Exit(1)
		}
		updateTaskStatus(numID, newStatus)

	default:
		fmt.Println("Unknown command:", command)
	}
}
