package main

import (
	"context"
	"log"
)

func main() {
	ctx := context.Background()

	taskUseCases, err := InitializeTaskUseCase(ctx)

	if err != nil {
		log.Fatalln(err)
	}

	tasks, err := taskUseCases.GetAllTasks()
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(tasks)
}
