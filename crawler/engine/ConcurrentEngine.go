package engine

import (
	"../fetcher"
	"../model"
	"../scheduler"
	"log"
)

// 并发引擎
type ConcurrentEngine struct {
	// 调度器
	Scheduler   scheduler.Scheduler
	WorkerCount int
}

func (e ConcurrentEngine) Run(seeds ...model.Request) {
	// 初始化 Scheduler 的队列，并启动配对 goroutine
	e.Scheduler.Run()
	out := make(chan model.ParseResult)
	for i := 0; i < e.WorkerCount; i++ {
		createWorker(e.Scheduler.WorkerChann(), out, e.Scheduler)
	}
	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	for {
		result := <-out
		for _, item := range result.Items {
			log.Printf("getItems, items: %v", item)
		}

		for _, r := range result.Requests {
			e.Scheduler.Submit(r)
		}
	}
}

func createWorker(in chan model.Request, out chan model.ParseResult, notifier scheduler.ReadyNotifier) {
	go func() {
		for {
			notifier.WorkerReady(in)
			r := <-in
			result, err := worker(r)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}

func worker(r model.Request) (model.ParseResult, error) {
	log.Printf("fetching url: %s", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("fetch error, url: %s, err: %v", r.Url, err)
		return model.ParseResult{}, err
	}
	return r.ParserFunc(body), nil
}
