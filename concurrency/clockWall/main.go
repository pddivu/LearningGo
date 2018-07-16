package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	servers := make(map[string]string)

	for _, v := range os.Args[1:] {
		//fmt.Println(v)
		if name, link, err := parseArg(v); err != nil {
			log.Printf("[ERROR] Error occured: %v", err)
		} else {
			servers[name] = link
			fmt.Printf("Name: %s; Link %s\n", name, link)
		}
	}
	for name, link := range servers {
		go func(name, link string) {
			port := "localhost:" + link
			conn, err := net.Dial("tcp", port)
			//fmt.Printf("Name: %s; Link %s\n", name, link)
			if err != nil {
				log.Fatal(err)
			}
			defer conn.Close()
			mustCopy(name, os.Stdout, conn)
		}(name, link)
	}
	select {}

}

func mustCopy(timezone string, dst io.Writer, src io.Reader) {
	/*	_, err := io.WriteString(dst, timezone)
		if err != nil {
		}*/
	//cw := timeWriter{timezone, dst}

	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

type timeWriter struct {
	timezone string
	origin   io.Writer
}

func (cw timeWriter) Write(p []byte) (n int, err error) {
	var totalN int
	//for {
	//fmt.Println(string(p))
	//n, err = fmt.Fprintf(cw.origin, "%s:\t%s", cw.timezone, string(p))
	//n, err := fmt.Fprintf(cw.origin, "\r%s:", string(p))
	if err != nil {
		return totalN, err
	}
	totalN += n
	time.Sleep(1 * time.Second)
	//}

	return totalN, nil
}
func parseArg(arg string) (name string, link string, err error) {
	equal := strings.Index(arg, "=")
	if equal < 0 {
		message := fmt.Sprintf("Cannot parse arg: %v", arg)
		return "", "", errors.New(message)
	}
	return arg[:equal], arg[equal+1:], nil
}
