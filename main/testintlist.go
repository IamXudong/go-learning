package main

import (
	"fmt"
	"github.com/stevzhang01/go-learning/intlist"
)

func main() {
	list := intlist.IntList{
		5,
		&intlist.IntList{8, nil},
	}
	fmt.Println(list.Sum())
}
