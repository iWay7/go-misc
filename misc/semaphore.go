package misc

import "time"

type Semaphore struct {
	channel chan string
	permits int
}

func NewSemaphore(permits int) *Semaphore {
	channel := make(chan string, permits)
	return &Semaphore{
		channel: channel,
		permits: permits,
	}
}

func (semaphore *Semaphore) Acquire() {
	semaphore.channel <- ""
}

func (semaphore *Semaphore) Release() {
	_ = <-semaphore.channel
}

func (semaphore *Semaphore) TryAcquire() bool {
	select {
	case semaphore.channel <- "":
		return true
	default:
		return false
	}
}

func (semaphore *Semaphore) TryAcquireInDuration(timeout time.Duration) bool {
	for {
		select {
		case semaphore.channel <- "":
			return true
		case <-time.After(timeout):
			return false
		}
	}
}

func (semaphore *Semaphore) Available() int {
	return semaphore.permits - len(semaphore.channel)
}

func (semaphore *Semaphore) IsEmpty() bool {
	return len(semaphore.channel) >= semaphore.permits
}

func (semaphore *Semaphore) IsFull() bool {
	return len(semaphore.channel) <= 0
}
