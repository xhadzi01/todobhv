package main

import (
	"encoding/json"
	"fmt"

	"github.com/xhadzi01/todobhv/taskManagement"
)

func main() {
	////////////////////////////////////////////////////////
	///////////////////TASK MARSHALING//////////////////////
	task1 := taskManagement.NewTask("Some name", "Some Description", taskManagement.TaskPriority_High)
	task2 := taskManagement.NewTask("Some other name", "Some other Description", taskManagement.TaskPriority_Low)

	strBytes, _ := json.Marshal(task1)
	strBytes2, _ := json.Marshal(task2)

	fmt.Println(string(strBytes))
	fmt.Println(string(strBytes2))

	task3 := taskManagement.NewTask("AAA", "AAA", taskManagement.TaskPriority_Low)

	json.Unmarshal(strBytes, &task3)
	////////////////////////////////////////////////////////
	//////////////DATA MANAGEMENT///////////////////////////

	dm := taskManagement.NewDataManagement(`C:\projectbins\Go\todobhv\testing`, `jackson`)

	dm.SaveTask(task1)
	dm.SaveTask(task2)

	task4, _ := dm.LoadTask(1)
	fmt.Println(task3.ToString())
	fmt.Println(task4.ToString())
}
