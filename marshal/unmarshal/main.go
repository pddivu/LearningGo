package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int `json:"New age"`
}

func main() {
	var p1 person
	fmt.Println(p1)
	fmt.Println(p1.First)
	bs := []byte(`{"First":"Divya","Last":"Divakaran","New age":34}`)
	json.Unmarshal(bs, &p1)
	fmt.Println(p1)
}
