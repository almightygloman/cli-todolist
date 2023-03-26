package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

type task struct {
	ID       int       `json:"ID"`
	DESC     string    `json:"DESC"`
	DONE     bool      `json:"DONE"`
	ASSIGNED time.Time `json:"ASSIGNED"`
	DUE      time.Time `json:"DUE"`
}

func main() {
	tasks := loadTasks()
	for {
		ui()
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')

		switch strings.TrimSuffix(input, "\n") {
		case "list":
			listall(tasks)

		case "add":
			tasks = add(tasks)
			saveTasksToFile(tasks)

		case "comp":
			var id int
			fmt.Print("Enter the ID of the task to remove: ")
			fmt.Scanln(&id)
			if err := completeTask(id, &tasks); err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Task removed successfully")
				saveTasksToFile(tasks)
			}
		case "exit":
			fmt.Println("goodbye!")
			os.Exit(0)
		}
	}
}

func ui() {
	fmt.Println("Please choose an option:")
	fmt.Println("Enter 'list' to display all tasks")
	fmt.Println("Enter 'add' to add a task")
	fmt.Println("Enter 'comp' to complete a task")
	fmt.Println("Enter 'exit' to quit")
}

//list method
func listall(tasks []task) {
	//for each task in tasks print task
	if len(tasks) == 0 {
		fmt.Println("no tasks")
	} else {
		for _, task := range tasks {
			fmt.Printf("%d. %s (assigned: %s, due date: %s, done: %t)\n",
				task.ID, task.DESC, task.ASSIGNED.Format("2006-01-02 15:04:05"), task.DUE.Format("2006-01-02 15:04"), task.DONE)
		}
	}
}

//add method
func add(tasks []task) []task {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter task description: ")
	description, _ := reader.ReadString('\n')
	description = strings.TrimSuffix(description, "\n")

	fmt.Println("Enter the Due Date (YYYY-MM-DD HH:MM): ")
	dueDateString, _ := reader.ReadString('\n')
	dueDateString = strings.TrimSuffix(dueDateString, "\n")
	dueDate, err := time.Parse("2006-01-02 15:04", dueDateString)
	if err != nil {
		fmt.Println("Invalid date format")
		return tasks
	}
	ass := time.Now()

	newTask := task{
		ID:       len(tasks) + 1,
		DESC:     description,
		DONE:     false,
		ASSIGNED: ass,
		DUE:      dueDate,
	}

	tasks = append(tasks, newTask)

	saveTasksToFile(tasks)

	fmt.Print("Task added")

	return tasks

}

//complete method
func completeTask(id int, tasks *[]task) error {
	taskIndex := -1
	for i, task := range *tasks {
		if task.ID == id {
			taskIndex = i
			break
		}
	}
	if taskIndex == -1 {
		return fmt.Errorf("task with ID %d not found", id)
	}

	*tasks = append((*tasks)[:taskIndex], (*tasks)[taskIndex+1:]...)

	for i := range *tasks {
		(*tasks)[i].ID = i + 1
	}

	return nil
}

//load json data
func loadTasks() []task {
	tasks := []task{}
	if _, err := os.Stat("tasks.json"); os.IsNotExist(err) {
		return tasks
	}

	fileBytes, err := ioutil.ReadFile("tasks.json")
	if err != nil {
		fmt.Println("Error loading tasks from file:", err)
		return tasks
	}

	if err := json.Unmarshal(fileBytes, &tasks); err != nil {
		fmt.Println("Error unmarshaling tasks from file:", err)
	}

	return tasks

}

//save json data
func saveTasksToFile(tasks []task) {
	file, err := os.Create("tasks.json")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()
	taskBytes, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling tasks to JSON:", err)
		return
	}

	if _, err := file.Write(taskBytes); err != nil {
		fmt.Println("Error writing tasks to file:", err)
	}
}
