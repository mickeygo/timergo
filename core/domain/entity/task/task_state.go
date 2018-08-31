package task

import (
	"time"
)

// TaskState 任务执行的状态
type TaskState struct {
	Id			string		`json:"id"`

	// 任务 Id
	TaskId		string		`json:"task_id"`

	// 执行开始时间
	ExecuteStart time.Time  `json:"execute_start"`

	// 执行结束时间
	ExecuteEnd	time.Time	`json:"execute_end"`

	// 执行结果
	Result		bool		`json:"result"`

	// 执行结果代码
	ResultCode	string 		`json:"result_code"`

	CreatedDate time.Time 	`json:"created_date"`
}