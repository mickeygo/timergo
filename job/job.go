package job

import (
	"time"
	"fmt"
)

var interval time.Duration = 10 * time.Second // 10 秒轮询间隔

type Job struct {

}

func (j *Job) Use() {

}

// 启动 job
func (j *Job) Run() {
	fmt.Println("Job Run")

	ticker := time.NewTicker(interval)
	for {
		t := <- ticker.C
		handler(t)
	}
}

// 处理事件
func handler(t time.Time) {
	fmt.Println(t)
}