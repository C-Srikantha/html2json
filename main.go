package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"task.com/task/service"
)

func handlereq() {
	mux := mux.NewRouter()
	mux.HandleFunc("/htmltojson", service.HtmlToJson)
	err := http.ListenAndServe(":8081", mux)
	if err != nil {
		fmt.Println(err)
	}
}
func main() {
	handlereq()
}
