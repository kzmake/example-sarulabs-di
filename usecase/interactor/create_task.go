package interactor

import (
	"fmt"

	"github.com/kzmake/example-sarulabs-di/domain/entity"
	"github.com/kzmake/example-sarulabs-di/usecase/port"
)

// CreateTaskInteractor ...
type CreateTaskInteractor struct {
}

// NewCreateTaskInteractor ...
func NewCreateTaskInteractor() port.CreateTaskInputPort {
	return &CreateTaskInteractor{}
}

// Execute ...
func (i *CreateTaskInteractor) Execute(input *port.CreateTaskInputData) (*entity.Task, error) {
	t := &entity.Task{
		ID:   "1",
		Data: "data",
	}

	fmt.Println("usecase...")

	return t, nil
}
