package di

import (
	"fmt"

	"github.com/sarulabs/di"
)

// NewContainer ...
func NewContainer(defs ...Definition) (Container, error) {
	builder, err := di.NewBuilder()
	if err != nil {
		return nil, fmt.Errorf("di error: %v", err)
	}

	if err := builder.Add(defs...); err != nil {
		return nil, fmt.Errorf("di error: %v", err)
	}

	return builder.Build(), nil
}
