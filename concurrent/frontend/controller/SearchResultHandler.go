package controller

import (
	"context"
	"crawler/concurrent/engine"
	"crawler/concurrent/frontend/model"
	"crawler/concurrent/frontend/view"
	"gopkg.in/olivere/elastic.v5"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)


//TODO list
// fill in query string
// support search button
// rewrite query string
// support paging
// add start page

type SearchResultHandler struct {
	view   view.SearchResultView
	client *elastic.Client
}

func CreateSearchResultHandler(
	template string) SearchResultHandler {
	client, err := elastic.NewClient(
		elastic.SetSniff(false))

	if err != nil {
		panic(err)
	}
	return SearchResultHandler{
		view:   view.CreateSearchResultView(template),
		client: client,
	}
}

// localhost:8888/search?q=男士 离异&from=20
func (h SearchResultHandler) ServeHTTP(
	w http.ResponseWriter, req *http.Request) {
	q := strings.TrimSpace(req.FormValue("q"))
	//q = rewriteQueryString(q)
	from, err := strconv.Atoi(
		req.FormValue("from"))

	if err != nil {
		from = 0
	}
	//fmt.Fprintf(w, "q=%s, from=%d\n", q, from)

	var page model.SearchResult
	page, err = h.getSearchResult(q, from)

	if err != nil {
		http.Error(w, err.Error(),
			http.StatusBadRequest)
	}
	err = h.view.Render(w, page)

	if err != nil {
		http.Error(w, err.Error(),
			http.StatusBadRequest)
	}
}

func (h SearchResultHandler) getSearchResult(
	q string, from int) (model.SearchResult, error) {
	var result model.SearchResult
	resp, err := h.client.
		Search("dating_profile").
		Query(elastic.NewQueryStringQuery(
			rewriteQueryString(q))).
		From(from).
		Do(context.Background())

	if err != nil {
		return result, err
	}
	result.PrevFrom = result.Start - len(result.Items)
	result.NextFrom = result.Start + len(result.Items)
	result.Query = q
	result.Hits = resp.TotalHits()
	result.Start = from
	result.Items = resp.Each(
		reflect.TypeOf(engine.Item{}))
	//for _, v := range resp.Each(
	//	reflect.TypeOf(engine.Item{})) {
	//		item := v.(engine.Item)
	//}
	return result, nil
}

func rewriteQueryString(q string) string {
	re := regexp.MustCompile(`([A-Z][a-z]*):`)
	return re.ReplaceAllString(q, "Payload.$1:")
}