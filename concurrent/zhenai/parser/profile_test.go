package parser

import (
	"crawler/concurrent/engine"
	"crawler/concurrent/model"
	"io/ioutil"
	"testing"
)


func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("user.html")

	if err != nil {
		panic(err)
	}
	//fmt.Printf("%s\n", contents)
	item := engine.Item{}
	profile := model.Profile{
		Name: "麦甜",
		Age: 28,
		Height: 155,
		Marriage: "丧偶 ",
	}
	item.Payload = profile
	result := ParseProfile(contents, &item)

	expected := model.Profile{
		Name: "麦甜",
		Age: 28,
		Height: 155,
		Income: "5-8千",
		Marriage: "丧偶 ",
		WorkPlace: "阿坝红原",
		//Occupation: "人事主管",
		//Education: "中专",
	}

	for _, p := range result.Items {
		profile := p.Payload.(model.Profile)
		if profile != expected {
			t.Errorf("exptect Profile is %+v; but " +
				"was %+v\n", expected, profile)
		}
	}

	//fmt.Printf("%s\n", contents)
}