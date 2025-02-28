package main

import (
	"encoding/json"
	"fmt"
	"os"
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
			fmt.Println("File does not exist")
			return []Task{}
		}
		fmt.Println("Error reading file: ", err)
		os.Exit(1)
	}

	if len(file) == 0 {
		return []Task{}
	}

	var tasks []Task
	err = json.Unmarshal(file, &tasks)
	if err != nil {
		fmt.Println("Error reading file: ", err)
		os.Exit(1)
	}

	return tasks
}

func saveTask(tasks []Task) {
	newData, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		fmt.Println("Error encoding JSON: ", err)
		return
	}
	err = os.WriteFile(taskFile, newData, 0644)
	if err != nil {
		fmt.Print("Error wrting to file: ", err)
	}
}

func addTask(title string) {
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

	fmt.Printf("'%v' is added in the list. \n", title)

}

func listTasks(choice string) {
	tasks := loadTasks()

	if len(tasks) == 0 {
		fmt.Println("List of Tasks is empty")
		return
	}

	found := false
	if choice != "" {
		switch choice {
		case "pending", "in-progress", "done":
			for _, task := range tasks {
				if task.Status == choice {
					fmt.Printf("[%v] %v - %v\n", task.ID, task.Title, task.Status)
					found = true
				}
			}
			if !found {
				fmt.Println("No task found for this status")
			}
		default:
			fmt.Println("Invalid choice!")
		}
	} else {
		for _, task := range tasks {
			fmt.Printf("[%v] %v - %v\n", task.ID, task.Title, task.Status)
		}
	}
}

func updateTaskStatus(givenID int, newStatus string) {
	tasks := loadTasks()
	if len(tasks) == 0 {
		fmt.Println("List of Tasks is empty")
		return
	}
	for i, task := range tasks {
		if task.ID == givenID {
			tasks[i].Status = newStatus
			saveTask(tasks)
			fmt.Println("Task updated successfully")
			return
		}
	}
	fmt.Println("Task is not found")
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
		fmt.Println("Task deleted successfully")
	} else {
		fmt.Println("Task not found")
	}
}
