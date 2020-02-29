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
}

// Run 实现的是并发版爬虫架构
func (e *ConcurrentEngine) Run(seeds ...Request) {
	in := make(chan Request)
	out := make(chan ParserResult)
	e.Scheduler.ConfigureMasterWorkerChan(in)

	for i := 0; i < e.WorkerCount; i++ {
		createWorker(in, out)
	}

	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}
	for {
		result := <-out
		for _, item := range result.Items {
			log.Printf("Got item %v", item)
		}
		for _, request := range result.Reuqests {
			e.Scheduler.Submit(request)
		}
	}
}

func createWorker(in chan Request, out chan ParserResult) {
	go func() {
		for {
			request := <-in
			result, err := worker(request)

			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
