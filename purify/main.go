package main

import (
	"bufio"
	"os"
)

// TODO: explode domain atoms and shuffle in subdomain?
// TODO: return wildcard ips
// TODO: use some domains from the same level to check for wildcards too
func main() {
	seedRandom()
	if err := errorPipeless(); err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(os.Stdin)
	domains, err := domainsFromScanner(scanner)
	if err != nil {
		panic(err)
	}
	// Build the recursive struct
	root := domains[0].SLD + "." + domains[0].TLD
	branches := Domain{name: root}
	for _, domain := range domains {
		addAtom(domain.TRD, &branches)
	}
	// Resolve nodes of the struct, filter out NX
	if err := processDomain(&branches); err != nil {
		panic(err)
	}
	// fmt.Println("---------------------")
	// fmt.Println(domains)
	// fmt.Println("---------------------")
	// returnInvalid(&branches)
	// fmt.Println("---------------------")
	returnValid(&branches, branches.name)
}
