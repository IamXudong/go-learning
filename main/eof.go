package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	for {
		r, _, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("read failed: %v\n", err)
		}
		fmt.Printf("%s", string(r))
	}
}
