// 捕获迭代变量
package main

import (
	"os"
)

func main() {
	tempDirs := os.Args[1:]
	var rmdirs []func()
	for _, dir := range tempDirs {
		dir := dir
		os.MkdirAll(dir, 0755)
		rmdirs = append(rmdirs, func() {
			os.RemoveAll(dir)
		})
	}

	for _, rmdir := range rmdirs {
		rmdir()
	}
}
