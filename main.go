package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/weppos/publicsuffix-go/publicsuffix"
)

// domainCurate removes junk from domain
func domainCurate(domain string) (string, error) {
	trimmed := strings.TrimSpace(domain)
	lowered := strings.ToLower(trimmed)
	if strings.Contains(lowered, "..") {
		return "", errors.New("Contains .. can't fix it...yet")
	}
	undotted := strings.Trim(lowered, ".")
	if !strings.Contains(undotted, ".") {
		return "", errors.New("Does NOT contains any dot.")
	}
	return undotted, nil
}

// hasSameSLD validates a slice for same domain on all
func hasSameSLD(domains []*publicsuffix.DomainName) bool {
	defsld := domains[0].SLD
	for _, domain := range domains {
		if defsld != domain.SLD {
			return false
		}
	}
	return true
}

// parseDomain ...
func parseDomain(domain string) (*publicsuffix.DomainName, error) {
	curated, err := domainCurate(domain)
	if err != nil {
		return nil, err
	}
	parsed, err := publicsuffix.Parse(curated)
	if err != nil {
		return nil, err
	}
	return parsed, nil
}

// maxLevel returns the max number of subdomains levels
func maxLevel(domains []*publicsuffix.DomainName) int {
	var maximun int
	for _, domain := range domains {
		current := domainLevel(domain)
		if current > maximun {
			maximun = current
		}
	}
	//fmt.Printf("Max level: %+v\n", maximun) // output for debug
	return maximun
}

func domainLevel(domain *publicsuffix.DomainName) int {
	return strings.Count(domain.TRD, ".")
}

func getLevel(level int, domains []*publicsuffix.DomainName) []*publicsuffix.DomainName {
	newdomains := make([]*publicsuffix.DomainName, 0)
	for _, domain := range domains {
		if level == domainLevel(domain) {
			newdomains = append(newdomains, domain)
		}
	}
	return newdomains
}

// domainsFromScanner ...
func domainsFromScanner(scanner *bufio.Scanner) ([]*publicsuffix.DomainName, error) {
	domains := make([]*publicsuffix.DomainName, 0)
	for scanner.Scan() {
		new, err := parseDomain(scanner.Text())
		if err != nil {
			panic(err)
		}
		domains = append(domains, new)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	if !hasSameSLD(domains) {
		return nil, errors.New("different domain in provided list")
	}
	return domains, nil
}

// errorPipeless errors if program is not called with pipe input
func errorPipeless() error {
	finfo, err := os.Stdin.Stat()
	if err != nil {
		return err
	}
	if finfo.Mode()&os.ModeNamedPipe == 0 {
		return errors.New("is not called with pipe data")
	}
	return nil
}

// kvDomain return a valid key value for a domain
func kvDomain(domain *publicsuffix.DomainName) (key, value string) {
	level := domainLevel(domain)
	root := domain.SLD + "." + domain.TLD
	if level == 0 {
		return root, domain.TRD
	} else {
		splitted := strings.Split(domain.TRD, ".")
		key = strings.Join(splitted[1:level+1], ".") + "." + root
		value = strings.Join(splitted[0:1], ".")
		return key, value
	}
}

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
	processDomain(&branches)
	fmt.Println(len(count))
}

func processDomain(domains *Domain) error {
	base := domains.name
	state, err := doesResolve(base)
	if err != nil {
		return err
	}
	if state != "NOERROR" {
		return errors.New(fmt.Sprintf("weird return state for root (%s)", state))
	}
	//
	processSubdomains(domains.subDomains, base)
	return nil
}

var count = make(map[string]int)

func processSubdomains(domains []*Domain, root string) error {
	for _, subdomain := range domains {
		current := subdomain.name + "." + root
		count[current]++
		if subdomain.subDomains == nil {
			fmt.Println(current)
		} else {
			processSubdomains(subdomain.subDomains, current)
		}
	}
	return nil
}
