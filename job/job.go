package job

import (
	"fmt"
	"time"
)

var interval = 10 * time.Second // 10 秒轮询间隔

// Job 定时器
type Job struct {
}

// Use 使用配置
func (j *Job) Use() {

}

// Run 启动 job
func (j *Job) Run() {
	fmt.Println("Job Run")

	Jobchan.Init(4, 1000)

	// ensureRunInZeroSecond()

	ticker := time.NewTicker(interval)
	for {
		t := <-ticker.C
		tasks := getAllPendingTasks(t)
		for index, task := range *tasks {
			Jobchan.C[index%Jobchan.Num] <- task
		}
	}
}

// 确保 job 开始运行的时间为 0 秒时间点
func ensureRunInZeroSecond() {
	ticker := time.NewTicker(time.Second)
	for {
		t := <-ticker.C
		if t.Second() == 0 {
			break
		}
	}

	ticker.Stop()
}
