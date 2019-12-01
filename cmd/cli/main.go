package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/kzmake/example-sarulabs-di/cli"
	"github.com/kzmake/example-sarulabs-di/di"
)

func main() {
	modules := []di.Definition{}
	modules = append(modules, di.NewCLIModules()...)
	modules = append(modules, di.NewCoreModules()...)

	ctn, err := di.NewContainer(modules...)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	task := ctn.Get("task").(map[string]*cobra.Command)
	task["new"].RunE = scoped(ctn, func(ctn di.Container) interface{} { return ctn.Get("handler") })

	app := ctn.Get("app").(cli.App)
	app.Run()
}

func scoped(ctn di.Container, invoke func(ctn di.Container) interface{}) func(*cobra.Command, []string) error {
	fn := func(cmd *cobra.Command, args []string) error {
		req, err := ctn.SubContainer()
		if err != nil {
			cmd.Println(err)
			return err
		}

		return invoke(req).(func(*cobra.Command, []string) error)(cmd, args)
	}

	return fn
}
