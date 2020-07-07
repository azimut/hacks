package main

import (
	"bufio"
	"fmt"
	"os"
)

// TODO: compare paths? keep one with trailing slash?
func main() {
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
