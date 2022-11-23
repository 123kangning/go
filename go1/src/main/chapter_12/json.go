package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Address struct {
	Type    string
	City    string
	Country string
}

type VCard struct {
	FirstName string
	LastName  string
	Addresses []*Address
	Remark    string
}

func main() {
	testMarshal()
	testUnmarshal()
}

func testMarshal() {
	pa := &Address{"private", "Aartselaar", "Belgium"}
	wa := &Address{"work", "Boom", "Belgium"}
	vc := VCard{"Jan", "Kersschot", []*Address{pa, wa}, "none"}
	// fmt.Printf("%v: \n", vc) // {Jan Kersschot [0x126d2b80 0x126d2be0] none}:
	// JSON format:
	js, _ := json.Marshal(vc)
	fmt.Printf("JSON format: %s", js)
	// using an encoder:
	file, _ := os.OpenFile("vcard.json", os.O_CREATE|os.O_WRONLY, 0666)
	defer file.Close()
	enc := json.NewEncoder(file)
	err := enc.Encode(vc)
	if err != nil {
		log.Println("Error in encoding json")
	}
}
func testUnmarshal() {
	file, errFile := os.Open("vcard.json")
	if errFile != nil {
		fmt.Println("Open error , err = ", errFile)
		return
	}
	defer file.Close()
	newV := new(VCard)
	buf := make([]byte, 512)
	nr, errRead := file.Read(buf)
	if errRead != nil {
		fmt.Println("file.Read error , err = ", errRead)
		return
	}
	errUnmarshal := json.Unmarshal(buf[:nr], newV)
	if errUnmarshal != nil {
		fmt.Println("Unmarshal error , err = ", errUnmarshal)
	}
	fmt.Println(newV)
}
