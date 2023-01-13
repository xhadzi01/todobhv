package state

import (
	"encoding/json"
	"time"
)

type TaskStatus uint8

const (
	TaskStatus_New TaskStatus = iota
	TaskStatus_InProgress
	TaskStatus_Done
	TaskStatus_Rejected
)

type Priority uint8

const (
	Priority_Low Priority = iota
	Priority_Medium
	Priority_High
)

type TaskID int8

// Task represents singular task that could be described as collection
// of attributes like 'name', 'description', 'priority' and so on.
type Task struct {
	ID             TaskID     `json:"ID"`
	name           string     `json:"Name"`
	description    string     `json:"Description"`
	creationDate   time.Time  `json:"CreationDate"`
	completionDate *time.Time `json:"CompletionDate"`
	status         TaskStatus `json:"Status"`
	priority       Priority   `json:"Priority"`
}

func (task Task) ToJson() ([]byte, error) {
	return json.Marshal(task)
}

func NewTask(name, description string) (task Task) {
	task.ID = 1
	task.name = name
	task.description = description
	task.creationDate = time.Now()
	task.completionDate = nil
	task.status = TaskStatus_New
	task.priority = Priority_Medium
	return
}

func (task *Task) GetID() TaskID {
	return task.ID
}

func (task *Task) GetName() string {
	return task.name
}

func (task *Task) GetDescription() string {
	return task.description
}

func (task *Task) GetState() TaskStatus {
	return task.status
}

func (task *Task) GetPriority() Priority {
	return task.priority
}

func (task *Task) GetDateCreated() time.Time {
	return task.creationDate
}

func (task *Task) GetDateCompleted() *time.Time {
	return task.completionDate
}

func (task *Task) SetName(newName string) error {
	// here will be sanity check
	task.name = newName
	return nil
}

func (task *Task) SetDescription(newDescription string) error {
	// here will be sanity check
	task.description = newDescription
	return nil
}

func (task *Task) SetState(newState TaskStatus) {
	task.status = newState
}

func (task *Task) SetPriority(newPriority Priority) {
	task.priority = newPriority
}
