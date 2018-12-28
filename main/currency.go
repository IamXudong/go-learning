// 声明币种数组
package main

import "fmt"

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	RMB
)

func main() {
	symbol := [...]string{USD: "＄", EUR: "€", GBP: "￡", RMB: "￥"}
	fmt.Printf("GBP: %d\t%s\n", GBP, symbol[GBP])
}
