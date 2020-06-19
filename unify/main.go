package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//TODO: permutation?
//TODO: integration test?
// flag to add custom domain
func main() {
	if err := errorPipeless(); err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(os.Stdin)
	subdomains, err := domainsFromScanner(scanner)
	if err != nil {
		panic(err)
	}
	domains := &Domain{}
	for _, subdomain := range subdomains {
		trd := subdomain.TRD
		addAtom(trd, domains)
	}
	printChains(domains, "")
	//permutateChains(domains)
}

func printChains(domain *Domain, acc string) {
	if acc != "" {
		fmt.Println(strings.TrimSuffix(acc, "."))
	}
	for _, sub := range domain.subDomains {
		printChains(sub, sub.name+"."+acc)
	}
}
