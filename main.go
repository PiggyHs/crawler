package main

import (
	"crawler/engine"
	"crawler/meishichina/parser"
	"crawler/scheduler"
)

func main() {
	eng := engine.ConcurrentEngine{
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 300,
	}
	eng.Run(engine.Request{
		Url:        "https://www.meishichina.com/YuanLiao/",
		ParserFunc: parser.ParseYuanliaoType,
	})
}
