// 常量声明
package main

import (
	"fmt"
)

type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func main() {
	fmt.Printf("Sunday\t=\t%[1]v\n", Sunday)
	fmt.Printf("Monday\t=\t%[1]v\n", Monday)
	fmt.Printf("Tuesday\t=\t%[1]v\n", Tuesday)
	fmt.Printf("Wednesday\t=\t%[1]v\n", Wednesday)
	fmt.Printf("Thursday\t=\t%[1]v\n", Thursday)
	fmt.Printf("Friday\t=\t%[1]v\n", Friday)
	fmt.Printf("Saturday\t=\t%[1]v\n", Saturday)
}
