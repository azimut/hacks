package main

import "strings"

// naiveAddAtom adds trd into domains
func naiveAddAtom(trd string, domains *Domain) {
	dotted := strings.Split(trd, ".")
	current := dotted[len(dotted)-1]
	domains.subDomains = append(domains.subDomains, &Domain{name: current})
	if len(dotted) > 1 {
		newtrd := strings.Join(dotted[0:len(dotted)-1], ".")
		naiveAddAtom(newtrd, domains.subDomains[len(domains.subDomains)-1])
	}
}

func addAtom(trd string, domains *Domain) {
	if trd == "" {
		return
	}
	dotted := strings.Split(trd, ".")
	current := dotted[len(dotted)-1]
	base := &Domain{}
	for _, value := range domains.subDomains {
		if current == value.name {
			base = value
			break
		}
	}
	if base.name == "" {
		naiveAddAtom(trd, domains)
	} else {
		addAtom(strings.Join(dotted[0:len(dotted)-1], "."), base)
	}
}

func getAtom(trd string, domains *Domain) *Domain {
	dotted := strings.Split(trd, ".")
	current := dotted[len(dotted)-1]
	for _, value := range domains.subDomains {
		if current == value.name {
			if len(dotted) == 1 {
				return value
			} else {
				newtrd := strings.Join(dotted[0:len(dotted)-1], ".")
				return getAtom(newtrd, value)
			}
		}
	}
	return nil
}
