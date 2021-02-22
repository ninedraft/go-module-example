package main

import (
	_ "embed"
	"fmt"
)

//go:embed go.mod
var src string

func main() {
	fmt.Println(src)
}
