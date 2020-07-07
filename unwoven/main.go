package main

import (
	"bufio"
	"fmt"
	"os"
)

// - check all are urls
func main() {
	if err := errorPipeless(); err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(os.Stdin)
	urls, err := urlsFromScanner(scanner)
	if err != nil {
		panic(err)
	}
	fmt.Println(urls)
}
