package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"
	"task_cli/core"
)

func main() {
	taskUseCases, err := InitializeTaskUseCase(context.Background())

	if err != nil {
		log.Fatalln(err)
	}

	addFlag := flag.Bool("add", false, "add new task")
	completeFlag := flag.Bool("complete", false, "complete task")

	flag.Parse()

	switch {
	case *addFlag:
		addTask(taskUseCases)
	case *completeFlag:
		completeTask(taskUseCases)
	default:
		printTasks(taskUseCases)
	}
}

func addTask(uc core.TaskUseCases) {
	title := strings.Join(flag.Args(), " ")
	task, err := uc.AddTask(title)

	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Printf(`Task added { Id: %d, Title: "%v" }`, task.Id, task.Title)
	}
}

func completeTask(uc core.TaskUseCases) {
	if len(flag.Args()) == 0 {
		fmt.Println("error: id not found")
		return
	}

	id, err := strconv.Atoi(flag.Args()[0])
	if err != nil {
		fmt.Println("error: id is invalid")
		return
	}

	task, err := uc.CompleteTask(id)

	if err != nil {
		fmt.Println("error: ", err)
	} else {
		fmt.Printf(
			`Task completed { Id: %d, Title: "%v", IsDone: true, DoneDate: %v }`,
			task.Id,
			task.Title,
			(*task.DoneDate).Format("2006-01-02 15:04:05"),
		)
	}
}

func printTasks(uc core.TaskUseCases) {
	tasks, err := uc.GetAllTasks()

	if err != nil {
		fmt.Println("error: ", err)
		return
	}

	for _, task := range tasks {
		doneDate := ""
		if task.DoneDate != nil {
			doneDate = (*task.DoneDate).Format("2006-01-02 15:04:05")
		}

		fmt.Printf(
			"{ Id: %d, Title: \"%v\", IsDone: %v, DoneDate: %v }\n",
			task.Id,
			task.Title,
			task.IsDone,
			doneDate,
		)
	}
}
