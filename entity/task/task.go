package task

import (
	"github.com/mickeygo/timergo/entity"
)

// 执行的任务
type Task struct {
	entity.Entity

	// 任务名称
	Name		string	`json:"name"`

	// 远程调用地址
	Remote		string	`json:"remote"`

	// 是否禁用
	IsDisabled 	bool	`json:"disabled"`
}