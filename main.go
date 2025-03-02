package main

//declaring packages
import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// to get taskID from arguements and convert it to int from string
func getTaskID() int {
	taskID := os.Args[2]
	numID, err := strconv.Atoi(taskID) // function to convert it into int
	if err != nil {
		fmt.Println(red+"❌Error: Task ID must be a number."+reset, err) //Error message to display
		os.Exit(1)
	}
	return numID //TaskID in int is returned
}

func main() {
	fmt.Println(green + "Welcome! This is Task Manager CLI" + reset) //This is opening line which prints everytime the program is called
	if len(os.Args) < 2 {
		fmt.Println(yellow + "Usuage: task-cli <command> [arguments]" + reset)
		return
	}

	command := os.Args[1] //commands like "add", "update", "list", "delete" are used here
	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println(yellow + "Usage: task-cli add <task>" + reset)
			return
		}
		task := strings.Join(os.Args[2:], " ") //this function helps to join the title string into a single string with spaces
		fmt.Println(yellow+"Adding task:"+reset, task)
		addTask(task) //addtask function is called from task.go

	case "list":
		fmt.Println(yellow + "Listing all tasks..." + reset)
		option := "" //it has a filter listing option according to the status of the task.
		if len(os.Args) > 2 {
			option = os.Args[2]
		}

		listTasks(option) //called from task.go

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println(red + "❌Error: Missing Task ID." + reset)
			fmt.Println(yellow + "Usage: go run main.go delete <task_id>" + reset)
			return
		}
		taskID := getTaskID() //taskId is retrieved using that function
		fmt.Println(yellow+"Deleting task with ID:"+reset, taskID)
		deleteTask(taskID) //called from task.go
	case "update":
		if len(os.Args) < 3 {
			fmt.Println(yellow + "Usage: go run main.go update <task_id> <status-changed>" + reset)
			return
		}
		newStatus := os.Args[3] //updated status is called
		TaskID := getTaskID()
		updateTaskStatus(TaskID, newStatus) //both taskID and the new status is given as arguements

	case "edit":
		//here task_id and new_title is taken and passed through the function
		if len(os.Args) < 4 {
			fmt.Println(yellow + "Usage: go run main.go edit <task_id> <new title>" + reset)
			return
		}
		taskID := getTaskID()
		newtitle := strings.Join(os.Args[3:], " ")
		editTaskTitle(taskID, newtitle)

	case "clear":
		clearAllTask() //this clears all tasks
	case "help":
		ShowHelp() //this prints the full manual of how to run this program

	default:
		fmt.Println(red+"❌Error: Unknown command:"+reset, command)
	}
}
