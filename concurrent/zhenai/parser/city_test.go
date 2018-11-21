package parser

import (
	"crawler/concurrent/fetcher"
	"log"
	"testing"
)

func TestParseCity(t *testing.T) {
	//contents, err := ioutil.ReadFile("city.html")
	contents, err := fetcher.Fetch("http://www.zhenai.com/zhenghun/laibin")
	if err != nil {
		panic(err)
	}

	//log.Printf("%s\n", contents)
	result := ParseCity(contents)

	log.Printf("%+v\n", result)

}
