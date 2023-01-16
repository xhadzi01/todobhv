package main

import (
	"encoding/json"
	"fmt"

	"github.com/xhadzi01/todobhv/taskManagement"
)

func main() {
	fmt.Println("Hiiii")
	task1 := taskManagement.NewTask("Some name", "Some Description", taskManagement.TaskPriority_High)
	task2 := taskManagement.NewTask("Some other name", "Some other Description", taskManagement.TaskPriority_Low)

	strBytes, _ := json.Marshal(task1)
	strBytes2, _ := json.Marshal(task2)

	fmt.Println(string(strBytes))
	fmt.Println(string(strBytes2))

	task3 := taskManagement.NewTask("AAA", "AAA", taskManagement.TaskPriority_Low)

	json.Unmarshal(strBytes, &task3)
	fmt.Println(task3.ToString())
}
