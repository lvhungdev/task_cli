package infras

import (
	"context"

	"task_cli/core"
	"task_cli/infras/ent"
)

type TaskRepository struct {
	client *ent.Client
	ctx    context.Context
}

func NewTaskRepository(client *ent.Client, ctx context.Context) TaskRepository {
	return TaskRepository{
		client: client,
		ctx:    ctx,
	}
}

func (repo TaskRepository) GetTaskById(id int) (core.Task, error) {
	task, err := repo.client.Task.
		Get(repo.ctx, id)

	if err != nil {
		return core.Task{}, err
	}

	return mapEntToCoreTask(*task), nil
}

func (repo TaskRepository) GetAllTasks() ([]core.Task, error) {
	tasks, err := repo.client.Task.
		Query().
		All(repo.ctx)

	if err != nil {
		return nil, err
	}

	var coreTasks []core.Task
	for _, task := range tasks {
		coreTasks = append(coreTasks, mapEntToCoreTask(*task))
	}

	return coreTasks, nil
}

func (repo TaskRepository) AddTask(task core.Task) (core.Task, error) {
	createdTask, err := repo.client.Task.
		Create().
		SetTitle(task.Title).
		SetIsDone(task.IsDone).
		SetNillableDoneDate(task.DoneDate).
		Save(repo.ctx)

	if err != nil {
		return core.Task{}, err
	}

	return mapEntToCoreTask(*createdTask), nil
}

func (repo TaskRepository) UpdateTask(task core.Task) (core.Task, error) {
	updatedTask, err := repo.client.Task.
		UpdateOneID(task.ID).
		SetTitle(task.Title).
		SetIsDone(task.IsDone).
		SetNillableDoneDate(task.DoneDate).
		Save(repo.ctx)

	if err != nil {
		return core.Task{}, err
	}

	return mapEntToCoreTask(*updatedTask), nil
}

func (repo TaskRepository) RemoveTask(id int) error {
	return repo.client.Task.
		DeleteOneID(id).
		Exec(repo.ctx)
}

func mapEntToCoreTask(task ent.Task) core.Task {
	return core.Task{
		ID:       task.ID,
		Title:    task.Title,
		IsDone:   task.IsDone,
		DoneDate: task.DoneDate,
	}
}
