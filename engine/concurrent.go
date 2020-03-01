package engine

// ConcurrentEngine 并发引擎
type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
	ItemChan    chan interface{}
}

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
		createWorker(e.Scheduler.WorkerChan(),
			out, e.Scheduler)
	}

	for _, r := range seeds {
		if isDuplicate(r.URL) {
			// log.Printf("Duplicate request: "+
			// 	"%s", r.URL)
			continue
		}
		e.Scheduler.Submit(r)
	}

	for {
		result := <-out
		for _, item := range result.Items {
			// if _, ok := item.(model.Profile); ok {
			// 	log.Printf("Got profile #%d:%v", profileCount, item)
			// 	profileCount++
			// }

			// if _, ok := item.(model.Profile); ok {
			// 	log.Printf("Got profile #%d: %v",
			// 		itemCount, item)
			// 	itemCount++

			// 	go func ()  {itemChan <- item}
			// }
			go func() { e.ItemChan <- item }()
		}

		// URL dedup
		for _, request := range result.Reuqests {
			if isDuplicate(request.URL) {
				// log.Printf("Duplicate request: "+
				// 	"%s", request.URL)
				continue
			}
			e.Scheduler.Submit(request)
		}
	}
}

func createWorker(in chan Request,
	out chan ParserResult, ready ReadyNotifier) {
	go func() {
		for {
			// tell scheduler i'm ready
			ready.WorkerReady(in)
			request := <-in
			result, err := worker(request)

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
