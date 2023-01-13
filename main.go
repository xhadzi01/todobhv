package main

import (
	"fmt"

	"github.com/xhadzi01/todobhv/state"
)

func main() {
	fmt.Println("Hiiii")
	task1 := state.NewTask("Some name", "Some Description", state.Priority_High)
	task2 := state.NewTask("Some other name", "Some other Description", state.Priority_Low)

	strBytes, _ := task1.ToJson()
	strBytes2, _ := task2.ToJson()

	fmt.Println(string(strBytes))
	fmt.Println(string(strBytes2))

	task3, _ := state.LoadTask(strBytes)
	fmt.Println(task3.ToString())
}
