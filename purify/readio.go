package main

import (
	"bufio"
	"errors"
	"os"
	"strings"

	"github.com/weppos/publicsuffix-go/publicsuffix"
)

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
