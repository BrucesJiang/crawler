package main

import (
	"crawler-single/zhenai/parser"
	"crawler-single/engine"
)

func main() {
	engine.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
