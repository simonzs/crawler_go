package scheduler

import (
	"crawler_go/engine"
)

// SimpleScheduler ...
type SimpleScheduler struct {
	workerChan chan engine.Request
}

// WorkerChan ...
func (s *SimpleScheduler) WorkerChan() chan engine.Request {
	return s.workerChan
}

// WorkerReady ...
func (s *SimpleScheduler) WorkerReady(w chan engine.Request) {
}

// Submit ...
func (s *SimpleScheduler) Submit(r engine.Request) {
	// Send request down to worker chan
	go func() { s.workerChan <- r }()
}

// Run ...
func (s *SimpleScheduler) Run() {
	s.workerChan = make(chan engine.Request)
}
