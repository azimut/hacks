package main

import (
	"fmt"
	"strings"
)

type Domain struct {
	name        string
	rcode       string
	raddresses  []string
	hasWildcard bool
	subDomains  []*Domain
}

func (d Domain) String() string {
	return stringifyDomain(&d)
}

func stringifyDomain(domain *Domain) string {
	root := domain.name
	raw := stringifyDomains(domain.subDomains, root)
	return strings.Join(raw, "\n")
}

func stringifyDomains(domains []*Domain, root string) []string {
	ret := make([]string, 0)
	for _, subdomain := range domains {
		current := subdomain.name + "." + root
		if len(subdomain.subDomains) == 0 {
			ret = append(ret, current)
		} else {
			ret = append(ret, stringifyDomains(subdomain.subDomains, current)...)
		}
	}
	return ret
}

func printBranch(domain *Domain, root string) {
	printBranchRaw(domain, root)
}
func printBranchRaw(domain *Domain, acc string) {
	fmt.Println(domain.name + "." + acc)
	for _, subdomain := range domain.subDomains {
		printBranchRaw(subdomain, domain.name+"."+acc)
	}
}
func returnInvalid(domain *Domain) {
	for _, subdomain := range domain.subDomains {
		if subdomain.rcode == "NXDOMAIN" {
			printBranch(subdomain, domain.name)
		} else {
			returnInvalid(subdomain)
		}
	}
}
