package core

import "time"

type Task struct {
	ID       int
	Title    string
	IsDone   bool
	DoneDate *time.Time
}
