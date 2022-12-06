package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

type empty struct{}

func main() {
	var e empty
	err := http.ListenAndServe(":8081", e)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}

func (in empty) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Println(in)
	if strings.Compare(req.URL.Path[1:6], "hello") == 0 {
		fmt.Fprintf(w, "hello %s", req.URL.Path[7:])
	} else if strings.Compare(req.URL.Path[1:11], "shouthello") == 0 {
		fmt.Fprintf(w, "hello %s", strings.ToUpper(req.URL.Path[12:]))
	}
}
