package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	str := ""
	fmt.Scanf("%s", &str)
	res, err := http.Get(str)
	checkError(err)
	data, err := ioutil.ReadAll(res.Body)
	checkError(err)
	fmt.Printf("Got: %q", string(data))
}

func checkError(err error) {
	if err != nil {
		log.Fatalf("Get : %v", err)
	}
}
