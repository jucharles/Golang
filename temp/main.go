package main

import (
	"flag"
	"fmt"
)

var temp = temp.CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}
