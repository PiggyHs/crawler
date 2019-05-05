package main

import (
	"crawler/engine"
	"crawler/meishichina/parser"
	"crawler/scheduler"
)

func main() {
	eng := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueueScheduler{},
		WorkerCount: 200,
	}
	eng.Run(engine.Request{
		Url:        "https://www.meishichina.com/YuanLiao/",
		ParserFunc: parser.ParseYuanliaoType,
	})
}
