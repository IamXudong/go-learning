package main

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
)

type myreader struct {
	i   int
	buf []byte
}

func (r *myreader) Read(p []byte) (n int, err error) {
	if r.i >= len(r.buf) {
		err = io.EOF
		return
	}
	n = copy(p, r.buf[r.i:])
	r.i += n
	return
}

func newReader(str string) io.Reader {
	var r myreader
	r.buf = []byte(str)
	return &r
}

type mylimitreader struct {
	r      io.Reader
	remain int64
}

func (lr *limitreader) Read(p []byte) (n int, err error) {
	if lr.remain <= 0 {
		return 0, io.EOF
	}
	if len(p) > lr.remain {
		p = p[0:lr.remain]
	}
	n, err = lr.r.Read(p)
	lr.remain -= int64(n)
	return
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &mylimitreader{r, n}
}

func main() {
	r := newReader("<html><body><h1>Hello world</h1><h2>Hello word2</h2></body></html>")
	doc, _ := html.Parse(r)
	fmt.Println(doc.FirstChild.LastChild.LastChild.FirstChild.Data)
}
