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

// domainCurate removes junk from domain
func domainCurate(domain string) (string, error) {
	trimmed := strings.TrimSpace(domain)
	lowered := strings.ToLower(trimmed)
	if strings.Contains(lowered, "..") {
		return "", errors.New("Contains .. can't fix it...yet")
	}
	undotted := strings.Trim(lowered, ".")
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
		text := scanner.Text()
		new, err := parseDomain(text)
		if err != nil {
			return nil, err
		}
		// drop dups
		for _, domain := range domains {
			if new == domain {
				continue
			}
		}
		domains = append(domains, new)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return domains, nil
}
