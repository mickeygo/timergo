package job

import (
	"github.com/mickeygo/timergo/core/domain/entity/task"
)

var Jobchan = jobchan{}

// Job 通道
type jobchan struct {
	// job channel
	C 			[]chan *task.Task
	Num 		int
	Capacity 	int
}

// 初始化
func (jc * jobchan) Init(num int, capacity int) {
	jc.C = make([]chan *task.Task, num)
	jc.Num = num
	jc.Capacity = capacity

	for index := 0; index < num; index++ {
		jc.C[index] = make(chan *task.Task, capacity)
		go handleTask(jc.C[index], index)
	}	
}

