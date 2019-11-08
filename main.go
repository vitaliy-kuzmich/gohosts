package main

import (
	"fmt"
	"gohosts/gohosts"
)

var Version = "(untracked)"

func main() {
	fmt.Println("gohosts version", Version)
	s := gohosts.Settings{}
	s.Read()

	gohosts.Merge(s.Hosts)
	gohosts.Write(s.Output)
}
