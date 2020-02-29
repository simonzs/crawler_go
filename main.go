package main

import (
	"crawler_go/engine"
	"crawler_go/zhenai/parser"
)

const baseURL = "https://www.zhenai.com/zhenghun"

func main() {
	// 单任务爬虫架构爬虫架构
	engine.SimpleEngine{}.Run(engine.Request{
		URL:        baseURL,
		ParserFunc: parser.ParserCityList,
	})

	// 并发版爬虫架构
	// e := engine.ConcurrentEngine{
	// 	Scheduler:   &scheduler.SimpleScheduler{},
	// 	WorkerCount: 1,
	// }
	// e.Run(engine.Request{
	// 	URL:        baseURL,
	// 	ParserFunc: parser.ParserCityList,
	// })
}
