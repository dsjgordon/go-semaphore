// Package semaphore provides a simple semaphore utility
package semaphore

type empty struct{}

type Semaphore struct {
	queue chan empty
}

// New creates a new semaphore with the given concurrency
func New(concurrency uint) *Semaphore {
	return &Semaphore{
		make(chan empty, concurrency),
	}
}

// Yield notifies the semaphore of intention to work
func (semaphore *Semaphore) Yield() {
	semaphore.queue <- empty{}
}

// Unyield notifies the semaphore of work completed
func (semaphore *Semaphore) Unyield() {
	<-semaphore.queue
}