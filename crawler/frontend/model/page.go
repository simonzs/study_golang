package model

import "imooc.com/ccmouse/learngo/crawler/engine"

type SearchResult struct {
	Hits  int
	Start int
	Items []engine.Item
}
