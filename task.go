package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

const (
	reset  = "\033[0m"
	red    = "\033[31m"
	green  = "\033[32m"
	yellow = "\033[33m"
)

const taskFile = "tasks.json"

type Task struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"` //"pending", "in-progress" or "done"
}

func loadTasks() []Task {

	file, err := os.ReadFile(taskFile)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println(yellow + "File does not exist" + reset)
			return []Task{}
		}
		fmt.Println(red+"❌Error: reading file unsuccessful: "+reset, err)
		return []Task{}
	}

	if len(file) == 0 {
		return []Task{}
	}

	var tasks []Task
	err = json.Unmarshal(file, &tasks)
	if err != nil {
		fmt.Println(red+"❌Error: Corrupted task.json. Creating a new file."+reset, err)
		return []Task{}
	}

	return tasks
}

func saveTask(tasks []Task) {
	newData, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		fmt.Println(red+"❌Error: encoding JSON unsuccessful: "+reset, err)
		return
	}
	err = os.WriteFile(taskFile, newData, 0644)
	if err != nil {
		fmt.Println(red+"❌Error: wrting to file unsuccessful: "+reset, err)
	}
}

func addTask(title string) {
	if strings.TrimSpace(title) == "" {
		fmt.Println(red + "❌Error: Title of the task cannot be empty!" + reset)
		return
	}
	tasks := loadTasks()
	var newID int

	if len(tasks) == 0 {
		newID = 1
	} else {
		newID = tasks[len(tasks)-1].ID + 1
	}

	newtask := Task{
		ID:     newID,
		Title:  title,
		Status: "pending",
	}

	tasks = append(tasks, newtask)
	saveTask(tasks)

	fmt.Printf(green+"'✅%v' is added in the list. \n"+reset, title)

}

func listTasks(choice string) {
	tasks := loadTasks()

	if len(tasks) == 0 {
		fmt.Println(red + "List of Tasks is empty" + reset)
		return
	}

	found := false
	if choice != "" {
		switch choice {
		case "pending", "in-progress", "done":
			for _, task := range tasks {
				if task.Status == choice {
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
					found = true
				}
			}

			if !found {
				fmt.Println(red + "No task found for this status" + reset)
			}
		default:
			fmt.Println(red + "Invalid choice!" + reset)
		}

	} else {
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

func updateTaskStatus(givenID int, newStatus string) {
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
	if len(tasks) == 0 {
		fmt.Println(yellow + "List of Tasks is empty" + reset)
		return
	}
	for i, task := range tasks {
		if task.ID == givenID {
			tasks[i].Status = newStatus
			saveTask(tasks)
			fmt.Println(green + "✅Task updated successfully" + reset)
			found = true
			break
		}
	}

	if !found {
		fmt.Println(red + "❌Error: Task is not found." + reset)
	}
}

func deleteTask(givenID int) {
	tasks := loadTasks()
	var updatedTask []Task
	found := false
	for _, task := range tasks {
		if task.ID != givenID {
			updatedTask = append(updatedTask, task)
		} else {
			found = true
		}
	}

	for i := range updatedTask {
		updatedTask[i].ID = i + 1
	}
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

func editTaskTitle(TaskId int, newtitle string) {
	Tasks := loadTasks()
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

func clearAllTask() {
	var Tasks []Task
	saveTask(Tasks)
	fmt.Println(green + "✅All task cleared!" + reset)
}
