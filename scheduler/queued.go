package scheduler

import (
	"github.com/simonzs/crawler_go/engine"
)

// QueuedScheduler ...
type QueuedScheduler struct {
	requestChan chan engine.Request
	workerChan  chan chan engine.Request
}

// WorkerChan ...
func (s *QueuedScheduler) WorkerChan() chan engine.Request {
	return make(chan engine.Request)
}

// Submit ...
func (s *QueuedScheduler) Submit(r engine.Request) {
	// Send request down to worker chan
	s.requestChan <- r
}

// WorkerReady ...
func (s *QueuedScheduler) WorkerReady(w chan engine.Request) {
	s.workerChan <- w
}

// ConfigureMasterWorkerChan ...
func (s *QueuedScheduler) ConfigureMasterWorkerChan(c chan engine.Request) {
	s.requestChan = c
}

// Run ...
func (s *QueuedScheduler) Run() {
	s.workerChan = make(chan chan engine.Request)
	s.requestChan = make(chan engine.Request)
	go func() {
		var requestQ []engine.Request
		var workerQ []chan engine.Request
		for {
			var activeRequet engine.Request
			var activeWorker chan engine.Request
			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeRequet = requestQ[0]
				activeWorker = workerQ[0]
			}
			select {
			case r := <-s.requestChan:
				requestQ = append(requestQ, r)
			case w := <-s.workerChan:
				workerQ = append(workerQ, w)
			case activeWorker <- activeRequet:
				requestQ = requestQ[1:]
				workerQ = workerQ[1:]
			}
		}
	}()
}
