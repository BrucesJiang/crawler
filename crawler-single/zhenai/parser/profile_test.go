package parser

import (
	"io/ioutil"
	"testing"
	"crawler-single/model"

)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("user.html")

	if err != nil {
		panic(err)
	}
	//fmt.Printf("%s\n", contents)
	result := ParseProfile(contents, "")

	expectedMarriage := "丧偶"
	expectedHuKou := "阿坝红原"

	for _, p := range result.Items {
		profile := p.(model.Profile)

		if profile.Marriage != expectedMarriage || 
			profile.Hukou != expectedHuKou {
				t.Errorf("exptect Marriage %s and HuKou %s; but " +
					"was %s, %s\n", expectedMarriage, expectedHuKou, 
					profile.Marriage, profile.Hukou)
			}
	}

	//fmt.Printf("%s\n", contents)
}