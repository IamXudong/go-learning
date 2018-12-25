// 测试boolconv包的实现
package main

import (
	"fmt"
	"go-learning/boolconv"
)

func main() {
	fmt.Printf("%s\t=\t%v\n", "boolconv.Btoi(true)", boolconv.Btoi(true))
	fmt.Printf("%s\t=\t%v\n", "boolconv.ItoB(8)", boolconv.Itob(8))
	fmt.Printf("%s\t=\t%v\n", "boolconv.ItoB(0)", boolconv.Itob(0))
	fmt.Printf("%s\t=\t%v\n", "boolconv.ItoB(5)", boolconv.Itob(5))
}
