package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var n = flag.Bool("n", false, "Omit trailing newline")
var sep = flag.String("s", " ", "Field separator")

func main() {
	fmt.Println(os.Args[:])
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}
