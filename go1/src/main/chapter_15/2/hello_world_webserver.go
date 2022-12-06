package main

import (
	"fmt"
	"log"
	"net/http"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Inside HelloServer handler")
	fmt.Fprintf(w, "Hello,"+req.URL.Path[1:])
}

func main() {
	//handler := http.HandlerFunc(HelloServer)
	//err := http.ListenAndServe(":8088", handler)

	err := http.ListenAndServe(":8088", http.HandlerFunc(HelloServer))
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}
