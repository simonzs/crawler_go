package engine

import "log"

// ConcurrentEngine 并发引擎
type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

// Scheduler ...
type Scheduler interface {
	Submit(Request)
	ConfigureMasterWorkerChan(chan Request)
	WorkerReady(chan Request)
	Run()
}

// Run 实现的是并发版爬虫架构
func (e *ConcurrentEngine) Run(seeds ...Request) {
	out := make(chan ParserResult)
	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++ {
		createWorker(out, e.Scheduler)
	}

	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	itemCount := 0

	for {
		result := <-out
		for _, item := range result.Items {
			log.Printf("Got item %d:%v", itemCount, item)
			itemCount++
		}
		for _, request := range result.Reuqests {
			e.Scheduler.Submit(request)
		}
	}
}

func createWorker(
	out chan ParserResult, s Scheduler) {

	in := make(chan Request)
	go func() {
		for {
			// tell scheduler i'm ready
			s.WorkerReady(in)
			request := <-in
			result, err := worker(request)

			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
