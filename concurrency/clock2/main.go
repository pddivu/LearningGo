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

	tz, link, err := parseArg(os.Args[1])
	if err != nil {
		log.Printf("[ERROR] Error occured: %v", err)
	}
	port := "localhost:" + link

	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
			continue
		}
		//tz := "Newyork-clock"

		go handleConn(conn, tz)

	}
}

func handleConn(c net.Conn, tz string) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, "\r"+tz+"    -"+(string(time.Now().Format("15:04:05"))))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}

func parseArg(arg string) (name string, link string, err error) {
	equal := strings.Index(arg, "=")
	if equal < 0 {
		message := fmt.Sprintf("Cannot parse arg: %v", arg)
		return "", "", errors.New(message)
	}
	return arg[:equal], arg[equal+1:], nil
}
