package state

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"
)

// Task represents singular task that could be described as collection
// of attributes like 'name', 'description', 'priority' and so on.
type Task struct {
	ID             TaskID
	name           string
	description    string
	creationDate   time.Time
	completionDate *time.Time
	status         TaskStatus
	priority       Priority
}

// We want to ue automatic struct decay into json. Not used for other uses.
type ExposedTask struct {
	ID             TaskID     `json:"ID"`
	Name           string     `json:"Name"`
	Description    string     `json:"Description"`
	CreationDate   time.Time  `json:"CreationDate"`
	CompletionDate *time.Time `json:"CompletionDate"`
	Status         TaskStatus `json:"Status"`
	Priority       Priority   `json:"Priority"`
}

func NewTask(name, description string, initialPriority Priority) (task *Task) {
	task = &Task{}
	task.ID = 1
	task.name = name
	task.description = description
	task.creationDate = time.Now()
	task.completionDate = nil
	task.status = TaskStatus_New
	task.priority = initialPriority
	return
}

func (task *Task) ToString() (str string, err error) {
	str = ""
	err = errors.New("Task instance is invalid")
	if task == nil {
		return
	}

	priorityStr, priorityErr := task.priority.ToString()
	if priorityErr != nil {
		priorityStr = "Unknown"
	}

	statusStr, statusErr := task.status.ToString()
	if statusErr != nil {
		statusStr = "Unknown"
	}

	completionDateStr := "Not Completed Yet"
	if task.completionDate != nil {
		completionDateStr = fmt.Sprintf("%v", *task.completionDate)
	}

	strBuilder := &strings.Builder{}
	_, err = fmt.Fprintf(strBuilder,
		"Task{ ID: %v, Name: '%v', Description: '%v', CreationDate: %v, CompletionDate: %v, Status: %v, Priority: %v",
		task.ID,
		task.name,
		task.description,
		task.creationDate,
		completionDateStr,
		statusStr,
		priorityStr)

	if err == nil {
		str = strBuilder.String()
	}
	return
}

func (task *Task) ToJson() ([]byte, error) {
	if task == nil {
		return make([]byte, 0), errors.New("Task instance is invalid")
	}

	exposed := ExposedTask{}
	exposed.ID = task.ID
	exposed.Name = task.name
	exposed.Description = task.description
	exposed.CreationDate = task.creationDate
	exposed.CompletionDate = task.completionDate
	exposed.Status = task.status
	exposed.Priority = task.priority

	return json.Marshal(exposed)
}

func LoadTask(data []byte) (task *Task, err error) {
	exposed := &ExposedTask{}
	task = nil
	err = json.Unmarshal(data, exposed)

	if err != nil {
		return
	}

	task = &Task{}
	task.ID = exposed.ID
	task.name = exposed.Name
	task.description = exposed.Description
	task.creationDate = exposed.CreationDate
	task.completionDate = exposed.CompletionDate
	task.status = exposed.Status
	task.priority = exposed.Priority
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
