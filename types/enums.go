package types

type TaskStatus string

const (
	TaskStatusTodo  TaskStatus = "todo"
	TaskStatusDoing TaskStatus = "doing"
	TaskStatusDone  TaskStatus = "done"
)

type ProviderEnum string

const (
	ProviderEnumGoogle ProviderEnum = "google"
)
