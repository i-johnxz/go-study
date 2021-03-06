package main

import (
	"./engine"
	"./model"
	"./scheduler"
	"./zhenai/parser"
)

func main() {
	engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 1000,
	}.Run(model.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
