package handler

import (
	"github.com/spf13/cobra"

	"github.com/kzmake/example-sarulabs-di/infrastructure/renderer"
	"github.com/kzmake/example-sarulabs-di/interface/controller"
)

// NewCreateTaskHandler ...
func NewCreateTaskHandler(c *controller.TaskController, r *renderer.TaskRenderer) func(*cobra.Command, []string) error {
	fn := func(cmd *cobra.Command, args []string) error {
		data, err := cmd.PersistentFlags().GetString("data")
		if err != nil {
			return err
		}

		o := c.CreateTask(data)
		r.CreateTask(cmd, o)

		return nil
	}

	return fn
}
