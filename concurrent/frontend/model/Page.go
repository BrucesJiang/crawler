package model

import "crawler/concurrent/engine"

type SearchResult struct {
	Hits  int
	Start int
	Items []engine.Item
}
