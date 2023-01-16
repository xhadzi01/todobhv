package taskManagement

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type DataManagement struct {
	dataFolder string
}

func NewDataManagement(basepath string, username string) (dm *DataManagement) {
	folderPath := filepath.Join(basepath, username)

	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		os.Mkdir(folderPath, os.ModePerm)
	}

	return &DataManagement{folderPath}
}

func (dm *DataManagement) LoadTask(id TaskID) (task *Task, err error) {
	task = nil
	filename := fmt.Sprintf("Task_%v.json", id)
	filepath := filepath.Join(dm.dataFolder, filename)
	data, err := os.ReadFile(filepath)
	if err == nil {
		reader := strings.NewReader(string(data))
		decoder := json.NewDecoder(reader)

		task = &Task{}
		decoder.Decode(task)
	}

	return
}

func (dm *DataManagement) SaveTask(task *Task) (err error) {
	filename := fmt.Sprintf("Task_%v.json", task.ID)
	filepath := filepath.Join(dm.dataFolder, filename)

	// if someone deleted folder with all tasks, recreate folder before writing task
	if _, err := os.Stat(dm.dataFolder); os.IsNotExist(err) {
		os.Mkdir(dm.dataFolder, os.ModePerm)
	}

	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE, 0666)
	if err == nil {
		defer file.Close()

		encoder := json.NewEncoder(file)
		encoder.Encode(task)
	}
	return
}
