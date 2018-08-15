package task

import (
	"time"
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

	CreatedDate time.Time `json:"created_date"`

	UpdatedDate time.Time `json:"updated_date"`
}

func NewTask(name string, remote string) Task {
	return Task {
		Name: name
		Remote: remote
		IsDisabled: false
		CreatedDate: time.Now()
	}
}

// 启用任务
func (t *Task) Enable() {
	if (t.IsDisabled) {
		t.IsDisabled = false;
		t.UpdatedDate = time.Now()
	}
}

// 禁用任务
func (t *Task) Disable() {
	if (!t.IsDisabled) {
		t.IsDisabled = true;
		t.UpdatedDate = time.Now()
	}
}