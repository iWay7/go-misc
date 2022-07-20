package misc

import (
	"sync"
	"time"
)

type Task struct {
	NextExecTimeInSeconds int64
	ExecFunc              func() int64
}

var taskList = []Task{}
var taskListLocker = sync.Mutex{}

func ScheduleTask(firstExecTimeInSeconds int64, execFunc func() int64) {
	taskListLocker.Lock()
	taskList = append(taskList, Task{
		NextExecTimeInSeconds: firstExecTimeInSeconds,
		ExecFunc:              execFunc,
	})
	taskListLocker.Unlock()
}

func runTasks() {
	for {
		taskListLocker.Lock()
		oldTaskList := taskList
		taskList = []Task{}
		taskListLocker.Unlock()
		newTaskList := []Task{}
		nowInSeconds := NowInSeconds()
		for _, task := range oldTaskList {
			if nowInSeconds >= task.NextExecTimeInSeconds {
				nextTimeInSeconds := task.ExecFunc()
				if nextTimeInSeconds >= 0 {
					task.NextExecTimeInSeconds = nextTimeInSeconds
					newTaskList = append(newTaskList, task)
				}
			} else {
				newTaskList = append(newTaskList, task)
			}
		}
		taskListLocker.Lock()
		taskList = append(taskList, newTaskList...)
		taskListLocker.Unlock()
		time.Sleep(time.Second)
	}
}

func initScheduler() {
	go runTasks()
}
