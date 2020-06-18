package main

import (
	"bufio"
	"fmt"
	"os"
)

// TODO: remove duplicate domains
// TODO: fill missing domains (between levels)
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
	root := domains[0].SLD + "." + domains[0].TLD
	branches := Domain{name: root}
	for _, domain := range domains {
		addAtom(domain.TRD, &branches)
	}
	fmt.Println(branches)
	err = processDomain(&branches)
	if err != nil {
		panic(err)
	}
}

func processDomain(domain *Domain) error {
	if err := processRoots(domain, domain.name); err != nil {
		return err
	}
	return nil
}

// processRoots updates DOMAIN argument with reply status
func processRoots(domain *Domain, root string) error {
	reply, err := doesResolve(root)
	if err != nil {
		return err
	}
	domain.rcode = rcode(reply)
	domain.raddresses = getAnswers(reply)
	if len(domain.subDomains) == 0 || rcode(reply) == "NXDOMAIN" {
		return nil
	}
	// wildcard, _, err := hasWildcard(domain.name)
	// if err != nil {
	// 	return err
	// }
	// if wildcard {
	// 	fmt.Printf("WILDCARD: %s\n", domain.name)
	// 	domain.hasWildcard = true
	// 	return nil
	// }
	for _, subdomain := range domain.subDomains {
		if len(subdomain.subDomains) > 0 {
			current := subdomain.name + "." + root
			processRoots(subdomain, current)
		}
	}
	return nil
}
