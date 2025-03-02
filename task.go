package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// ANSI escape codes for terminal color formatting
const (
	reset  = "\033[0m"  // Reset all formatting
	red    = "\033[31m" // Error messages and warnings
	green  = "\033[32m" // Success messages
	yellow = "\033[33m" // Informational messages
)

// File configuration
const taskFile = "tasks.json" // Persistent storage file for tasks

// Task structure represents a single task item
type Task struct {
	ID     int    `json:"id"`     // Unique identifier for the task
	Title  string `json:"title"`  // Description of the task
	Status string `json:"status"` // Current state: "pending", "in-progress", or "done"
}

// loadTasks reads and parses tasks from the JSON file
func loadTasks() []Task {
	// Read file contents
	file, err := os.ReadFile(taskFile)
	if err != nil {
		// Handle missing file case
		if os.IsNotExist(err) {
			fmt.Println(yellow + "File does not exist" + reset)
			return []Task{}
		}
		// Handle other read errors
		fmt.Println(red+"❌Error: reading file unsuccessful: "+reset, err)
		return []Task{}
	}

	// Handle empty file case
	if len(file) == 0 {
		return []Task{}
	}

	// Unmarshal JSON data into Task slice
	var tasks []Task
	err = json.Unmarshal(file, &tasks)
	if err != nil {
		// Handle corrupted JSON data
		fmt.Println(red+"❌Error: Corrupted task.json. Creating a new file."+reset, err)
		return []Task{}
	}

	return tasks
}

// saveTask writes tasks to the JSON file with proper formatting
func saveTask(tasks []Task) {
	// Marshal tasks with indentation for readability
	newData, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		fmt.Println(red+"❌Error: encoding JSON unsuccessful: "+reset, err)
		return
	}

	// Write to file with standard permissions
	err = os.WriteFile(taskFile, newData, 0644)
	if err != nil {
		fmt.Println(red+"❌Error: writing to file unsuccessful: "+reset, err)
	}
}

// addTask creates a new task and appends it to the list
func addTask(title string) {
	// Validate input
	if strings.TrimSpace(title) == "" {
		fmt.Println(red + "❌Error: Title of the task cannot be empty!" + reset)
		return
	}

	tasks := loadTasks()
	var newID int

	// Generate new ID based on existing tasks
	if len(tasks) == 0 {
		newID = 1
	} else {
		newID = tasks[len(tasks)-1].ID + 1
	}

	// Create new task object
	newtask := Task{
		ID:     newID,
		Title:  title,
		Status: "pending", // Default status
	}

	// Update and save task list
	tasks = append(tasks, newtask)
	saveTask(tasks)

	fmt.Printf(green+"'✅%v' is added in the list. \n"+reset, title)
}

// listTasks displays tasks with color-coding based on status
func listTasks(choice string) {
	tasks := loadTasks()

	// Handle empty task list
	if len(tasks) == 0 {
		fmt.Println(red + "List of Tasks is empty" + reset)
		return
	}

	found := false
	// Filter tasks if status parameter is provided
	if choice != "" {
		// Validate status filter
		switch choice {
		case "pending", "in-progress", "done":
			// Iterate through tasks and print matching ones
			for _, task := range tasks {
				if task.Status == choice {
					color := reset
					// Set color based on status
					switch task.Status {
					case "pending":
						color = yellow
					case "in-progress":
						color = red
					case "done":
						color = green
					}
					fmt.Printf("%s[%v] %v - %v\n"+reset, color, task.ID, task.Title, task.Status)
					found = true
				}
			}

			// Handle no matches case
			if !found {
				fmt.Println(red + "No task found for this status" + reset)
			}
		default:
			fmt.Println(red + "Invalid choice!" + reset)
		}
	} else {
		// Print all tasks when no filter is specified
		for _, task := range tasks {
			color := reset
			switch task.Status {
			case "pending":
				color = yellow
			case "in-progress":
				color = red
			case "done":
				color = green
			}
			fmt.Printf("%s[%v] %v - %v\n"+reset, color, task.ID, task.Title, task.Status)
		}
	}
}

// updateTaskStatus modifies the status of an existing task
func updateTaskStatus(givenID int, newStatus string) {
	// Validate new status
	validStatuses := map[string]bool{
		"pending":     true,
		"in-progress": true,
		"done":        true,
	}

	if !validStatuses[newStatus] {
		fmt.Println(red + "❌Error: Invalid Status. Use 'pending', 'in-progress' or 'done'." + reset)
		return
	}

	tasks := loadTasks()
	found := false

	// Handle empty task list
	if len(tasks) == 0 {
		fmt.Println(yellow + "List of Tasks is empty" + reset)
		return
	}

	// Find and update task
	for i, task := range tasks {
		if task.ID == givenID {
			tasks[i].Status = newStatus
			saveTask(tasks)
			fmt.Println(green + "✅Task updated successfully" + reset)
			found = true
			break
		}
	}

	// Handle task not found
	if !found {
		fmt.Println(red + "❌Error: Task is not found." + reset)
	}
}

// deleteTask removes a task from the list and renumbers IDs
func deleteTask(givenID int) {
	tasks := loadTasks()
	var updatedTask []Task
	found := false

	// Filter out the task to delete
	for _, task := range tasks {
		if task.ID != givenID {
			updatedTask = append(updatedTask, task)
		} else {
			found = true
		}
	}

	// Renumber remaining tasks
	if found {
		for i := range updatedTask {
			updatedTask[i].ID = i + 1
		}
		saveTask(updatedTask)
		fmt.Println(green + "✅Task deleted successfully" + reset)
	} else {
		fmt.Println(red + "❌Error: Task not found" + reset)
	}
}

// editTaskTitle modifies the title of an existing task
func editTaskTitle(TaskId int, newtitle string) {
	Tasks := loadTasks()
	// Find and update task title
	for i, task := range Tasks {
		if task.ID == TaskId {
			Tasks[i].Title = newtitle
			saveTask(Tasks)
			fmt.Println(green + "✅Task title edited successfully" + reset)
			return
		}
	}
	fmt.Println(red + "❌Error: TaskId not found" + reset)
}

// clearAllTask resets the task list completely
func clearAllTask() {
	var Tasks []Task
	saveTask(Tasks)
	fmt.Println(green + "✅All task cleared!" + reset)
}

// ShowHelp displays usage instructions and examples
func ShowHelp() {
	fmt.Println(yellow + "📖 Task Manager CLI - Help Menu" + reset)
	fmt.Println(green + "\nUsage:" + reset)
	fmt.Println("  task-cli add <task_title>       → Add a new task")
	fmt.Println("  task-cli list                   → List all tasks")
	fmt.Println("  task-cli list <status>          → List tasks by status (pending, done, in-progress)")
	fmt.Println("  task-cli update <id> <status>   → Update task status (pending, done, in-progress)")
	fmt.Println("  task-cli edit <id> <new_title>  → Edit a task title")
	fmt.Println("  task-cli delete <id>            → Delete a task")
	fmt.Println("  task-cli clear                  → Clear all tasks")
	fmt.Println("  task-cli help                   → Show this help menu")
	fmt.Println(green + "\nExamples:" + reset)
	fmt.Println("  task-cli add Buy groceries")
	fmt.Println("  task-cli update 2 done")
	fmt.Println("  task-cli list pending")
	fmt.Println("  task-cli edit 3 Read a book")
}
