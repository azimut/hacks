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
	branches := Domain{name: root}
	for _, domain := range domains {
		addAtom(domain.TRD, &branches)
	}
	if err := resolveWildcards(&branches, branches.name); err != nil {
		panic(err)
	}
}

// Unify CORE
func resolveWildcards(domain *Domain, root string) error {
	if len(domain.subDomains) == 0 {
		return nil
	}
	resolve, err := doesResolve(root)
	if err != nil {
		return err
	}
	if rcode(resolve) != "NOERROR" {
		return nil
	}
	wild, data, err := hasWildcard(root)
	if err != nil {
		return err
	}
	// fmt.Println(root)
	// fmt.Println(".")
	// fmt.Println(wild)
	// fmt.Println(data)
	if wild && len(data) > 0 {
		for _, result := range data {
			fmt.Printf("%s %s\n", root, result)
		}
		return nil
	}
	//fmt.Println("before")
	for _, subdomain := range domain.subDomains {
		//fmt.Println("in")
		resolveWildcards(subdomain, subdomain.name+"."+root)
	}
	//fmt.Println("after")
	return nil
}
