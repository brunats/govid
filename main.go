package main

import (
	"fmt"

	"github.com/brunats/govid/internal/cli"
)

func main() {
	cli.Parse()

	country := cli.Country()
	format := cli.Format()

	fmt.Println("country:", country)
	fmt.Println("format:", format)
}
