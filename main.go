package main

import (
	"flag"
	"log"

	"github.com/rtward/advent-of-code-2023/day01"
)

func main() {
	flag.Parse()
	day := flag.Arg(0)

	switch day {
	case "01":
		day01.Run()
	default:
		log.Fatalf("Invalid day specified: \"%s\"", day)
	}
}
