package main

import "fmt"

type Domain struct {
	name       string
	subDomains []*Domain
}

func (d Domain) String() string {
	return stringifyDomains(&d)
}

// TODO: is backwards
func stringifyDomains(domain *Domain) string {
	var acc string
	for _, value := range domain.subDomains {
		if len(value.subDomains) == 0 {
			acc += fmt.Sprintf("%s\n", value.name)
		} else {
			acc += value.name + "."
			acc += stringifyDomains(value)
		}
	}
	return acc
}
