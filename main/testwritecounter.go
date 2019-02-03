package main

import (
	"fmt"
	"github.com/stevzhang01/go-learning/writecounter"
	"os"
)

func main() {
	fmt.Println("===test ByteCounter")
	var c writecounter.ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c)
	c = 0
	name := "Dolly"
	fmt.Fprintf(&c, "Hello, %s", name)
	fmt.Println(c)
	fmt.Println("===test LineCounter")
	var l writecounter.LineCounter
	fmt.Fprintf(&l, "Hello, %s\n%s", name, name)
	fmt.Println(l)
	fmt.Println("===test WordsCounter")
	var w writecounter.WordsCounter
	fmt.Fprintf(&w, "Hello, %s", name)
	fmt.Println(w)
	fmt.Println("==test CountingWriter")
	writer, count := writecounter.CountingWriter(os.Stdout)
	fmt.Fprintf(writer, "Hello, %s\n", name)
	fmt.Println("*count: ", *count)
}
