package parser

import (
	"crawler/concurrent/engine"
	"crawler/concurrent/model"
	"log"
	"regexp"
	"strconv"
)

var urlRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>[^<]+</a>`)

var nameRe = regexp.MustCompile(`<a href="http://album.zhenai.com/u/[0-9]+"[^>]*>([^<]+)</a>`)

var idRe = regexp.MustCompile(`http://album.zhenai.com/u/([0-9]+)`)
//const incomingRe = `<td><span class="grayL">月[^<]+</span>([0-9\-]+)元</td>`

var genderRe = regexp.MustCompile(`<td width="180"><span class="grayL">性别[\pP]+</span>([^<]+)</td>`)

var ageRe = regexp.MustCompile(`<td width="180"><span class="grayL">年龄[\pP]+</span>([0-9]+)</td>`)

var marriageRe = regexp.MustCompile(`<td width="180"><span[^>]+>婚[^<]+</span>([^<]+)</td>`)

var workPlaceRe = regexp.MustCompile(`<td><span class="grayL">居住地[\pP]+</span>([^<]+)</td>`)

var heightRe = regexp.MustCompile(`<td width="180"><span class="grayL">身[^<]+</span>([0-9]+)</td>`)

func ParseCity(contents []byte) engine.ParseResult{

	urls := extractField(contents, urlRe)
	names := extractField(contents, nameRe)
	ids := extractField(contents, idRe)
	genders := extractField(contents, genderRe)
	ages := extractField(contents, ageRe)
	mars := extractField(contents, marriageRe)
	workplaces := extractField(contents, workPlaceRe)
	heights := extractField(contents, heightRe)

	result := engine.ParseResult{}
	for i, m := range urls {
		profile := model.Profile{}
		profile.Name = names[i]
		profile.Gender = genders[i]
		profile.WorkPlace = workplaces[i]
		age, err := strconv.Atoi(ages[i])
		if err == nil {
			profile.Age = age
		}

		profile.Marriage = mars[i]
		height, err := strconv.Atoi(heights[i])
		if err == nil {
			profile.Height = height
		}
		log.Printf("Profile %+v\n", profile)
		item := engine.Item{
			Url: urls[i],
			Id: ids[i],
			Payload: profile,
		}

		//name := string(m[2])
		//result.Items = append(
		//	result.Items,
		//	"User :" + name)
		result.Requests = append(
			result.Requests,
			engine.Request{
				Url: m,
				ParseFunc: func(bytes []byte) engine.ParseResult {
					return ParseProfile(bytes, &item)
				},
			})
		//log.Printf("Got User %v\n", profile)
	}

	return result
}

//func extraceFields(contents []byte, regexp *regexp.Regexp) [][]string {
//	matches := regexp.FindAllSubmatch(contents, -1)
//
//	//log.Printf("Item: %s, Number: %d\n", matches, len(matches))
//
//	ms := make([][]string, 0)
//	urls := make([]string, 0)
//	names := make([]string,0)
//	for _, m := range matches {
//		if len(m) >= 3 {
//			urls = append(urls, string(m[1]))
//			names = append(names, string(m[2]))
//		}else {
//			urls = append(urls, "")
//			names = append(names, "")
//		}
//	}
//	ms = append(ms, urls)
//	ms = append(ms, names)
//	return ms
//}

func extractField(contents []byte, regexp *regexp.Regexp) []string {
	matches := regexp.FindAllSubmatch(contents, -1)

	//log.Printf("Item: %s, Number: %d\n", matches, len(matches))

	ms := make([]string, 0)

	for _, m := range matches {
		if len(m) >= 2 {
			ms = append(ms, string(m[1]))
		}else {
			ms = append(ms, "")
		}
	}
	return ms
}
