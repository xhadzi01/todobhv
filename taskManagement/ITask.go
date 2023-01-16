package taskManagement

type TaskID int8

type ITaskObserver interface {
	GetID() TaskID
	GetName() string
	GetDescription() string
	GetState() TaskStatus
	GetPriority() TaskPriority
	GetDateCreated() string
	GetDateCompleted() *string
}

type ITaskSetter interface {
	*ITaskObserver
	SetName(string) error
	SetDescription(string) error
	SetState(TaskStatus)
	SetPriority(TaskPriority)
}
