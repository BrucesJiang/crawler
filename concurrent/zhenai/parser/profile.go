package parser

import (
	"crawler/concurrent/engine"
	"crawler/concurrent/model"
	"regexp"
	"strconv"
)

var marriageRe = regexp.MustCompile(
	`<div class="m-btn purple"[^>]*>([^<]+)</div>`)
//var xingZuoRe = regexp.MustCompile(
//	`<div class="m-btn purple"[^>]*>([^<]+)</div>`)
//var occupationRe = regexp.MustCompile(
//	`<div class="m-btn purple"[^>]*>([^<]+)</div>`)
//var eductionRe = regexp.MustCompile(
//	`<div class="m-btn purple"[^>]*>([^<]+)/div>`)


var ageRe = regexp.MustCompile(
	`<div class="m-btn purple"[^>]*>([0-9]+)岁</div>`)
var heightRe = regexp.MustCompile(
	`<div class="m-btn purple"[^>]*>([0-9]+)cm</div>`)

var workPlaceRe = regexp.MustCompile(
	`<div class="m-btn purple"[^>]*>工作地:([^<]+)</div>`)

var incomeRe = regexp.MustCompile(
	`<div class="m-btn purple"[^>]*>月收入:([^<]+)</div>`)



func ParseProfile(contents []byte) engine.ParseResult {

	profile := model.Profile{}

	//marriage := extractString(contents, marriageRe)
	//profile.Marriage = marriage
	//
	//xingZuo := extractString(contents, xingZuoRe)
	//profile.XingZuo = xingZuo
	//
	//occupation := extractString(contents, occupationRe)
	//profile.Occupation = occupation
	//
	//income := extractString(contents, incomeRe)
	//profile.Income = income
	//
	//eduction := extractString(contents, eductionRe)
	//profile.Education = eduction
	ms := extractField(contents, marriageRe)
	profile.Marriage = ms[0]
	profile.XingZuo = ms[1]
	profile.Occupation = ms[2]
	profile.Education = ms[3]

	age, err := strconv.Atoi(extractString(contents, ageRe))
	if err == nil {
		profile.Age = age
	}
	height, err := strconv.Atoi(extractString(contents, heightRe))
	if err == nil {
		profile.Height = height
	}

	workPlace := extractString(contents, workPlaceRe)
	profile.WorkPlace = workPlace



	//log.Printf("Got Profile %v\n", profile)
	result := engine.ParseResult{
		Items: []interface{}{profile},
	}
	return result
}

func extractField(contents []byte, regexp *regexp.Regexp) []string {
	mathes := regexp.FindAllSubmatch(contents, -1)

	ms := make([]string, 0)
	for _, match := range mathes {
		for _, m := range match {
			if len(m) >= 2 {
				ms = append(ms, string(m[1]))
			}else {
				ms = append(ms, "")
			}
		}
	}
	return ms
}


func extractString(contents []byte, regexp *regexp.Regexp) string {
	matches := regexp.FindSubmatch(contents)

	if len(matches) >= 2 {
		return string(matches[1])
	}else {
		return ""
	}
}

