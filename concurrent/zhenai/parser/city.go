package parser

import (
	"crawler/concurrent/engine"
	"log"
	"regexp"
)

const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

func ParseCity(contents []byte) engine.ParseResult{
	re := regexp.MustCompile(cityRe)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}

	for _, m := range matches {
		name := string(m[2])
		result.Items = append(
			result.Items,
			"User :" + name)
		result.Requests = append(
			result.Requests,
			engine.Request{
				Url: string(m[1]),
				ParseFunc: ParseProfile,
			})
		log.Printf("Got User: %s, Url: %s\n", name, m[1])
	}

	return result
}