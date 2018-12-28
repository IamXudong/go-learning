// 输出标准输入的sha256, sha384或sha512散列值
package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

var option = flag.String("o", "256", "sha option 256, 384 or 512")

func main() {
	flag.Parse()
	var text string
	input := bufio.NewScanner(os.Stdin)

	if input.Scan() {
		text = input.Text()
	} else {
		os.Exit(1)
	}

	switch *option {
	case "256":
		fmt.Printf("%x\n", sha256.Sum256([]byte(text)))
	case "384":
		fmt.Printf("%x\n", sha512.Sum384([]byte(text)))
	case "512":
		fmt.Printf("%x\n", sha512.Sum512([]byte(text)))
	default:
		fmt.Println("Sha option must be 256, 384 or 512")
	}
}
