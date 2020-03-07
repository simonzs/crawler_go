package engine

// ConcurrentEngine 并发引擎
type ConcurrentEngine struct {
	Scheduler        Scheduler
	WorkerCount      int
	ItemChan         chan Item
	ReqeustProcessor Processor
}

// Processor ...
type Processor func(Request) (ParserResult, error)

// Scheduler ...
type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	WorkerChan() chan Request
	Run()
}

// ReadyNotifier ...
type ReadyNotifier interface {
	WorkerReady(chan Request)
}

// Run 实现的是并发版爬虫架构
func (e *ConcurrentEngine) Run(seeds ...Request) {
	out := make(chan ParserResult)
	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++ {
		e.createWorker(e.Scheduler.WorkerChan(),
			out, e.Scheduler)
	}

	for _, r := range seeds {
		if isDuplicate(r.URL) {
			continue
		}
		e.Scheduler.Submit(r)
	}

	for {
		result := <-out
		for _, item := range result.Items {
			go func() { e.ItemChan <- item }()
		}

		// URL dedup
		for _, request := range result.Reuqests {
			if isDuplicate(request.URL) {
				continue
			}
			e.Scheduler.Submit(request)
		}
	}
}

func (e *ConcurrentEngine) createWorker(
	in chan Request,
	out chan ParserResult, ready ReadyNotifier) {
	go func() {
		for {
			// tell scheduler i'm ready
			ready.WorkerReady(in)
			request := <-in
			result, err := e.ReqeustProcessor(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}

var visitedUrls = make(map[string]bool)

func isDuplicate(url string) bool {
	if visitedUrls[url] {
		return true
	}
	visitedUrls[url] = true
	return false
}
