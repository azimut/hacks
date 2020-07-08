package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var timeout int

// TODO: compare paths? keep one with trailing slash?
func main() {
	flag.IntVar(&timeout, "t", 5, "timeout after seconds")
	flag.Parse()
	if err := errorPipeless(); err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(os.Stdin)
	urls, err := urlsFromScanner(scanner)
	if err != nil {
		panic(err)
	}
	for _, url := range urls {
		fmt.Println(url)
	}
}
