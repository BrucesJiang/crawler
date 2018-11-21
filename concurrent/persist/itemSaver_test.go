package persist

import (
	"crawler/concurrent/engine"
	"crawler/concurrent/model"
	"encoding/json"
	"golang.org/x/net/context"
	"gopkg.in/olivere/elastic.v5"
	"testing"
)

func TestSave(t *testing.T) {
	expected := engine.Item{
		Url: "http://album.zhenai.com/u/104063556",
		Type: "zhenai",
		Id: "104063556",
		Payload: model.Profile{

			Name: "真爱永恒",
			Gender: "男士",
			Age: 35,
			Height: 167,
			Income: "8千-1.2万",
			Marriage: "离异",
			XingZuo: "",
			WorkPlace: "内蒙古阿拉善盟",
		},
	}

	//TODO: Try to start up elastic search
	// here using docker go client
	client, err := elastic.NewClient(
		elastic.SetSniff(false))

	if err != nil {
		panic(err)
	}
	const index = "dating_test"
	//Save expected item
	err = save(client, index, expected)

	if err != nil {
		panic(err)
	}

	// Fetch saved item
	resp, err := client.Get().
		Index(index).
		Type(expected.Type).
		Id(expected.Id).
		Do(context.Background())

	if err != nil {
		panic(err)
	}

	t.Logf("%+v", resp)
	var actual engine.Item
	bts , err := resp.Source.MarshalJSON()

	t.Logf("%s", bts)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(
		bts, &actual)

	if err != nil {
		panic(err)
	}
	actualProfile, err := model.FromJsonObj(actual.Payload)
	actual.Payload = actualProfile
	//Verify result
	if actual != expected{
		t.Errorf("got %+v; exptected %+v\n",actual, expected)
	}
}