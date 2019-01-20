package main

import (
	"fmt"
	"github.com/stevzhang01/go-learning/github"
	"log"
	"os"
	"time"
)

func main() {
	trange := os.Args[1]
	result, err := github.SearchIssues(os.Args[2:])
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	fmt.Printf("%d Issues: \n", result.TotalCount)
	switch trange {
	case "M":
		fmt.Println("Items in one Month:")
	case "Y":
		fmt.Println("Items in one Year:")
	case "O":
		fmt.Println("Items over one Year:")
	default:
	}
	now := time.Now()
	for _, item := range result.Items {
		m := item.CreatedAt.Month()
		y := item.CreatedAt.Year()
		switch trange {
		case "M":
			if y != now.Year() || m != now.Month() {
				goto End
			}
		case "Y":
			if y != now.Year() {
				goto End
			}
		case "O":
			if y == now.Year() {
				continue
			}
		default:
		}
		fmt.Printf("| #%-5d | %s |%9.9s | %.55s\n", item.Number, item.CreatedAt, item.User.Login, item.Title)
	}
End:
}
