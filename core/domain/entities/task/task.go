package task

import (
	"time"
	"github.com/mickeygo/timergo/core/domain/entities"
)

// Task 执行的任务
type Task struct {
	// 实体 Id
	Id 			string 		`db:"id"`

	// 任务名称
	Name		string		`db:"name"`

	// 调用地址(scheme + host + [port] + route + [params])
	Url			string 		`db:"url"`

	// 地址调用方式，有 get | post | put | delete
	Method		string		`db:"method"`

	// 执行的内容类型，有 application/x-www-form-urlencoded | application/json
	ContentType	string 		`db:"content_type"`

	// 执行的参数，使用 json 格式。针对不同类型的会转换为不同的参数形式
	Params		string		`db:"params"`

	// 是否已禁用
	IsDisabled 	bool		`db:"disabled"`

	// 是否已删除
	IsDeleted	bool		`db:"deleted"`

	// 计数器，记录任务执行的次数
	Count		int32		`db:"count"`

	CreatedDate time.Time 	`db:"created_date"`

	UpdatedDate time.Time 	`db:"updated_date"`
}

// NewTask 新建一个任务
func NewTask(name string, url string, method string, contentType string, params string) *Task {
	return &Task {
		Id: entities.NewID(),
		Name: name,
		Url: url,
		Method: method,
		ContentType: contentType,
		Params: params,
		IsDisabled: false,
		IsDeleted: false,
		Count: 0,
		CreatedDate: time.Now(),
	}
}

// IsNew 是否是新任务（任务还未开始执行）
func IsNew(t *Task) bool {
	return t.Count == 0
}

// Enable 启用任务
func (t *Task) Enable() {
	if (t.IsDisabled) {
		t.IsDisabled = false
		t.UpdatedDate = time.Now()
	}
}

// Disable 禁用任务
func (t *Task) Disable() {
	if (!t.IsDisabled) {
		t.IsDisabled = true
		t.UpdatedDate = time.Now()
	}
}

// Delete 删除任务
func (t *Task) Delete() {
	if (!t.IsDeleted) {
		t.IsDeleted = true
		t.UpdatedDate = time.Now()
	}
}

// NewState 生成新的执行状态
func (t *Task) NewState(start time.Time, end time.Time, result bool, resultCode string) *TaskState {
	return &TaskState {
		Id: entities.NewID(),
		TaskId: t.Id,
		ExecuteStart: start,
		ExecuteEnd: end,
		Result: result,
		ResultCode: resultCode,
		CreatedDate: time.Now(),
	}
}