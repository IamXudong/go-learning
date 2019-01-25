// fetch 下载 URL 并返回本地文件的名字和长度
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

func main() {
	url := os.Args[1]
	file, size, err := fetch(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch %s: %v\n", url, err)
		os.Exit(1)
	}
	fmt.Printf("file: %s\nsize: %d\n", file, size)
}

func fetch(url string) (string, int64, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()
	local := path.Base(resp.Request.URL.Path)
	if local == "/" || local == "." {
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	n, err := io.Copy(f, resp.Body)
	if closeErr := f.Close(); err == nil {
		err = closeErr
	}
	return local, n, err
}
