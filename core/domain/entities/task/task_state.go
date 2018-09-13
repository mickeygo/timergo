package task

import (
	"time"
)

// TaskState 任务执行的状态
type TaskState struct {
	Id			string		`db:"id"`

	// 任务 Id
	TaskId		string		`db:"task_id"`

	// 执行开始时间
	ExecuteStart time.Time  `db:"execute_start"`

	// 执行结束时间
	ExecuteEnd	time.Time	`db:"execute_end"`

	// 执行结果
	Result		bool		`db:"result"`

	// 执行结果代码
	ResultCode	string 		`db:"result_code"`

	CreatedDate time.Time 	`db:"created_date"`
}