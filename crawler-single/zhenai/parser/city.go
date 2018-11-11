package parser

import (
	"crawler/crawler-single/engine"
	"regexp"

)


const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`
func ParseCity(
	contents []byte) engine.ParseResult {
		re :=  regexp.MustCompile(cityRe)
		matches := re.FindAllSubmatch(contents, -1)
		
		result := engine.ParseResult{}
		for _, m := range matches {
			name := string(m[2])
			result.Items = append(
				result.Items, "User " + string(m[2]))
			result.Requests = append(
				result.Requests, engine.Request{
					Url: string(m[1]),
					ParserFunc: func (c []byte)  engine.ParseResult{
						return ParseProfile(c, name)
					}, //函数式编程的思想，采用闭包的概念，
					//因为添加后面一个name字段，Requests API的API更改会导致其它地方
					//代码更改过多，因此采用闭包的概念。这很常用
			})
		}
		return result
}