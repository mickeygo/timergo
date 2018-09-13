package job

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/mickeygo/timergo/core/domain/entities/task"
)

// 获取所有待执行的任务
func getAllPendingTasks(t time.Time) *[]*task.Task {
	tasks := make([]*task.Task, 1000)
	for index := 0; index < 1000; index++ {
		tasks[index] = task.NewTask("", "https://www.baidu.com", "get", "", "")
	}

	return &tasks
}

func handleTask(ch chan *task.Task, index int) {
	for {
		t := <-ch
		resp, err := http.Get(t.Url)
		if err != nil {
			fmt.Println(err)
		}

		if resp.StatusCode >= 200 || resp.StatusCode <= 304 {
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err == nil {
				fmt.Println(string(body))
			}
		} else {
			// handle request error
			fmt.Println(resp.StatusCode)
		}

		// fmt.Printf("%d %v \n", index, *t)
	}
}
