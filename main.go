package main

import (
	"crawler/engine"
	"crawler/meishichina/parser"
	"crawler/persist"
	"crawler/scheduler"
)

func main() {
	itemChan, err := persist.ItemSaver(
		"data_food")
	if err != nil {
		panic(err)
	}

	eng := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueueScheduler{},
		WorkerCount: 200,
		ItemChan:    itemChan,
	}
	eng.Run(engine.Request{
		Url:        "https://www.meishichina.com/YuanLiao/",
		ParserFunc: parser.ParseYuanliaoType,
	})
	//eng := engine.ConcurrentEngine{
	//	Scheduler:   &scheduler.QueueScheduler{},
	//	WorkerCount: 200,
	//}
	//eng.Run(engine.Request{
	//	Url:        "https://www.readnovel.com/free/all?pageSize=10&gender=2&catId=30013&isFinish=1&isVip=1&size=-1&updT=-1&orderBy=0&pageNum=1",
	//	ParserFunc: yuewenparser.ParseArticle,
	//})
}
