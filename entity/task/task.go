package task

import (
	"time"
	"github.com/mickeygo/timergo/entity"
)

// 执行的任务
type Task struct {
	entity.Entity

	// 任务名称
	Name		string		`json:"name"`

	// 调用地址(scheme + host + [port] + route + [params])
	Url			string 		`json:"url"`

	// 地址调用方式，有 get | post | put | delete
	Method		string		`json:"method"`

	// 执行的内容类型，有 application/x-www-form-urlencoded | application/json
	ContentType	string 		`json:"content-type"`

	// 执行的参数，使用 json 格式。针对不同类型的会转换为不同的参数形式
	Params		string		`json:“params”`

	// 是否已禁用
	IsDisabled 	bool		`json:"disabled"`

	// 是否已删除
	IsDeleted	bool		`json:"deleted"`

	CreatedDate time.Time 	`json:"created_date"`

	UpdatedDate time.Time 	`json:"updated_date"`
}

// 新建一个任务
func NewTask(name string, url string, method string, contentType string, params string) *Task {
	return &Task {
		Name: name,
		Url: url
		Method: method,
		ContentType: contentType,
		Params: params,
		IsDisabled: false,
		IsDeleted: false,
		CreatedDate: time.Now()
	}
}

// 启用任务
func (t *Task) Enable() {
	if (t.IsDisabled) {
		t.IsDisabled = false
		t.UpdatedDate = time.Now()
	}
}

// 禁用任务
func (t *Task) Disable() {
	if (!t.IsDisabled) {
		t.IsDisabled = true
		t.UpdatedDate = time.Now()
	}
}

// 删除任务
func (t *Task) Delete() {
	if (!t.IsDeleted) {
		t.IsDeleted = true
		t.UpdatedDate = time.Now()
	}
}

// 生成新的执行状态
func (t *Task) NewState(start time.Time, end time.Time, result bool, resultCode string) *TaskState {
	return &TaskState {
		TaskId: t.Id,
		ExecuteStart: start,
		ExecuteEnd: end,
		Result: result,
		ResultCode: resultCode,
		CreatedDate: time.Now()
	}
}