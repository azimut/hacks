package main

import (
	"bufio"
	"os"
)

func main() {
	if err := errorPipeless(); err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(os.Stdin)
	nmaps, err := nmapsFromScanner(scanner)
	if err != nil {
		panic(err)
	}
	for _, nmap := range nmaps {
		if err := printNmap(nmap); err != nil {
			panic(err)
		}
	}
}
