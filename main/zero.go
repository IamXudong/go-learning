// [32]byte数组清0
package main

import "fmt"

func main() {
	a := [32]byte{31: 'a'}
	b := [32]byte{30: 'b'}
	fmt.Println(a)
	zeroA(&a)
	fmt.Println(a)

	fmt.Println(b)
	zeroB(&b)
	fmt.Println(b)

}

func zeroA(ptr *[32]byte) {
	for i := range ptr {
		ptr[i] = 0
	}
}

func zeroB(ptr *[32]byte) {
	*ptr = [32]byte{}
}
