package main

import (
	"gohosts/gohosts"
)

var Version = "(untracked)"

func main() {
	s := gohosts.Settings{}
	s.Read()

	gohosts.Merge(s.Hosts)
	gohosts.Write(s.Output)
}
