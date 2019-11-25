package controller

import (
	"fmt"

	"github.com/kzmake/example-sarulabs-di/usecase/port"
)

// TaskController ...
type TaskController struct {
	createIn  port.CreateTaskInputPort
	createOut port.CreateTaskOutputPort
}

// NewTaskController ...
func NewTaskController(input port.CreateTaskInputPort, output port.CreateTaskOutputPort) *TaskController {
	return &TaskController{
		createIn:  input,
		createOut: output,
	}
}

// CreateTask ...
func (c *TaskController) CreateTask(d string) (*port.CreateTaskOutputData, error) {
	in := &port.CreateTaskInputData{Data: d}
	_t, _ := c.createIn.Execute(in)

	fmt.Println("controller...")

	return c.createOut.Execute(_t)
}
