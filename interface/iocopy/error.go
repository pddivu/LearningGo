package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	msg := "I am a good person"
	rdr := strings.NewReader(msg)
	_, err := io.Copy(os.Stdout, rdr)
	if err != nil {
		fmt.Println(err)
		return
	}

	rdr2 := bytes.NewBuffer([]byte(msg))
	_, err = io.Copy(os.Stdout, rdr2)
	if err != nil {
		fmt.Println(err)
		return
	}

	res, err := http.Get("http://192.168.1.1")
	if err != nil {
		fmt.Println(err)
		return
	}
	io.Copy(os.Stdout, res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
}
