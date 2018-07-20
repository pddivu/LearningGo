package main

import (
	"encoding/json" //encoding
	"fmt"
)

type person struct {
	First      string //Capital letter to export
	Last       string
	Age        int
	unexported int // this will not be exported
}

func main() {
	p1 := person{"Divya", "Divakaran", 34, 6}
	//json to data transfer
	bs, _ := json.Marshal(p1)
	fmt.Println(p1)
	fmt.Println(bs)
	fmt.Printf("%T \n", bs) //[]uint8
	fmt.Println(string(bs)) //{"First":"Divya","Last":"Divakaran","Age":34}
}
