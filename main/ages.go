// map键排序
package main

import (
	"fmt"
	"sort"
)

func main() {
	ages := map[string]int{"Jim": 24, "Bob": 36, "Sam": 47, "Sarah": 21, "Alice": 19}
	var names []string
	for name := range ages {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}
}
