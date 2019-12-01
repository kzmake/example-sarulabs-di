package presenter

import (
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
func (p *CreateTaskPresenter) Execute(t *entity.Task) *port.CreateTaskOutputData {
	payload := &port.CreateTaskOutputData{
		Task: t,
	}

	return payload
}
