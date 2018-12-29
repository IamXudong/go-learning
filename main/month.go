// slice相关操作
package main

import "fmt"

var month = [...]string{1: "January", 2: "February", 3: "March", 4: "April", 5: "may", 6: "June", 7: "July", 8: "August", 9: "September", 10: "October", 11: "November", 12: "December"}
var q2 = month[4:7]
var summer = month[6:9]

func main() {
	fmt.Println(q2)
	fmt.Println(summer)

	// 查找两个slice中相同的元素
	for _, s := range summer {
		for _, q := range q2 {
			if s == q {
				fmt.Printf("%s apears in both\n", s)
			}
		}
	}

	endlessSummer := summer[:5]
	fmt.Println(endlessSummer)
}
