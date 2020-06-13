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

// sameSLD validates a slice for same domain on all
func hasSameSLD(domains []*publicsuffix.DomainName) bool {
	defsld := domains[0].SLD
	for _, domain := range domains {
		if defsld != domain.SLD {
			return false
		}
	}
	return true
}

// domainLevel returns the level of root domain from .
func domainLevel(TRD string) int {
	return strings.Count(TRD, ".")
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
		raw := scanner.Text()
		curated, err := domainCurate(raw)
		if err != nil {
			panic(err)
		}
		parsed, err := publicsuffix.Parse(curated)
		if err != nil {
			panic(err)
		}
		domains = append(domains, parsed)
	}
	if !hasSameSLD(domains) {
		fmt.Fprintln(os.Stderr, "error: different domain in provided list")
		os.Exit(1)
	}
}
