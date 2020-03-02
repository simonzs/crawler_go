package main

import (
	"crawler_go/engine"
	"crawler_go/persist"
	"crawler_go/scheduler"
	"crawler_go/zhenai/parser"
)

func main() {
	// 单任务爬虫架构爬虫架构
	// engine.SimpleEngine{}.Run(engine.Request{
	// 	URL:        baseURL,
	// 	ParserFunc: parser.ParserCityList,
	// })

	itemChan, err := persist.ItemSaver(
		"dating_profile")
	if err != nil {
		panic(err)
	}

	// 并发版爬虫架构 Request
	e := engine.ConcurrentEngine{
		// Scheduler: &scheduler.SimpleScheduler{},
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan:    itemChan,
	}

	cityListURL := "https://www.zhenai.com/zhenghun"
	e.Run(engine.Request{
		URL:        cityListURL,
		ParserFunc: parser.ParserCityList,
	})

	// cityURL := "http://www.zhenai.com/zhenghun/chengdu"
	// e.Run(engine.Request{
	// 	URL:        cityURL,
	// 	ParserFunc: parser.ParserCity,
	// })
}
