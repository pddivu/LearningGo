package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {

	for

		port := "localhost:" + os.Args[i]
		conn, err := net.Dial("tcp", port)
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()
		go mustCopy(os.Stdout, conn)
	}
}

func mustCopy(dst io.Writer, src io.Reader) {
	fmt.Printf(format)
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
