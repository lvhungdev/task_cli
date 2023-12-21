package core

type TaskRepository interface {
	GetTaskById(id int) (Task, error)
	GetAllTasks() ([]Task, error)
	AddTask(task Task) (Task, error)
	UpdateTask(task Task) (Task, error)
	RemoveTask(id int) error
}
