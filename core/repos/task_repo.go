package repos

import (
	"github.com/mickeygo/timergo/core/domain/entity/task"
)

// Task 仓储接口
type TaskRepository interface {
	Find(id string) *task.Task
	FindAll() []*task.Task
}
