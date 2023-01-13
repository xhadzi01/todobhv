package state

import (
	"errors"
	"time"
)

type TaskStatus uint8

const (
	TaskStatus_New TaskStatus = iota
	TaskStatus_InProgress
	TaskStatus_Done
	TaskStatus_Rejected
)

func (taskStatus TaskStatus) ToString() (str string, err error) {
	err = nil

	switch taskStatus {
	case TaskStatus_New:
		str = "New"
	case TaskStatus_InProgress:
		str = "InProgress"
	case TaskStatus_Done:
		str = "Done"
	case TaskStatus_Rejected:
		str = "Rejected"
	default:
		str = ""
		err = errors.New("coud not translate TaskStatus")
	}
	return
}

type Priority uint8

const (
	Priority_Low Priority = iota
	Priority_Medium
	Priority_High
)

func (priority Priority) ToString() (str string, err error) {
	err = nil

	switch priority {
	case Priority_Low:
		str = "Low"
	case Priority_Medium:
		str = "Medium"
	case Priority_High:
		str = "High"
	default:
		str = ""
		err = errors.New("coud not translate Priority state")
	}
	return
}

type TaskID int8

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
