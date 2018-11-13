package parser

import (
	"crawler/concurrent/engine"
	"crawler/concurrent/model"
	"log"
	"regexp"
	"strconv"
)

var marriageRe = regexp.MustCompile(
	`<div class="m-btn purple"[^>]*>([^<]+)</div>`)

var ageRe = regexp.MustCompile(
	`<div class="m-btn purple"[^>]*>([0-9]+)岁</div>`)

var xingZuoRe = regexp.MustCompile(
	`<div class="m-btn purple"[^>]*>([^<]+)</div>`)

var heightRe = regexp.MustCompile(
	`<div class="m-btn purple"[^>]*>([0-9]+)cm</div>`)

var workPlaceRe = regexp.MustCompile(
	`<div class="m-btn purple"[^>]*>工作地:([^<]+)</div>`)

var occupationRe = regexp.MustCompile(
	`<div class="m-btn purple"[^>]*>([^<]+)</div>`)

var incomeRe = regexp.MustCompile(
	`<div class="m-btn purple"[^>]*>月收入:([^<]+)</div>`)

var eductionRe = regexp.MustCompile(
	`<div class="m-btn purple"[^>]*>([^<]+)/div>`)

func ParseProfile(contents []byte) engine.ParseResult {

	profile := model.Profile{}

	marriage := extractString(contents, marriageRe)
	profile.Marriage = marriage

	age, err := strconv.Atoi(extractString(contents, ageRe))
	if err == nil {
		profile.Age = age
	}

	xingZuo := extractString(contents, xingZuoRe)
	profile.XingZuo = xingZuo

	height, err := strconv.Atoi(extractString(contents, heightRe))
	if err == nil {
		profile.Height = height
	}

	workPlace := extractString(contents, workPlaceRe)
	profile.WorkPlace = workPlace

	occupation := extractString(contents, occupationRe)
	profile.Occupation = occupation

	income := extractString(contents, incomeRe)
	profile.Income = income

	eduction := extractString(contents, eductionRe)
	profile.Education = eduction

	log.Printf("Got Profile %v\n", profile)
	result := engine.ParseResult{
		Items: []interface{}{profile},
	}
	return result
}

func extractString(contents []byte, regexp *regexp.Regexp) string {
	matches := regexp.FindSubmatch(contents)

	if len(matches) >= 2 {
		return string(matches[1])
	}else {
		return ""
	}
}

