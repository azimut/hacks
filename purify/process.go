package main

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
	for _, subdomain := range domain.subDomains {
		if len(subdomain.subDomains) > 0 {
			current := subdomain.name + "." + root
			processRoots(subdomain, current)
		}
	}
	return nil
}
