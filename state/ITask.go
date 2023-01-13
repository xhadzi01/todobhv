package state

import "time"

type ITaskObserver interface {
	GetID() TaskID
	GetName() string
	GetDescription() string
	GetState() TaskStatus
	GetPriority() Priority
	GetDateCreated() time.Time
	GetDateCompleted() *time.Time
}

type ITaskSetter interface {
	*ITaskObserver
	SetName(string) error
	SetDescription(string) error
	SetState(TaskStatus)
	SetPriority(Priority)
}
