package main

import (
	"fmt"
	"strings"

	"github.com/bazzebazz/go-modules/slices"
	"rsc.io/quote/v3"
)

func main() {
	list := []string{"GolaNG", "GopherS", "JaVi", quote.HelloV3()}
	slices.Filter(list, func(item string) bool {
		return strings.HasPrefix(strings.ToLower(item), "h")
	})

	fmt.Println(quote.Concurrency())
}
