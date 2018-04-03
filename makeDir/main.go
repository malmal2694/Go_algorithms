package main

import (
	"flag"
	_ "fmt"
	"makeDir"
	_ "os"
	_ "path/filepath"
)

func main() {
	var numEpisodes, series int
	var path string

	flag.IntVar(&numEpisodes, "e", 12, "Defalut number of episode is 12")
	flag.IntVar(&series, "s", 1, "Defalut series is 1")
	flag.StringVar(&path, "p", "~/Videos/Dont Watch", "Defalut path is \"~/Videos/Dont Watch\"")
	flag.Parse()

	makeDir.MakeDir(numEpisodes, series, path)
}
