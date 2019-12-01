package di

import (
	"github.com/spf13/cobra"

	"github.com/kzmake/example-sarulabs-di/cli"
)

// NewCLIModules ...
func NewCLIModules() []Definition {
	defs := []Definition{
		{
			Name:  "app",
			Scope: App,
			Build: func(ctn Container) (interface{}, error) {
				println("build app")
				builder := ctn.Get("builder").(cli.Builder)
				targets := []map[string]*cobra.Command{
					ctn.Get("task").(map[string]*cobra.Command),
				}
				return builder.Targets(targets...).Build(), nil
			},
		},
		{
			Name:  "builder",
			Scope: App,
			Build: func(ctn Container) (interface{}, error) {
				println("build builder")
				root := ctn.Get("root").(*cobra.Command)
				operations := ctn.Get("operations").([]*cobra.Command)
				return cli.New().Root(root).Operations(operations...), nil
			},
		},
		{
			Name:  "root",
			Scope: App,
			Build: func(ctn Container) (interface{}, error) {
				println("build root")
				return cli.NewRootCommand(), nil
			},
		},
		{
			Name:  "operations",
			Scope: App,
			Build: func(ctn Container) (interface{}, error) {
				println("build operations")
				return cli.NewOperationCommands(), nil
			},
		},
		{
			Name:  "task",
			Scope: App,
			Build: func(ctn Container) (interface{}, error) {
				println("build task")
				return cli.NewTaskCommand(), nil
			},
		},
	}

	return defs
}
