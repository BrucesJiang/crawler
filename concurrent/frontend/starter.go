package main

import (
	"crawler/concurrent/frontend/controller"
	"net/http"
)

func main() {
	 //访问本地CSS，图片以及js等资源拿不到。没有配置相应的服务
	//配置查找的根目录到view
	 http.Handle("/", http.FileServer(
	 	http.Dir("concurrent/frontend/view")))

	 http.Handle("/search",
	 	controller.CreateSearchResultHandler(
	 		"concurrent/frontend/view/template.html"))
	 err := http.ListenAndServe(":8888", nil)
	 if err != nil {
	 	panic(err)
	 }
}