package parser

import (
	"crawler/concurrent/engine"
	"crawler/concurrent/model"
	"log"
	"regexp"
	"strconv"
)

//const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

const userRe = `<table><tbody><tr><th><a href="(http://album.zhenai.com/u/[0-9])"[^>]*>([^<]+)</a></th></tr> <tr><td width="180"><span class="grayL">性别：</span>([^<]+)</td> <td><span class="grayL">居住地：</span>([^<]+)</td></tr> <tr><td width="180"><span class="grayL">年龄：</span>([^<]+)</td> <td><span class="grayL">学   历：</span>([^<]+)</td> <!----></tr> <tr><td width="180"><span class="grayL">婚况：</span>([^<]+)</td> <td width="180"><span class="grayL">身   高：</span>([^<]+)</td></tr></tbody></table>`

func ParseCity(contents []byte) engine.ParseResult{
	//re := regexp.MustCompile(cityRe)
	re := regexp.MustCompile(userRe)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}

	profile := model.Profile{}
	for _, m := range matches {
		profile.Url = string(m[1])
		profile.Name = string(m[2])
		profile.Gender = string(m[3])
		profile.WorkPlace = string(m[4])
		age, err := strconv.Atoi(string(m[5]))
		if err == nil {
			profile.Age = age
		}
		profile.Education = string(m[6])
		profile.Marriage = string(m[7])
		height, err := strconv.Atoi(string(m[5]))
		if err == nil {
			profile.Height = height
		}

		//name := string(m[2])
		//result.Items = append(
		//	result.Items,
		//	"User :" + name)
		result.Requests = append(
			result.Requests,
			engine.Request{
				Url: string(m[1]),
				ParseFunc: func(bytes []byte) engine.ParseResult {
					return ParseProfile(bytes, profile)
				},
			})
		log.Printf("Got User %v\n", profile)
	}

	return result
}