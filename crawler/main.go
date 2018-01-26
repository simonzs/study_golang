package main

import (
	"imooc.com/ccmouse/learngo/crawler/engine"
	"imooc.com/ccmouse/learngo/crawler/persist"
	"imooc.com/ccmouse/learngo/crawler/scheduler"
	"imooc.com/ccmouse/learngo/crawler/zhenai/parser"
	"imooc.com/ccmouse/learngo/crawler_distributed/config"
)

func main() {
	itemChan, err := persist.ItemSaver(
		"dating_profile")
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan:    itemChan,
	}

	e.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(
			parser.ParseCityList,
			config.ParseCityList),
	})
}
