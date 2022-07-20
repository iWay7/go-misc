package misc

import (
	"container/list"
	"sync"
)

var receiverList = list.New()
var listLocker = sync.Mutex{}

func RegisterEventReceiver(receiver func(event string, data interface{})) {
	listLocker.Lock()
	receiverList.PushBack(receiver)
	listLocker.Unlock()
}

func BroadcastEvent(event string, data interface{}) {
	var funcArray []func(event string, data interface{})
	listLocker.Lock()
	for element := receiverList.Front(); element != nil; element = element.Next() {
		funcArray = append(funcArray, element.Value.(func(event string, data interface{})))
	}
	listLocker.Unlock()
	for _, funcItem := range funcArray {
		funcItem(event, data)
	}
}
