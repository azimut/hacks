package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if err := errorPipeless(); err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(os.Stdin)
	domains, err := domainsFromScanner(scanner)
	if err != nil {
		panic(err)
	}
	root := domains[0].SLD + "." + domains[0].TLD
	roots := &Domain{name: root}
	for _, domain := range domains {
		addAtom(domain.TRD, roots)
	}
	printNodes(roots, roots.name)
}

func printNodes(domain *Domain, root string) {
	fmt.Println(root)
	for _, subdomain := range domain.subDomains {
		if len(subdomain.subDomains) > 0 {
			printNodes(subdomain, subdomain.name+"."+root)
		}
	}
}
