package parser

import (
	"crawler/single/engine"
	"crawler/single/model"
	"fmt"
	"regexp"
	"strconv"

)

var marriageRe = regexp.MustCompile(
	`<div class="m-btn purple"[^>]*>([^<]+)</div>`)
var xingZuoRe = regexp.MustCompile(
	`<div class="m-btn purple"[^>]*>([^<]+)</div>`)
var occupationRe = regexp.MustCompile(
	`<div class="m-btn purple"[^>]>([^<]+)</div>`)
var eductionRe = regexp.MustCompile(
	`<div class="m-btn purple"[^>]*>([^<]+)/div>`)

var ageRe = regexp.MustCompile(
	`<div class="m-btn purple"[^>]*>([0-9]+)岁</div>`)

var heightRe = regexp.MustCompile(
	`<div class="m-btn purple"[^>]*>([0-9]+)cm</div>`)

var hukouRe = regexp.MustCompile(
	`<div class="m-btn purple"[^>]*>工作地:([^<]+)</div>`)

var incomeRe = regexp.MustCompile(
	`<div class="m-btn purple"[^>]*>月收入:([^<]+)</div>`)

func ParseProfile(
	contents []byte, name string) engine.ParseResult {

		profile := model.Profile{}

		profile.Name = name

		age, err := strconv.Atoi(extractString(contents, ageRe))
		if err == nil {
			profile.Age = age
		}

		marriage := extractString(contents, marriageRe)
		profile.Marriage = marriage

		ms := extractField(contents, marriageRe)
		fmt.Printf("MS #%s\n", ms)

		xingZuo := extractString(contents, xingZuoRe)
		profile.XingZuo = xingZuo

		height, err := strconv.Atoi(extractString(contents, heightRe))
		if err == nil {
			profile.Height = height
		}

		hukou := extractString(contents, hukouRe)
		profile.Hukou = hukou

		occupation := extractString(contents, occupationRe)
		profile.Occupation = occupation

		income := extractString(contents, incomeRe)
		profile.Income = income

		eduction := extractString(contents, eductionRe)
		profile.Education = eduction

		result := engine.ParseResult{
			Items: []interface{}{profile},
		}		
		return result
}

func extractField(contents []byte, regexp *regexp.Regexp) []string {
	mathes := regexp.FindAllSubmatch(contents, -1)

	ms := make([]string, 0)
	for _, match := range mathes {
		if len(match) >= 2 {
			ms = append(ms, string(match[1]))
		}else {
			ms = append(ms, "")
		}
	}
	return ms
}


func extractString(
	contents []byte, re *regexp.Regexp) string{
		match := re.FindSubmatch(contents)
		
		if len(match) >= 2 {
			fmt.Printf("%s", match[1])
			return string(match[1])
		}else {
			//fmt.Printf("%s", match[0])
			//fmt.Printf("匹配个数  %d\n", len(match))
			return " "
		}
}