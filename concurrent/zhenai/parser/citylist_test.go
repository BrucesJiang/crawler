package parser

import (
	"testing"
)

func TestParseCityList(t *testing.T) {
	contents, err := fetcher.Fetch("http://www.zhenai.com/zhenghun")

	if err != nil {
		panic(err)
	}

	result := ParseCityList(contents)  

	const resultSize = 470

	expectedUrls := []string{
		"", "", "",
	}

	expectedCities := []string{
		"", "", "",    
	}

	if len(result.Requests) != resultSize {
		t.Errorf("result should have %d " + 
			"requests; but had %d",
			resultSize, len(result.Requests))
	}

	for i, url := range expectedUrls {
		if result.Requests[i].Url != url {
			t.Errorf("expected url #%d: %s; but " +
				"was %s",
				i, url, result.Requests[i].Url)
		}
	}

	if len(result.Items) != resultSize {
		t.Errorf("result should have %d " + 
		"items; but had %d",
		resultSize, len(result.Items))
	}

	for i, city := range expectedCities {
		if result.Items[i].(string) != city {
			t.Errorf("expected city #%d: %s; but " +
				"was %s",
				i, city, result.Requests[i].Url)
		}
	}
}