package task

// 执行的任务
type Task struct {
	Id		string	`json:"id"`
	Name	string	`json:"name"`

	// 远程调用地址
	Remote	string	`json:"remote"`

	// 是否禁用
	IsDisabled bool	`json:"disabled"`
}

// 添加任务
func Add(t Task) Task {
	return t
}

// 禁用任务
func Disable(id string) bool {
	return true
}

// 禁用任务
func (t *Task)Disable() bool {
	if (t.IsDisabled) {
		return true
	}

	t.IsDisabled = true

	return true
}

// 启用任务
func Enable(id string) bool {
	return true
}

// 启用任务
func (t *Task) Enable() bool  {
	if (!t.IsDisabled) {
		return true
	}

	t.IsDisabled = false

	return true
}