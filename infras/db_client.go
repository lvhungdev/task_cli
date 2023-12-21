package infras

import (
	"context"
	"task_cli/infras/ent"
)

import _ "github.com/mattn/go-sqlite3"

func NewClient() (*ent.Client, error) {
	client, err := ent.Open("sqlite3", "file:task_db.db?cache=shared&_fk=1")
	if err != nil {
		return nil, err
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		return nil, err
	}

	return client, nil
}
