package main

import (
	"crawler_go/frontend/controller"
	"fmt"
	"net/http"
)

func main() {

	http.Handle("/", http.FileServer(
		http.Dir("frontend/view")))

	http.Handle("/search",
		controller.CreateSearchResultHandler(
			"frontend/view/template.html"))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
