package main

import (
	"fmt"

	"github.com/stevzhang01/go-learning/strings"
)

func main() {
	fmt.Printf("%s\t=\t%v\n", "strings.HasPrefix('abc', 'ab')", strings.HasPrefix("abc", "ab"))
	fmt.Printf("%s\t=\t%v\n", "strings.HasPrefix('aec', 'ab')", strings.HasPrefix("aec", "ab"))
	fmt.Printf("%s\t=\t%v\n", "strings.HasSuffix('eab', 'ab')", strings.HasSuffix("eab", "ab"))
	fmt.Printf("%s\t=\t%v\n", "strings.HasSuffix('eab', 'ac')", strings.HasSuffix("eab", "ac"))
	fmt.Printf("%s\t=\t%v\n", "strings.Contain('abkef', 'bk')", strings.Contain("abkef", "bk"))
	fmt.Printf("%s\t=\t%v\n", "strings.Contain('abkef', 'be')", strings.Contain("abkef", "be"))
}
