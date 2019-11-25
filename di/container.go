package di

import (
	"fmt"

	"github.com/sarulabs/di"
)

const (
	// App is the name of the application scope.
	App = di.App
	// Request is the name of the request scope.
	Request = di.Request
	// SubRequest is the name of the subrequest scope.
	SubRequest = di.SubRequest
)

// Container ...
type Container = di.Container

// Def ...
type Def = di.Def

// NewContainer ...
func NewContainer(defs ...Def) (Container, error) {
	builder, err := di.NewBuilder()
	if err != nil {
		return nil, fmt.Errorf("di error: %v", err)
	}

	if err := builder.Add(defs...); err != nil {
		return nil, fmt.Errorf("di error: %v", err)
	}

	return builder.Build(), nil
}
