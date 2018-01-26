package main

import (
	"log"

	"fmt"

	"gopkg.in/olivere/elastic.v5"
	"imooc.com/ccmouse/learngo/crawler_distributed/config"
	"imooc.com/ccmouse/learngo/crawler_distributed/persist"
	"imooc.com/ccmouse/learngo/crawler_distributed/rpcsupport"
)

func main() {
	log.Fatal(serveRpc(
		fmt.Sprintf(":%d", config.ItemSaverPort),
		config.ElasticIndex))
}

func serveRpc(host, index string) error {
	client, err := elastic.NewClient(
		elastic.SetSniff(false))
	if err != nil {
		return err
	}

	return rpcsupport.ServeRpc(host,
		&persist.ItemSaverService{
			Client: client,
			Index:  index,
		})
}
