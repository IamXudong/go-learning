package main

import (
	"io"
	"log"
	"net"
	"os"
	"fmt"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	go mustCopy(os.Stdout, conn, done)
	mustCopy(conn, os.Stdin, nil)
	<-done
	conn.Close()
}

func mustCopy(dst io.Writer, src io.ReadWriter, ch chan struct{}) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
	if ch != nil {
		fmt.Println("done")
		ch <- struct{}{}
	}
}
