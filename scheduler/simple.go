package scheduler

import (
	"crawler_go/engine"
)

// SimpleScheduler ...
type SimpleScheduler struct {
	workerChan chan engine.Request
}

// ConfigureMasterWorkerChan ...
func (s *SimpleScheduler) ConfigureMasterWorkerChan(c chan engine.Request) {
	s.workerChan = c
}

// Submit ...
func (s *SimpleScheduler) Submit(r engine.Request) {
	// Send request down to worker chan
	go func() { s.workerChan <- r }()
}
