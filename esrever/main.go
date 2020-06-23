package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

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
		addr = strings.TrimRight(addr, ".")
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s,%s\n", ip, addr)
	}
}
