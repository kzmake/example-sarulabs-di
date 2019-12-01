package di

import (
	"github.com/kzmake/example-sarulabs-di/infrastructure/handler"
	"github.com/kzmake/example-sarulabs-di/infrastructure/renderer"
	"github.com/kzmake/example-sarulabs-di/interface/controller"
	"github.com/kzmake/example-sarulabs-di/interface/presenter"
	"github.com/kzmake/example-sarulabs-di/usecase/interactor"
	"github.com/kzmake/example-sarulabs-di/usecase/port"
)

// NewCoreModules ...
func NewCoreModules() []Definition {
	defs := []Definition{
		{
			Name:  "handler",
			Scope: Request,
			Build: func(ctn Container) (interface{}, error) {
				println("build handler")
				controller := ctn.Get("controller").(*controller.TaskController)
				renderer := ctn.Get("renderer").(*renderer.TaskRenderer)

				return handler.NewCreateTaskHandler(controller, renderer), nil
			},
		},
		{
			Name:  "controller",
			Scope: Request,
			Build: func(ctn Container) (interface{}, error) {
				println("build controller")
				input := ctn.Get("input").(port.CreateTaskInputPort)
				output := ctn.Get("output").(port.CreateTaskOutputPort)

				return controller.NewTaskController(input, output), nil
			},
		},
		{
			Name:  "renderer",
			Scope: Request,
			Build: func(ctn Container) (interface{}, error) {
				println("build renderer")
				return renderer.NewTaskRenderer(), nil
			},
		},
		{
			Name:  "input",
			Scope: Request,
			Build: func(ctn Container) (interface{}, error) {
				println("build input")
				return interactor.NewCreateTaskInteractor(), nil
			},
		},
		{
			Name:  "output",
			Scope: Request,
			Build: func(ctn Container) (interface{}, error) {
				println("build output")
				return presenter.NewCreateTaskPresenter(), nil
			},
		},
	}

	return defs
}
