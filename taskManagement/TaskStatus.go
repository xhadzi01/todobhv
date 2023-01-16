package taskManagement

import "errors"

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
