package core

import "time"

type Task struct {
	Id       int
	Title    string
	IsDone   bool
	DoneDate *time.Time
}
