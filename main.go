package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Task struct {
	TaskName  string
	completed bool
}

var tasks []Task

func addTask(task string) {
	newTask := Task{TaskName: task, completed: false}
	tasks = append(tasks, newTask)

	fmt.Println("Added new task: " + newTask.TaskName)
}

func listTasks() {
	for i, task := range tasks {
		status := "Not Done"
		if task.completed {
			status = "Done"
		}

		fmt.Printf("%d: %s [%s]\n", i+1, task.TaskName, status)
	}
}

func markCompleted(index int) {
	if index >= 1 && index <= len(tasks) {
		tasks[index-1].completed = true
		fmt.Println("Marked completed: " + tasks[index-1].TaskName)
	} else {
		fmt.Println("Not marked completed: " + tasks[len(tasks)-1].TaskName)
	}
}

func editTask(index int, newTask string) {
	if index >= 1 && index <= len(tasks) {
		tasks[index-1].TaskName = newTask
		fmt.Println("Edited task: " + tasks[index-1].TaskName)
	} else {
		fmt.Println("Invalid Index")
	}
}

func deleteTask(index int) {
	if index >= 1 && index <= len(tasks) {
		tasks = append(tasks[:index-1], tasks[index:]...)
		fmt.Println("Task Deleted")
	} else {
		fmt.Println("Invalid Index")
	}
}

func main() {
	//var indexnput int
	var taskInput, newtaskInput string

	fmt.Println("Options")
	fmt.Println("1. Add new task")
	fmt.Println("2. List tasks")
	fmt.Println("3. Mark completed")
	fmt.Println("4. Edit task")
	fmt.Println("5. Delete task")
	fmt.Println("6. Exit")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter choice (1, 2, 3, 4, 5, 6): ")
		scanner.Scan()
		input := scanner.Text()

		choice, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid choice")
			continue
		}

		switch choice {
		case 1:
			fmt.Print("Enter task: ")
			scanner.Scan()
			taskInput = scanner.Text()
			addTask(taskInput)
		case 2:
			listTasks()
		case 3:
			fmt.Print("Enter index: ")
			scanner.Scan()
			indexinput, _ := strconv.Atoi(scanner.Text())
			markCompleted(indexinput)
		case 4:
			fmt.Println("Enter index: ")
			scanner.Scan()
			indexinput, _ := strconv.Atoi(scanner.Text())
			fmt.Print("Enter task: ")
			scanner.Scan()
			newtaskInput = scanner.Text()
			editTask(indexinput, newtaskInput)
		case 5:
			fmt.Println("Enter index: ")
			scanner.Scan()
			indexinput, _ := strconv.Atoi(scanner.Text())
			deleteTask(indexinput)
		case 6:
			os.Exit(0)

		default:
			fmt.Println("Invalid choice")
		}
	}

}
