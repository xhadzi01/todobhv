package taskManagement

type IDataManagement interface {
	LoadTask(TaskID) (*Task, error)
	SaveTask(*Task) error
}
