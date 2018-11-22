package view

import (
	"crawler/concurrent/engine"
	"crawler/concurrent/frontend/model"
	common "crawler/concurrent/model"
	"os"
	"testing"
)

func TestTemplate(t *testing.T) {

	 //template := template2.Must(
	 //	template2.ParseFiles("template.html"))
	view := CreateSearchResultView("template.html")

	 out, err := os.Create("template_test.html")
	 page := model.SearchResult{}
	 page.Hits = 123
	 //page.Start = 1
	item := engine.Item{
		Url: "http://album.zhenai.com/u/1902329077",
		Type: "zhenai",
		Id: "1902329077",
		Payload: common.Profile{
			Name: "麦甜",
			Age: 28,
			Gender: "女士",
			Height: 155,
			Income: "5-8千",
			Marriage: "丧偶 ",
			XingZuo: "天秤座",
			WorkPlace: "阿坝红原",
		},
	}

	for i := 0; i < 10; i ++ {
		page.Items = append(page.Items, item)
	}
	 //err := template.Execute(os.Stdout, page)
	 //err = template.Execute(out, page)
	 err = view.Render(out, page)
	 if err != nil {
	 	panic(err)
	 }
}