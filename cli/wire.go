//go:build wireinject
// +build wireinject

package main

import (
	"context"

	"github.com/google/wire"

	"task_cli/core"
	"task_cli/infras"
)

func InitializeTaskUseCase(ctx context.Context) (core.TaskUseCases, error) {
	wire.Build(
		core.NewTaskUseCases,
		infras.NewTaskRepository,
		infras.NewClient,
		wire.Bind(new(core.TaskRepository), new(infras.TaskRepository)),
	)

	return core.TaskUseCases{}, nil
}
