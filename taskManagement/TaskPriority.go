package taskManagement

import (
	"errors"
)

type TaskPriority uint8

const (
	TaskPriority_Low TaskPriority = iota
	TaskPriority_Medium
	TaskPriority_High
)

func (priority TaskPriority) ToString() (str string, err error) {
	err = nil

	switch priority {
	case TaskPriority_Low:
		str = "Low"
	case TaskPriority_Medium:
		str = "Medium"
	case TaskPriority_High:
		str = "High"
	default:
		str = ""
		err = errors.New("coud not translate Priority state")
	}
	return
}
