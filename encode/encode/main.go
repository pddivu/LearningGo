package main

import (
	"encoding/json"
	"os"
)

type person struct {
	Fisrt       string
	Last        string
	Age         int
	notexported bool
}

func main() {
	p1 := person{"Divya", "Divakaran", 34, true}
	json.NewEncoder(os.Stdout).Encode(p1)
}
