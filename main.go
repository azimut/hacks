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

// domainLevel returns the level of root domain from .
func domainLevel(domain string) int {
	return strings.Count(domain, ".")
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
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	result, err := publicsuffix.Parse("1.www.example.com")
	fmt.Println(result.TLD)
	fmt.Println(result.SLD)
	fmt.Println(result.TRD)
}
