package main

import (
	"fmt"
	_ "github.com/paradoxxl/go-web01/rest01"
	"github.com/paradoxxl/go-web01/rping"
)

func main() {
	fmt.Println("Trying to start Server")
	rping.Start()
}
