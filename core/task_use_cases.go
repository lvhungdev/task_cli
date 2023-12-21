package core

import "time"

type TaskUseCases struct {
	repo TaskRepository
}

func NewTaskUseCases(repo TaskRepository) TaskUseCases {
	return TaskUseCases{repo: repo}
}

func (uc TaskUseCases) GetAllTasks() ([]Task, error) {
	return uc.repo.GetAllTasks()
}

func (uc TaskUseCases) GetAllUnfinishedTasks() ([]Task, error) {
	tasks, err := uc.repo.GetAllTasks()

	if err != nil {
		return nil, err
	}

	var unfinishedTasks []Task
	for _, t := range tasks {
		if !t.IsDone {
			unfinishedTasks = append(unfinishedTasks, t)
		}
	}

	return unfinishedTasks, nil
}

func (uc TaskUseCases) AddTask(title string) (Task, error) {
	task := Task{
		Title:    title,
		IsDone:   false,
		DoneDate: nil,
	}

	return uc.repo.AddTask(task)
}

func (uc TaskUseCases) CompleteTask(id int) (Task, error) {
	task, err := uc.repo.GetTaskById(id)
	if err != nil {
		return Task{}, err
	}

	now := time.Now()

	task.IsDone = true
	task.DoneDate = &now

	return uc.repo.UpdateTask(task)
}
