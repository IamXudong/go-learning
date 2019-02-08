package main

import (
	"bytes"
	"fmt"
	"io"
)

const debug = false

func main() {
	var buf io.Writer
	if debug {
		buf = new(bytes.Buffer)
	}
	f(buf)
	fmt.Println("end")
}

func f(w io.Writer) {
	if w != nil {
		w.Write([]byte("debug info"))
		fmt.Println(w)
	} else {
		fmt.Println("no debug")
	}
}
