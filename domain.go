package main

import "strings"

type Domain struct {
	name       string
	subDomains []*Domain
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
