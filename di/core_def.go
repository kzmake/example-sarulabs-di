package di

import (
	"github.com/kzmake/example-sarulabs-di/interface/controller"
	"github.com/kzmake/example-sarulabs-di/interface/presenter"
	"github.com/kzmake/example-sarulabs-di/usecase/interactor"
	"github.com/kzmake/example-sarulabs-di/usecase/port"
)

// CoreDefs ...
var CoreDefs = []Def{
	{
		Name:  "input",
		Scope: App,
		Build: func(ctn Container) (interface{}, error) {
			return interactor.NewCreateTaskInteractor(), nil
		},
	},
	{
		Name:  "output",
		Scope: App,
		Build: func(ctn Container) (interface{}, error) {
			return presenter.NewCreateTaskPresenter(), nil
		},
	},
	{
		Name:  "controller",
		Scope: App,
		Build: func(ctn Container) (interface{}, error) {
			input := ctn.Get("input").(port.CreateTaskInputPort)
			output := ctn.Get("output").(port.CreateTaskOutputPort)

			return controller.NewTaskController(input, output), nil
		},
	},
}
