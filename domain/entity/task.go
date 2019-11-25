package entity

import (
	"github.com/kzmake/example-sarulabs-di/domain/vo"
)

// Task ...
type Task struct {
	ID   vo.TaskID
	Data vo.TaskData
}
