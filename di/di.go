package di

import (
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

// Definition ...
type Definition = di.Def
