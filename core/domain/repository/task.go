package repository

import (
	"github.com/mickeygo/timergo/core/domain/entities/task"
)

// ITaskRepository Task 仓储接口
type ITaskRepository interface {
	Find(id string) *task.Task
	FindAll() []*task.Task

	Add(t *task.Task) error
	Update(t *task.Task) error
	Delete(t *task.Task) error
}