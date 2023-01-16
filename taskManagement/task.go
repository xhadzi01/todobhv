package taskManagement

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

// Task represents singular task that could be described as collection
// of attributes like 'name', 'description', 'priority' and so on.
type Task struct {
	ID             TaskID       `json:"ID"`
	Name           string       `json:"Name"`
	Description    string       `json:"Description"`
	CreationDate   string       `json:"CreationDate"`
	CompletionDate *string      `json:"CompletionDate"`
	Status         TaskStatus   `json:"Status"`
	Priority       TaskPriority `json:"Priority"`
}

func NewTask(name, description string, initialPriority TaskPriority) (task *Task) {
	task = &Task{}
	task.ID = 1
	task.Name = name
	task.Description = description
	task.CreationDate = time.Now().Format("15:04:05 02-01-2006")
	task.CompletionDate = nil
	task.Status = TaskStatus_New
	task.Priority = initialPriority
	return
}

func (task *Task) ToString() (str string, err error) {
	str = ""
	err = errors.New("Task instance is invalid")
	if task == nil {
		return
	}

	priorityStr, priorityErr := task.Priority.ToString()
	if priorityErr != nil {
		priorityStr = "Unknown"
	}

	statusStr, statusErr := task.Status.ToString()
	if statusErr != nil {
		statusStr = "Unknown"
	}

	completionDateStr := "Not Completed Yet"
	if task.CompletionDate != nil {
		completionDateStr = *task.CompletionDate
	}

	strBuilder := &strings.Builder{}
	_, err = fmt.Fprintf(strBuilder,
		"Task{ ID: %v, Name: '%v', Description: '%v', CreationDate: %v, CompletionDate: %v, Status: %v, Priority: %v",
		task.ID,
		task.Name,
		task.Description,
		task.CreationDate,
		completionDateStr,
		statusStr,
		priorityStr)

	if err == nil {
		str = strBuilder.String()
	}
	return
}

func (task *Task) GetID() TaskID {
	return task.ID
}

func (task *Task) GetName() string {
	return task.Name
}

func (task *Task) GetDescription() string {
	return task.Description
}

func (task *Task) GetState() TaskStatus {
	return task.Status
}

func (task *Task) GetPriority() TaskPriority {
	return task.Priority
}

func (task *Task) GetDateCreated() string {
	return task.CreationDate
}

func (task *Task) GetDateCompleted() *string {
	return task.CompletionDate
}

func (task *Task) SetName(newName string) error {
	// here will be sanity check
	task.Name = newName
	return nil
}

func (task *Task) SetDescription(newDescription string) error {
	// here will be sanity check
	task.Description = newDescription
	return nil
}

func (task *Task) SetState(newState TaskStatus) {
	task.Status = newState
}

func (task *Task) SetPriority(newPriority TaskPriority) {
	task.Priority = newPriority
}
