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

	num := 4  // 处理消息通道数
	var chans = make([]chan int, num)
	for index := 0; index < num; index++ {
		chans[index] = make(chan int, 1000)

		go func(ch chan int, idx int) {
			for {
				i := <-ch
				fmt.Printf("%d -- %d \r\n", idx, i)
			}
		}(chans[index], index)
	}

	// ensureRunInZeroSecond()

	ticker := time.NewTicker(interval)
	for {
		t := <-ticker.C
		fmt.Println(t)

		func(count int) {
			for index := 0; index < count; index++ {
				chans[index % len(chans)] <-index
			}
		}(1000)
	}
}

// 处理事件
func handle(t time.Time, ch chan int) {
	// 获取当前时间待执行的任务
}

// 确保 job 开始运行的时间为 0 秒时间点
func ensureRunInZeroSecond() {
	ticker := time.NewTicker(time.Second)
	for	{
		t := <-ticker.C
		if (t.Second() == 0) {
			break
		}
	}

	ticker.Stop()
}