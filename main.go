package main

import (
	"github.com/kzmake/example-sarulabs-di/interface/controller"

	"github.com/kzmake/example-sarulabs-di/di"
)

func main() {
	ctn, _ := di.NewContainer(di.CoreDefs...)

	c := ctn.Get("controller").(*controller.TaskController)

	o, _ := c.CreateTask("hoge")
}
