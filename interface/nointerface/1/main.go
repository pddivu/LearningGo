package main

import "fmt"

type vehicle struct {
	color  string
	number string
	speed  int
}

type car struct {
	vehicle
	model string
	year  int
}

type plane struct {
	vehicle
	country string
	year    int
}
type vehicles interface{}

func main() {
	prius := car{vehicle{"red", "7und", 100}, "red", 2012}
	crv := car{}
	corolla := car{}

	cars := []car{prius, crv, corolla}

	boeing1 := plane{}
	boeing2 := plane{}
	boeing3 := plane{}

	planes := []plane{boeing1, boeing2, boeing3}

	rides := []vehicles{prius, crv, corolla, boeing1, boeing2, boeing3}
	//rides:= []interface{}{prius, crv, corolla, boeing1, boeing2, boeing3}
	//rides := []vehicles{cars, planes}
	for key, value := range cars {
		fmt.Println(key, "-", value)
	}
	for key, value := range planes {
		fmt.Println(key, "-", value)
	}

	for key, value := range rides {
		fmt.Println(key, "-", value)
	}
}
