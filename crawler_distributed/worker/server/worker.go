package main

import (
	"fmt"

	"log"

	"imooc.com/ccmouse/learngo/crawler_distributed/config"
	"imooc.com/ccmouse/learngo/crawler_distributed/rpcsupport"
	"imooc.com/ccmouse/learngo/crawler_distributed/worker"
)

func main() {
	log.Fatal(rpcsupport.ServeRpc(
		fmt.Sprintf(":%d", config.WorkerPort0),
		worker.CrawlService{}))
}
