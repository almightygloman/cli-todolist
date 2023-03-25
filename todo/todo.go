package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type task struct {
	desc     string
	done     bool
	assigned time.Time
	due      time.Time
}

//array to collet tasks

func main() {
	tasks := []task{}
	for {
		ui()
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')

		switch strings.TrimSuffix(input, "\n") {
		case "list":
			listall(tasks)

		case "add":
			//add()
		case "check":
			//complete()
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
	fmt.Println("Enter 'check' to complete a task")
	fmt.Println("Enter 'exit' to quit")
}

//list method
func listall(tasks []task) {
	//for each task in tasks print task
	if len(tasks) == 0 {
		fmt.Println("no tasks")
	} else {
		for i, task := range tasks {
			fmt.Println("Tasks:")
			fmt.Printf("%d. %s (done: %t)\n", i+1, task.desc, task.done)
		}
	}
}

//add method
func add(tasks []task) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter the task's description: ")

}

//complete method
func complete(tasks []task) {

}
