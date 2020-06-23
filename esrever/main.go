package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/miekg/dns"
)

func main() {
	if err := errorPipeless(); err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(os.Stdin)
	ips, err := ipsFromScanner(scanner)
	if err != nil {
		panic(err)
	}
	for _, ip := range ips {
		addr, err := dns.ReverseAddr(ip)
		if err != nil {
			panic(err)
		}
		fmt.Println(addr)
	}
}
