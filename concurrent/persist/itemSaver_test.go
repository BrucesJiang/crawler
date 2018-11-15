package persist

import (
	"crawler/concurrent/model"
	"log"
	"testing"
)

func TestItemSaver(t *testing.T) {
	expected := model.Profile{
		Age : 34,
		Height: 162,
		Income: "3001-5000元",
		Gender: "女",
		Name: "安静的雪",
	}
	log.Printf("%+v", expected)
}