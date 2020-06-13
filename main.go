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

// maxLevels returns the max number of subdomains levels
func maxLevels(domains []*publicsuffix.DomainName) int {
	var maximun int
	for _, domain := range domains {
		current := domainLevel(domain)
		if current > maximun {
			maximun = current
		}
	}
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

// TODO: remove duplicate domains
func main() {
	finfo, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}
	if finfo.Mode()&os.ModeNamedPipe == 0 {
		fmt.Fprintln(os.Stderr, "error: no stdin pipe data")
		os.Exit(1)
	}
	scanner := bufio.NewScanner(os.Stdin)
	domains := make([]*publicsuffix.DomainName, 0)
	for scanner.Scan() {
		new, err := parseDomain(scanner.Text())
		if err != nil {
			panic(err)
		}
		domains = append(domains, new)
	}
	if !hasSameSLD(domains) {
		fmt.Fprintln(os.Stderr, "error: different domain in provided list")
		os.Exit(1)
	}
	for _, domain := range getLevel(1, domains) {
		fmt.Printf("%s\n", domain)
	}
}
