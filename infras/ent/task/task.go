// Code generated by ent, DO NOT EDIT.

package task

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the task type in the database.
	Label = "task"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldIsDone holds the string denoting the is_done field in the database.
	FieldIsDone = "is_done"
	// FieldDoneDate holds the string denoting the done_date field in the database.
	FieldDoneDate = "done_date"
	// Table holds the table name of the task in the database.
	Table = "tasks"
)

// Columns holds all SQL columns for task fields.
var Columns = []string{
	FieldID,
	FieldTitle,
	FieldIsDone,
	FieldDoneDate,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the Task queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByTitle orders the results by the title field.
func ByTitle(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTitle, opts...).ToFunc()
}

// ByIsDone orders the results by the is_done field.
func ByIsDone(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsDone, opts...).ToFunc()
}

// ByDoneDate orders the results by the done_date field.
func ByDoneDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDoneDate, opts...).ToFunc()
}
