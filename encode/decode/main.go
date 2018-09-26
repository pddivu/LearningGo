package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	var p1 person
	data := strings.NewReader(`{"First":"Divya","Last":"Divakaran","Age":34}`)
	json.NewDecoder(data).Decode(&p1)
	fmt.Println(data)
}
