package main

import (
	"crawler_go/engine"
	"crawler_go/zhenai/parser"
)

const baseURL = "https://www.zhenai.com/zhenghun"

func main() {
	engine.Run(engine.Request{
		URL:        baseURL,
		ParserFunc: parser.ParserCityList,
	})
}
