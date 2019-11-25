package presenter

import (
	"fmt"

	"github.com/kzmake/example-sarulabs-di/domain/entity"
	"github.com/kzmake/example-sarulabs-di/usecase/port"
)

// CreateTaskPresenter ...
type CreateTaskPresenter struct {
}

// NewCreateTaskPresenter ...
func NewCreateTaskPresenter() port.CreateTaskOutputPort {
	return &CreateTaskPresenter{}
}

// Execute ...
func (p *CreateTaskPresenter) Execute(t *entity.Task) (*port.CreateTaskOutputData, error) {
	dpo := &port.CreateTaskOutputData{
		Task: t,
	}

	fmt.Println("presenter...")

	return dpo, nil
}
